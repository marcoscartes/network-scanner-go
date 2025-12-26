package web

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"network-scanner-go/internal/database"
	"network-scanner-go/internal/scanner"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

//go:embed templates/index.html
var indexHTML string

var (
	scanState = make(map[string]*database.ScanProgress)
	stateMu   sync.RWMutex
)

// Server represents the web server
type Server struct {
	router *mux.Router
	port   string
}

// NewServer creates a new web server
func NewServer(port string) *Server {
	s := &Server{
		router: mux.NewRouter(),
		port:   port,
	}

	s.setupRoutes()
	return s
}

// setupRoutes configures the HTTP routes
func (s *Server) setupRoutes() {
	s.router.HandleFunc("/", s.handleIndex).Methods("GET")
	s.router.HandleFunc("/api/scan-all-ports/{ip}", s.handleScanAllPorts).Methods("POST")
	s.router.HandleFunc("/api/scan-progress/{ip}", s.handleScanProgress).Methods("GET")

	// Notification endpoints
	s.router.HandleFunc("/api/notifications", s.handleGetNotifications).Methods("GET")
	s.router.HandleFunc("/api/notifications/{id}/read", s.handleMarkNotificationRead).Methods("POST")
	s.router.HandleFunc("/api/notifications/{id}", s.handleDeleteNotification).Methods("DELETE")
	s.router.HandleFunc("/api/notifications/config", s.handleGetNotificationConfig).Methods("GET")
	s.router.HandleFunc("/api/notifications/config", s.handleUpdateNotificationConfig).Methods("PUT")

	// History and Statistics endpoints
	s.router.HandleFunc("/api/history/device/{mac}", s.handleGetDeviceHistory).Methods("GET")
	s.router.HandleFunc("/api/history/network", s.handleGetNetworkHistory).Methods("GET")
	s.router.HandleFunc("/api/stats/overview", s.handleGetStatsOverview).Methods("GET")
	s.router.HandleFunc("/api/stats/trends", s.handleGetNetworkTrends).Methods("GET")
	s.router.HandleFunc("/api/stats/uptime/{mac}", s.handleGetDeviceUptime).Methods("GET")
}

// handleIndex renders the dashboard
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	devices, err := database.GetAllDevices()
	if err != nil {
		http.Error(w, "Failed to load devices", http.StatusInternalServerError)
		return
	}

	// Sort by IP address
	sort.Slice(devices, func(i, j int) bool {
		return devices[i].IP < devices[j].IP
	})

	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		log.Printf("Template parse error: %v", err)
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"devices": devices,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Template execute error: %v", err)
	}
}

// handleScanAllPorts initiates a full port scan
func (s *Server) handleScanAllPorts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	stateMu.Lock()
	if progress, exists := scanState[ip]; exists && progress.Status == "running" {
		stateMu.Unlock()
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Scan already running for this IP",
		})
		return
	}

	// Initialize scan state
	scanState[ip] = &database.ScanProgress{
		Status:      "running",
		Progress:    0,
		CurrentPort: 0,
		TotalPorts:  65535,
		OpenPorts:   []int{},
		PortsFound:  0,
		StartTime:   time.Now(),
	}
	stateMu.Unlock()

	// Start scan in goroutine
	go func() {
		log.Printf("Starting full port scan for %s\n", ip)

		openPorts := scanner.ScanAllPorts(ip, func(current, total int, ports []int) {
			stateMu.Lock()
			if state, ok := scanState[ip]; ok {
				state.Progress = (current * 100) / total
				state.CurrentPort = current
				state.OpenPorts = ports
				state.PortsFound = len(ports)
			}
			stateMu.Unlock()
		})

		// Mark as complete
		stateMu.Lock()
		if state, ok := scanState[ip]; ok {
			state.Status = "complete"
			state.Progress = 100
			now := time.Now()
			state.EndTime = &now
			state.OpenPorts = openPorts
			state.PortsFound = len(openPorts)
		}
		stateMu.Unlock()

		// Update database
		devices, _ := database.GetAllDevices()
		for _, device := range devices {
			if device.IP == ip {
				device.OpenPorts = openPorts
				database.UpsertDevice(device)
				break
			}
		}

		log.Printf("Full port scan complete for %s. Found %d open ports\n", ip, len(openPorts))
	}()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "started",
		"ip":     ip,
	})
}

// handleScanProgress returns the progress of a port scan
func (s *Server) handleScanProgress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	stateMu.RLock()
	progress, exists := scanState[ip]
	stateMu.RUnlock()

	if !exists {
		http.Error(w, "No scan found for this IP", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":       progress.Status,
		"progress":     progress.Progress,
		"current_port": progress.CurrentPort,
		"total_ports":  progress.TotalPorts,
		"open_ports":   progress.OpenPorts,
		"ports_found":  progress.PortsFound,
		"elapsed_time": progress.ElapsedTime(),
	}

	json.NewEncoder(w).Encode(response)
}

// Start starts the web server
func (s *Server) Start() error {
	log.Printf("Starting web server on port %s\n", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}

// handleGetNotifications returns all notifications
func (s *Server) handleGetNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := database.GetAllNotifications()
	if err != nil {
		http.Error(w, "Failed to load notifications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// handleMarkNotificationRead marks a notification as read
func (s *Server) handleMarkNotificationRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var notificationID int
	_, err := fmt.Sscanf(id, "%d", &notificationID)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	err = database.MarkNotificationAsRead(notificationID)
	if err != nil {
		http.Error(w, "Failed to mark notification as read", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// handleDeleteNotification deletes a notification
func (s *Server) handleDeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var notificationID int
	_, err := fmt.Sscanf(id, "%d", &notificationID)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteNotification(notificationID)
	if err != nil {
		http.Error(w, "Failed to delete notification", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// handleGetNotificationConfig returns the notification configuration
func (s *Server) handleGetNotificationConfig(w http.ResponseWriter, r *http.Request) {
	config, err := database.GetNotificationConfig()
	if err != nil {
		http.Error(w, "Failed to load notification config", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// handleUpdateNotificationConfig updates the notification configuration
func (s *Server) handleUpdateNotificationConfig(w http.ResponseWriter, r *http.Request) {
	var config database.NotificationConfig
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = database.SaveNotificationConfig(&config)
	if err != nil {
		http.Error(w, "Failed to save notification config", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// handleGetDeviceHistory returns the history of a specific device
func (s *Server) handleGetDeviceHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mac := vars["mac"]

	// Get time range from query parameters (default: last 30 days)
	days := 30
	if daysParam := r.URL.Query().Get("days"); daysParam != "" {
		fmt.Sscanf(daysParam, "%d", &days)
	}

	from := time.Now().AddDate(0, 0, -days)
	to := time.Now()

	history, err := database.GetDeviceHistory(mac, from, to)
	if err != nil {
		http.Error(w, "Failed to load device history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

// handleGetNetworkHistory returns the network history
func (s *Server) handleGetNetworkHistory(w http.ResponseWriter, r *http.Request) {
	// Get time range from query parameters (default: last 30 days)
	days := 30
	if daysParam := r.URL.Query().Get("days"); daysParam != "" {
		fmt.Sscanf(daysParam, "%d", &days)
	}

	trends, err := database.GetNetworkTrends(days)
	if err != nil {
		http.Error(w, "Failed to load network history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trends)
}

// handleGetStatsOverview returns an overview of network statistics
func (s *Server) handleGetStatsOverview(w http.ResponseWriter, r *http.Request) {
	devices, err := database.GetAllDevices()
	if err != nil {
		http.Error(w, "Failed to load devices", http.StatusInternalServerError)
		return
	}

	// Calculate statistics
	totalDevices := len(devices)
	activeDevices := 0
	totalPorts := 0
	now := time.Now()
	oneDayAgo := now.Add(-24 * time.Hour)

	for _, device := range devices {
		if device.LastSeen.After(oneDayAgo) {
			activeDevices++
		}
		totalPorts += len(device.OpenPorts)
	}

	// Get today's stats
	todayStats, _ := database.CalculateDailyStats(now)

	overview := map[string]interface{}{
		"total_devices":      totalDevices,
		"active_devices":     activeDevices,
		"total_ports":        totalPorts,
		"new_devices_today":  0,
		"disconnected_today": 0,
	}

	if todayStats != nil {
		overview["new_devices_today"] = todayStats.NewDevices
		overview["disconnected_today"] = todayStats.DisconnectedDevices
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(overview)
}

// handleGetNetworkTrends returns network trends over time
func (s *Server) handleGetNetworkTrends(w http.ResponseWriter, r *http.Request) {
	// Get number of days from query parameters (default: 30)
	days := 30
	if daysParam := r.URL.Query().Get("days"); daysParam != "" {
		fmt.Sscanf(daysParam, "%d", &days)
	}

	trends, err := database.GetNetworkTrends(days)
	if err != nil {
		http.Error(w, "Failed to load network trends", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trends)
}

// handleGetDeviceUptime returns the uptime percentage of a device
func (s *Server) handleGetDeviceUptime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mac := vars["mac"]

	// Get period from query parameters (default: 7 days)
	days := 7
	if daysParam := r.URL.Query().Get("days"); daysParam != "" {
		fmt.Sscanf(daysParam, "%d", &days)
	}

	period := time.Duration(days) * 24 * time.Hour
	uptime, err := database.GetDeviceUptime(mac, period)
	if err != nil {
		http.Error(w, "Failed to calculate uptime", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"mac":    mac,
		"uptime": uptime,
		"period": days,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
