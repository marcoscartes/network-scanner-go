package main

import (
	"flag"
	"log"
	"network-scanner-go/internal/database"
	"network-scanner-go/internal/history"
	"network-scanner-go/internal/notifications"
	"network-scanner-go/internal/scanner"
	"network-scanner-go/internal/security"
	"network-scanner-go/internal/web"
	"sync"
	"time"
)

func main() {
	// Parse command line flags
	ipRange := flag.String("range", "", "IP range to scan (e.g., 192.168.1.0/24)")
	interval := flag.Int("interval", 60, "Scan interval in seconds")
	webPort := flag.String("web-port", "5050", "Web interface port")
	dbPath := flag.String("db", "scanner.db", "Database file path")

	// Notification flags
	notifyNewDevices := flag.Bool("notify-new-devices", true, "Notify when new devices are detected")
	notifyDisconnected := flag.Bool("notify-disconnected", true, "Notify when devices disconnect")
	notifyPortChanges := flag.Bool("notify-port-changes", true, "Notify when port changes are detected")
	webhookURL := flag.String("webhook-url", "", "Webhook URL for notifications")
	notificationRetentionDays := flag.Int("notification-retention", 7, "Days to retain notifications")

	// History flags
	historyRetentionDays := flag.Int("history-retention-days", 90, "Number of days to retain historical data")

	flag.Parse()

	// Initialize database
	err := database.Init(*dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Detect network range if not specified
	if *ipRange == "" {
		detected, err := scanner.GetLocalNetwork()
		if err != nil {
			log.Printf("Failed to detect network: %v. Using default 192.168.1.0/24", err)
			*ipRange = "192.168.1.0/24"
		} else {
			*ipRange = detected
			log.Printf("Auto-detected network range: %s", *ipRange)
		}
	}

	// Load existing devices
	devices, err := database.GetAllDevices()
	if err != nil {
		log.Printf("Failed to load devices: %v", err)
	} else {
		log.Printf("Loaded %d devices from database", len(devices))
	}

	// Initialize notification system
	notificationManager := notifications.NewManager()
	defer notificationManager.Stop()

	// Update notification config with webhook if provided
	if *webhookURL != "" {
		config := notificationManager.GetConfig()
		config.WebhookURL = *webhookURL
		config.EnabledChannels = append(config.EnabledChannels, "webhook")
		if err := notificationManager.UpdateConfig(config); err != nil {
			log.Printf("Failed to update notification config: %v", err)
		}
	}

	// Initialize change detector
	detector := notifications.NewDetector()
	if len(devices) > 0 {
		detector.UpdateState(devices)
	}

	// Load security rules
	security.LoadRules(security.GetDefaultRulesPath())

	log.Println("Notification system initialized")

	// Start web server in goroutine
	server := web.NewServer(*webPort)
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Web server failed: %v", err)
		}
	}()

	// Initialize history tracking
	lastStatsDay := ""
	lastSnapshotTime := time.Time{}

	// Main scanning loop
	for {
		now := time.Now()
		log.Printf("Starting network scan for %s", *ipRange)

		// Discover devices
		discoveredDevices, err := scanner.DiscoverDevices(*ipRange)
		if err != nil {
			log.Printf("Scan error: %v", err)
			time.Sleep(time.Duration(*interval) * time.Second)
			continue
		}

		log.Printf("Found %d active devices", len(discoveredDevices))

		// Enrich devices in parallel
		var wg sync.WaitGroup
		for _, device := range discoveredDevices {
			wg.Add(1)
			go func(d *database.Device) {
				defer wg.Done()

				// Load existing device data to preserve previously discovered ports
				existingDevices, _ := database.GetAllDevices()
				var existingPorts []int
				for _, existing := range existingDevices {
					if existing.MAC == d.MAC {
						existingPorts = existing.OpenPorts
						break
					}
				}

				scanner.IdentifyDevice(d)

				if len(d.MetricsURLs) > 0 {
					log.Printf("Found metrics at: %v on %s", d.MetricsURLs, d.IP)
				}

				// Merge ports: combine existing ports with newly discovered ports
				if len(existingPorts) > 0 {
					portMap := make(map[int]bool)
					// Add existing ports
					for _, port := range existingPorts {
						portMap[port] = true
					}
					// Add newly discovered ports
					for _, port := range d.OpenPorts {
						portMap[port] = true
					}
					// Convert back to slice
					mergedPorts := make([]int, 0, len(portMap))
					for port := range portMap {
						mergedPorts = append(mergedPorts, port)
					}
					// Sort ports
					for i := 0; i < len(mergedPorts); i++ {
						for j := i + 1; j < len(mergedPorts); j++ {
							if mergedPorts[i] > mergedPorts[j] {
								mergedPorts[i], mergedPorts[j] = mergedPorts[j], mergedPorts[i]
							}
						}
					}
					d.OpenPorts = mergedPorts
				}

				// Check for vulnerabilities
				d.Vulnerabilities = security.CheckDevice(d.OpenPorts, d.Type)

				// Save to database
				if err := database.UpsertDevice(d); err != nil {
					log.Printf("Failed to save device %s: %v", d.IP, err)
				}
			}(device)
		}

		wg.Wait()

		server.Broadcast(map[string]interface{}{
			"type": "discovery_complete",
			"data": discoveredDevices,
		})

		// Detect and notify changes
		previousDevices := make([]*database.Device, 0)
		for _, dev := range detector.GetPreviousDevices() {
			previousDevices = append(previousDevices, dev)
		}

		changes := detector.CompareDeviceStates(previousDevices, discoveredDevices)
		for _, change := range changes {
			// Check if notification is enabled for this type
			shouldNotify := false
			switch change.Type {
			case "new_device":
				shouldNotify = *notifyNewDevices
			case "disconnected":
				shouldNotify = *notifyDisconnected
			case "port_change":
				shouldNotify = *notifyPortChanges
			}

			if shouldNotify {
				if err := notificationManager.NotifyChange(change); err != nil {
					log.Printf("Failed to send notification: %v", err)
				}
				server.Broadcast(map[string]interface{}{
					"type": "notification",
					"data": change,
				})
			}
		}

		// Update detector state
		detector.UpdateState(discoveredDevices)

		// Record network snapshot for historical tracking
		if err := history.RecordNetworkSnapshot(discoveredDevices); err != nil {
			log.Printf("Failed to record network snapshot: %v", err)
		}

		// Record individual device changes
		for _, change := range changes {
			changeType := "update"
			switch change.Type {
			case "new_device":
				changeType = "new"
			case "disconnected":
				changeType = "disconnect"
			}
			if err := history.RecordDeviceState(change.Device, changeType); err != nil {
				log.Printf("Failed to record device change: %v", err)
			}
		}

		// Check for daily stats calculation
		// Run if it's a new day since last calculation
		currentDay := now.Format("2006-01-02")
		if lastStatsDay != currentDay {
			log.Printf("New day detected (%s). Calculating daily statistics...", currentDay)
			if _, err := history.CalculateDailyStats(now); err != nil {
				log.Printf("Failed to calculate daily stats: %v", err)
			} else {
				lastStatsDay = currentDay
			}

			// Clean old history data
			if err := history.CleanOldHistory(*historyRetentionDays); err != nil {
				log.Printf("Failed to clean old history: %v", err)
			}

			// Clean old notifications
			if err := database.DeleteOldNotifications(*notificationRetentionDays); err != nil {
				log.Printf("Failed to clean old notifications: %v", err)
			}
		}

		// Record network snapshot periodically (e.g., every hour)
		if now.Sub(lastSnapshotTime) >= 1*time.Hour {
			log.Println("Recording hourly network snapshot...")
			if err := history.RecordNetworkSnapshot(discoveredDevices); err != nil {
				log.Printf("Failed to record network snapshot: %v", err)
			}
			lastSnapshotTime = now
		}

		log.Printf("Scan complete. Sleeping for %d seconds...", *interval)
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}
