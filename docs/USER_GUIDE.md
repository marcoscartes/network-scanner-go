# 📖 User Guide - Network Scanner

**Version**: 2.0.0  
**Last updated**: 2025-12-28

---

## 📋 Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Quick Start](#quick-start)
4. [Main Features](#main-features)
5. [Web Dashboard](#web-dashboard)
6. [Device Management](#device-management)
7. [Advanced Search](#advanced-search)
8. [Notifications](#notifications)
9. [Security](#security)
10. [Historical Analysis](#historical-analysis)
11. [Import/Export Data](#importexport-data)
12. [Configuration](#configuration)
13. [Troubleshooting](#troubleshooting)

---

## 🎯 Introduction

Network Scanner is a professional network discovery and monitoring tool written in Go. It allows you to:

- ✅ Automatically discover all devices on your network
- ✅ Scan ports and detect services
- ✅ Identify device types
- ✅ Detect security vulnerabilities
- ✅ Receive notifications for network changes
- ✅ Visualize historical statistics
- ✅ Organize devices with tags and groups

---

## 💻 Installation

### Requirements
- **Operating System**: Windows, Linux, or macOS
- **Go**: Version 1.20 or higher (only required for building from source)
- **Permissions**: Administrator/root privileges for certain features (e.g., ICMP pings, ARP access)

### Pre-compiled Binary
1. Download the binary for your specific operating system.
2. Place it in a folder of your choice.
3. You're ready to go!

### Building from Source

```bash
# Clone the repository
git clone https://github.com/your-username/network-scanner-go.git
cd network-scanner-go

# Install dependencies
go mod tidy

# Build
go build -o scanner.exe cmd/scanner/main.go
```

---

## 🚀 Quick Start

### Basic Usage

```bash
# Run with automatic configuration
./scanner.exe

# Specify network range
./scanner.exe -range 192.168.1.0/24

# Change web server port
./scanner.exe -web-port 8080

# Change scan interval (in seconds)
./scanner.exe -interval 120
```

### Accessing the Dashboard

1. Launch the scanner.
2. Open your web browser.
3. Navigate to: `http://localhost:5050`
4. Start exploring your network!

---

## 🎯 Main Features

### 1. Network Discovery

The scanner automatically handles:
- Detecting your local network range.
- Sending pings to all potential IPs.
- Identifying active devices.
- Resolving hostnames.
- Retrieving MAC addresses and identifying manufacturers.

**Frequency**: Every 60 seconds by default (configurable).

### 2. Port Scanning

**Quick Scan** (Automatic):
- Scans common ports (22, 80, 443, 3306, etc.).
- Runs during every discovery cycle.
- Identifies well-known services.

**Full Scan** (Manual):
- Scans all 65,535 ports.
- Takes approximately 2-5 minutes per device.
- Displays real-time progress.
- Detects services on non-standard ports.

### 3. Device Identification

The system automatically categorizes:
- **Computers**: Ports 22, 3389, 445.
- **Routers**: Ports 80, 443, 23.
- **Printers**: Ports 631, 9100.
- **IoT Devices**: Based on specific port patterns.
- **Servers**: Multiple service signatures.
- **Mobile Devices**: Specific port signatures.

### 4. Vulnerability Detection

The scanner identifies:
- ✅ Insecure services (Telnet, unencrypted FTP).
- ✅ Critical open ports.
- ✅ Weak configurations.
- ✅ Outdated services.

For each finding, it provides:
- A description of the issue.
- Severity levels (Info, Warning, Critical).
- A recommended solution.
- Links to official documentation.

---

## 🖥️ Web Dashboard

### Main Interface

The dashboard includes:

#### Top Navigation Bar
- **Logo & Title**: Quick link to the main view.
- **Live Indicator**: WebSocket connection status.
- **Interval Selector**: Configure auto-refresh frequency.
- **Data Menu**: Export and import options.
- **Notifications**: Alert bell with a counter badge.
- **Theme Toggle**: Switch between dark and light modes.

#### Stats Widgets
- **Total Devices**: Total unique devices discovered.
- **Active Now**: Devices currently online.
- **New (24h)**: Devices discovered within the last 24 hours.
- **Security Risks**: Count of detected vulnerabilities.

#### Search Bar
- Real-time results filtering.
- Advanced syntax support.
- Keyboard shortcut: `/`.

#### Device Table

Columns:
1. **Status**: Green indicator for online devices.
2. **Device/IP**: Custom name or IP address.
3. **MAC Address**: Physical address of the hardware.
4. **Vendor**: Hardware manufacturer.
5. **Type & Tags**: Category and custom labels.
6. **Group**: Assigned organization group.
7. **Security Risks**: Red badge for vulnerabilities.
8. **Open Ports**: Clickable service badges.
9. **Metrics**: Detected Prometheus endpoints.
10. **Last Seen**: Timestamp of the last successful scan.
11. **Actions**: Management buttons.

#### Charts

**Network Activity (30 days)**:
- Visualizes total (blue) vs active (green) device trends.

**Device Distribution**:
- Doughnut chart showing your network composition by device type.

---

## 🏷️ Device Management

### Editing a Device

1. Click the **Edit** (pencil) icon in the actions column.
2. The management modal will appear.

**Editable Fields**:
- **Name**: Assign a friendly name (e.g., "Main Kitchen Printer").
- **Type**: Override the automatic classification.
- **Tags**: Comma-separated labels for custom sorting.
- **Group**: Organize into logical groups (e.g., "Guest Network").
- **Known/Trusted**: Flag devices you trust.
- **Notes**: Extra information or context.

### Full Port Scanning

1. Click the **+** button next to a device.
2. Confirm the scan request (2-5 minutes duration).
3. Monitor progress via the dynamic progress bar showing current port and findings.
4. Results update automatically in the table upon completion.

### Reviewing Vulnerabilities

1. Open the device's management modal.
2. Security risks will appear in the "Security" section.
3. Click **Rescan** to verify fixes.
4. Each risk includes severity, description, solution, and a "Learn More" link.

---

## 🔍 Advanced Search

### Basic Syntax

```
# Simple text search
router

# Search by IP
192.168.1.1

# Search by MAC
aa:bb:cc:dd:ee:ff
```

### Advanced Filters

```
# By Device Type
type:router
type:computer

# By Open Port
port:80
port:22

# By Group
group:office
group:iot

# By Tag
tag:personal
tag:critical

# By Trusted Status
known:true
known:false

# By Manufacturer
vendor:apple
vendor:cisco

# Combined Filters (implicit AND)
type:router port:80
group:office known:true
```

---

## 🔔 Notifications

### Alert Types
- **new_device**: Notify when a new IP appears.
- **device_offline**: Notify when a previously seen device disappears.
- **port_change**: Notify when services are added or removed.
- **vulnerability_detected**: Notify when a security risk is found.

### Management
- Access the panel via the **Bell** icon in the header.
- Features include: Mark as Read, Delete, Mark All as Read, and Clear All.
- All notifications are persisted in the database.

---

## 🔒 Security Rules

The system scans for:
- **Insecure Communication**: Telnet (23), FTP (21).
- **Dangerous Access**: SMB (445), RDP (3389).
- **Public Databases**: Unauthorized access to MySQL, Postgres, MongoDB, Redis.

### Severity Levels
- **🔵 Info**: General information.
- **🟡 Warning**: Configuration review recommended.
- **🔴 Critical**: Immediate intervention required.

---

## 📊 Historical Analysis

### Data Retention
- System stores snapshots every hour and daily summaries.
- Default retention is **90 days** (configurable via `-history-retention-days`).
- Automatic cleanup runs daily.

---

## 📤 Import/Export Data

### Exporting
1. Click **Data** in the top bar.
2. Select **Export JSON**.
3. You'll receive a timestamped JSON file with all device metadata.

### Importing
1. Click **Data** → **Import JSON**.
2. Upload your file and confirm.
3. Existing devices (matched by MAC) will be updated with your custom names, tags, and notes.

---

## ⚙️ CLI Configuration Summary

```bash
-range                     Network range (e.g., 10.0.0.0/24)
-interval                  Scan interval in seconds (default: 60)
-web-port                  Dashboard port (default: 5050)
-db                        Database path (default: scanner.db)
-history-retention-days    History storage duration (default: 90)
```

---

## 🔧 Troubleshooting

### No devices found
- Verify you are scanning the correct subnet.
- Try running with elevated privileges (`sudo` or as Administrator).
- Check if your firewall is blocking ICMP or ARP.

### Dashboard won't load
- Ensure the scanner process is active.
- Verify port 5050 is not used by another application.

---

## 🎓 Next Steps

1. **Brand your network**: Start by naming your known devices.
2. **Review Security**: Look for red badges and follow the remediation steps.
3. **Monitor Trends**: Check the charts after a few days of operation.

---

**Happy scanning! 🚀**
