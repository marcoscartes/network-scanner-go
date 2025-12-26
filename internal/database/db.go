package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB
var dbMu sync.Mutex

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

// Init initializes the database
func Init(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS devices (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			mac TEXT UNIQUE NOT NULL,
			ip TEXT NOT NULL,
			vendor TEXT,
			type TEXT,
			open_ports TEXT,
			metrics_urls TEXT,
			last_seen INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS vendor_cache (
			mac TEXT PRIMARY KEY,
			vendor TEXT NOT NULL,
			cached_at INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS notifications (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT NOT NULL,
			device_ip TEXT,
			device_mac TEXT,
			message TEXT NOT NULL,
			timestamp INTEGER NOT NULL,
			read INTEGER DEFAULT 0,
			severity TEXT DEFAULT 'info'
		);

		CREATE TABLE IF NOT EXISTS notification_config (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			enabled_channels TEXT,
			email_config TEXT,
			telegram_config TEXT,
			webhook_url TEXT,
			updated_at INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS device_history (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			device_mac TEXT NOT NULL,
			ip TEXT NOT NULL,
			hostname TEXT,
			vendor TEXT,
			open_ports TEXT,
			timestamp INTEGER NOT NULL,
			change_type TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS network_stats (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date INTEGER NOT NULL UNIQUE,
			total_devices INTEGER DEFAULT 0,
			new_devices INTEGER DEFAULT 0,
			disconnected_devices INTEGER DEFAULT 0,
			total_ports INTEGER DEFAULT 0,
			active_devices INTEGER DEFAULT 0
		);

		CREATE INDEX IF NOT EXISTS idx_devices_ip ON devices(ip);
		CREATE INDEX IF NOT EXISTS idx_devices_last_seen ON devices(last_seen);
		CREATE INDEX IF NOT EXISTS idx_notifications_timestamp ON notifications(timestamp);
		CREATE INDEX IF NOT EXISTS idx_notifications_read ON notifications(read);
		CREATE INDEX IF NOT EXISTS idx_notifications_device_mac ON notifications(device_mac);
		CREATE INDEX IF NOT EXISTS idx_device_history_mac ON device_history(device_mac);
		CREATE INDEX IF NOT EXISTS idx_device_history_timestamp ON device_history(timestamp);
		CREATE INDEX IF NOT EXISTS idx_network_stats_date ON network_stats(date);
	`)

	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Database initialized successfully")
	return nil
}

// UpsertDevice inserts or updates a device
func UpsertDevice(device *Device) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	openPortsJSON, _ := json.Marshal(device.OpenPorts)
	metricsURLsJSON, _ := json.Marshal(device.MetricsURLs)

	_, err := db.Exec(`
		INSERT INTO devices (mac, ip, vendor, type, open_ports, metrics_urls, last_seen)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(mac) DO UPDATE SET
			ip = excluded.ip,
			vendor = excluded.vendor,
			type = excluded.type,
			open_ports = excluded.open_ports,
			metrics_urls = excluded.metrics_urls,
			last_seen = excluded.last_seen
	`, device.MAC, device.IP, device.Vendor, device.Type,
		string(openPortsJSON), string(metricsURLsJSON), device.LastSeen.Unix())

	return err
}

// GetAllDevices retrieves all devices from the database
func GetAllDevices() ([]*Device, error) {
	rows, err := db.Query(`
		SELECT id, mac, ip, vendor, type, open_ports, metrics_urls, last_seen
		FROM devices
		ORDER BY last_seen DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []*Device
	for rows.Next() {
		var device Device
		var openPortsJSON, metricsURLsJSON string
		var lastSeenUnix int64

		err := rows.Scan(&device.ID, &device.MAC, &device.IP, &device.Vendor,
			&device.Type, &openPortsJSON, &metricsURLsJSON, &lastSeenUnix)
		if err != nil {
			continue
		}

		json.Unmarshal([]byte(openPortsJSON), &device.OpenPorts)
		json.Unmarshal([]byte(metricsURLsJSON), &device.MetricsURLs)
		device.LastSeen = time.Unix(lastSeenUnix, 0)

		devices = append(devices, &device)
	}

	return devices, nil
}

// GetCachedVendor retrieves a cached vendor lookup
func GetCachedVendor(mac string) (string, bool) {
	var vendor string
	err := db.QueryRow("SELECT vendor FROM vendor_cache WHERE mac = ?", mac).Scan(&vendor)
	if err != nil {
		return "", false
	}
	return vendor, true
}

// SaveCachedVendor saves a vendor lookup to cache
func SaveCachedVendor(mac, vendor string) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	_, err := db.Exec(`
		INSERT OR REPLACE INTO vendor_cache (mac, vendor, cached_at)
		VALUES (?, ?, ?)
	`, mac, vendor, time.Now().Unix())
	return err
}

// Close closes the database connection
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// SaveNotification saves a notification to the database
func SaveNotification(notification *Notification) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	result, err := db.Exec(`
		INSERT INTO notifications (type, device_ip, device_mac, message, timestamp, read, severity)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, notification.Type, notification.DeviceIP, notification.DeviceMAC,
		notification.Message, notification.Timestamp.Unix(), notification.Read, notification.Severity)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	notification.ID = int(id)
	return nil
}

// GetAllNotifications retrieves all notifications
func GetAllNotifications() ([]*Notification, error) {
	rows, err := db.Query(`
		SELECT id, type, device_ip, device_mac, message, timestamp, read, severity
		FROM notifications
		ORDER BY timestamp DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		var notification Notification
		var timestampUnix int64
		var readInt int

		err := rows.Scan(&notification.ID, &notification.Type, &notification.DeviceIP,
			&notification.DeviceMAC, &notification.Message, &timestampUnix, &readInt, &notification.Severity)
		if err != nil {
			continue
		}

		notification.Timestamp = time.Unix(timestampUnix, 0)
		notification.Read = readInt == 1

		notifications = append(notifications, &notification)
	}

	return notifications, nil
}

// GetUnreadNotifications retrieves unread notifications
func GetUnreadNotifications() ([]*Notification, error) {
	rows, err := db.Query(`
		SELECT id, type, device_ip, device_mac, message, timestamp, read, severity
		FROM notifications
		WHERE read = 0
		ORDER BY timestamp DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		var notification Notification
		var timestampUnix int64
		var readInt int

		err := rows.Scan(&notification.ID, &notification.Type, &notification.DeviceIP,
			&notification.DeviceMAC, &notification.Message, &timestampUnix, &readInt, &notification.Severity)
		if err != nil {
			continue
		}

		notification.Timestamp = time.Unix(timestampUnix, 0)
		notification.Read = readInt == 1

		notifications = append(notifications, &notification)
	}

	return notifications, nil
}

// MarkNotificationAsRead marks a notification as read
func MarkNotificationAsRead(id int) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	_, err := db.Exec("UPDATE notifications SET read = 1 WHERE id = ?", id)
	return err
}

// DeleteNotification deletes a notification
func DeleteNotification(id int) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	_, err := db.Exec("DELETE FROM notifications WHERE id = ?", id)
	return err
}

// GetNotificationConfig retrieves the notification configuration
func GetNotificationConfig() (*NotificationConfig, error) {
	var config NotificationConfig
	var enabledChannelsJSON, emailConfigJSON, telegramConfigJSON string
	var updatedAtUnix int64

	err := db.QueryRow(`
		SELECT id, enabled_channels, email_config, telegram_config, webhook_url, updated_at
		FROM notification_config
		ORDER BY id DESC
		LIMIT 1
	`).Scan(&config.ID, &enabledChannelsJSON, &emailConfigJSON, &telegramConfigJSON,
		&config.WebhookURL, &updatedAtUnix)

	if err == sql.ErrNoRows {
		// Return default config if none exists
		return &NotificationConfig{
			EnabledChannels: []string{"console"},
			EmailConfig:     make(map[string]string),
			TelegramConfig:  make(map[string]string),
			WebhookURL:      "",
			UpdatedAt:       time.Now(),
		}, nil
	}

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(enabledChannelsJSON), &config.EnabledChannels)
	json.Unmarshal([]byte(emailConfigJSON), &config.EmailConfig)
	json.Unmarshal([]byte(telegramConfigJSON), &config.TelegramConfig)
	config.UpdatedAt = time.Unix(updatedAtUnix, 0)

	return &config, nil
}

// SaveNotificationConfig saves the notification configuration
func SaveNotificationConfig(config *NotificationConfig) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	enabledChannelsJSON, _ := json.Marshal(config.EnabledChannels)
	emailConfigJSON, _ := json.Marshal(config.EmailConfig)
	telegramConfigJSON, _ := json.Marshal(config.TelegramConfig)

	_, err := db.Exec(`
		INSERT INTO notification_config (enabled_channels, email_config, telegram_config, webhook_url, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, string(enabledChannelsJSON), string(emailConfigJSON), string(telegramConfigJSON),
		config.WebhookURL, time.Now().Unix())

	return err
}

// GetDeviceHistory retrieves the history of a specific device
func GetDeviceHistory(mac string, from, to time.Time) ([]*DeviceHistory, error) {
	query := `
		SELECT id, device_mac, ip, hostname, vendor, open_ports, timestamp, change_type
		FROM device_history
		WHERE device_mac = ? AND timestamp >= ? AND timestamp <= ?
		ORDER BY timestamp DESC
	`

	rows, err := db.Query(query, mac, from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []*DeviceHistory
	for rows.Next() {
		var h DeviceHistory
		var openPortsJSON string
		var timestampUnix int64
		var hostname sql.NullString

		err := rows.Scan(&h.ID, &h.DeviceMAC, &h.IP, &hostname, &h.Vendor,
			&openPortsJSON, &timestampUnix, &h.ChangeType)
		if err != nil {
			continue
		}

		if hostname.Valid {
			h.Hostname = hostname.String
		}

		json.Unmarshal([]byte(openPortsJSON), &h.OpenPorts)
		h.Timestamp = time.Unix(timestampUnix, 0)

		history = append(history, &h)
	}

	return history, nil
}

// GetNetworkTrends retrieves network statistics for the last N days
func GetNetworkTrends(days int) ([]*NetworkStats, error) {
	startDate := time.Now().AddDate(0, 0, -days)

	query := `
		SELECT id, date, total_devices, new_devices, disconnected_devices, total_ports, active_devices
		FROM network_stats
		WHERE date >= ?
		ORDER BY date ASC
	`

	rows, err := db.Query(query, startDate.Unix())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trends []*NetworkStats
	for rows.Next() {
		var stat NetworkStats
		var dateUnix int64

		err := rows.Scan(&stat.ID, &dateUnix, &stat.TotalDevices, &stat.NewDevices,
			&stat.DisconnectedDevices, &stat.TotalPorts, &stat.ActiveDevices)
		if err != nil {
			continue
		}

		stat.Date = time.Unix(dateUnix, 0)
		trends = append(trends, &stat)
	}

	return trends, nil
}

// CalculateDailyStats calculates and stores daily statistics
func CalculateDailyStats(date time.Time) (*NetworkStats, error) {
	// Normalize date to start of day
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	stats := &NetworkStats{
		Date: startOfDay,
	}

	// Get total devices seen on this day
	err := db.QueryRow(`
		SELECT COUNT(DISTINCT device_mac)
		FROM device_history
		WHERE timestamp >= ? AND timestamp < ?
	`, startOfDay.Unix(), endOfDay.Unix()).Scan(&stats.TotalDevices)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Get new devices (first appearance)
	err = db.QueryRow(`
		SELECT COUNT(DISTINCT device_mac)
		FROM device_history
		WHERE change_type = 'new'
		AND timestamp >= ? AND timestamp < ?
	`, startOfDay.Unix(), endOfDay.Unix()).Scan(&stats.NewDevices)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Get disconnected devices
	err = db.QueryRow(`
		SELECT COUNT(DISTINCT device_mac)
		FROM device_history
		WHERE change_type = 'disconnect'
		AND timestamp >= ? AND timestamp < ?
	`, startOfDay.Unix(), endOfDay.Unix()).Scan(&stats.DisconnectedDevices)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Get currently active devices
	devices, err := GetAllDevices()
	if err == nil {
		stats.ActiveDevices = len(devices)
		totalPorts := 0
		for _, device := range devices {
			totalPorts += len(device.OpenPorts)
		}
		stats.TotalPorts = totalPorts
	}

	// Save stats to database
	_, err = db.Exec(`
		INSERT OR REPLACE INTO network_stats 
		(date, total_devices, new_devices, disconnected_devices, total_ports, active_devices)
		VALUES (?, ?, ?, ?, ?, ?)
	`, startOfDay.Unix(), stats.TotalDevices, stats.NewDevices,
		stats.DisconnectedDevices, stats.TotalPorts, stats.ActiveDevices)

	if err != nil {
		log.Printf("Failed to save daily stats: %v", err)
		return nil, err
	}

	return stats, nil
}

// GetDeviceUptime calculates the uptime percentage of a device over a period
func GetDeviceUptime(mac string, period time.Duration) (float64, error) {
	startTime := time.Now().Add(-period)

	// Get all history records for this device in the period
	history, err := GetDeviceHistory(mac, startTime, time.Now())
	if err != nil {
		return 0, err
	}

	if len(history) == 0 {
		return 0, nil
	}

	// Count how many times the device was seen (not disconnected)
	activeCount := 0
	totalCount := len(history)

	for _, record := range history {
		if record.ChangeType != "disconnect" {
			activeCount++
		}
	}

	uptime := (float64(activeCount) / float64(totalCount)) * 100.0
	return uptime, nil
}
