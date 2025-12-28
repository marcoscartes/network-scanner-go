# 🚀 Quick Start Guide

**Estimated time**: 5 minutes  
**Skill Level**: Beginner

---

## 📋 Prerequisites

- ✅ Windows, Linux, or macOS.
- ✅ Connection to a local network.
- ✅ A modern web browser.

---

## ⚡ Quick Start in 3 Steps

### Step 1: Download and Run

```bash
# Download the binary for your operating system and execute it:

# Windows
scanner.exe

# Linux/Mac
./scanner
```

### Step 2: Open the Dashboard

Open your browser and navigate to:
```
http://localhost:5050
```

### Step 3: Explore Your Network!

The scanner will automatically begin discovering devices on your local network.

---

## 🎯 First Steps

### 1. View Discovered Devices

The dashboard presents a table with all detected devices:
- **IP Address**: The assigned IP.
- **MAC Address**: The hardware physical address.
- **Vendor**: Manufacturer details.
- **Type**: Automatically identified category (Router, Computer, etc.).
- **Open Ports**: Detected active services.

### 2. Personalize a Device

1. Click the **Edit** (pencil) icon next to any device.
2. Add custom metadata:
   - **Name**: "My Main Router"
   - **Tags**: "critical, infrastructure"
   - **Group**: "Backbone"
3. Click **Save Changes**.

### 3. Full Port Scan

For a deep-dive analysis of a specific device:
1. Click the **+** button next to the device.
2. Confirm the scan request.
3. Monitor progress in real-time.
4. Results will update automatically specialized services are found.

### 4. Review Security Risks

If a device has potential vulnerabilities:
1. A red badge with a risk count will appear.
2. Click the badge or the **Edit** icon.
3. Review the specific findings and follow recommended fixes.

### 5. Search Filters

Use the header search bar for real-time filtering:

```
# Search by Category
type:router

# Search by Service
port:80

# Search by Organization
group:office

# Combine Filters
type:router port:80
```

---

## 🔧 Basic CLI Configuration

### Change Network Range
```bash
scanner.exe -range 10.0.0.0/24
```

### Change Scan Interval
```bash
scanner.exe -interval 300 # Every 5 minutes
```

### Change Web Dashboard Port
```bash
scanner.exe -web-port 8080
```

---

## 🎨 Dashboard Interface Highlights

### Top Bar
- **Live Indicator**: Green means your WebSocket connection for live updates is active.
- **Auto-Refresh**: Toggle and set timing for automatic table reloads.
- **Theme Toggle**: Switch between Dark and Light UI modes.

### Statistics Widgets
- **Total Devices**: Count of unique hardware seen.
- **Active Now**: Devices reachable in the last scan.
- **Security Risks**: Total vulnerabilities across your network.

---

## 💡 Quick Tips

### Key Shortcuts
- **`/`**: Instantly focus the search bar.
- **`Ctrl+F5`**: Force-reload the dashboard assets.

### Port Tooltips
Hover your cursor over any port badge to see the service name, common usage descriptions, and security status.

### Visual Cues
- 🟢 **Green Dot**: Device is currently online.
- 🔴 **Red Badge**: Security risks detected.
- 🛡️ **Shield Icon**: Device flagged as "Trusted".

---

## ⚙️ CLI Options Reference

```bash
-range                     Network range (e.g., 192.168.1.0/24)
-interval                  Scan interval in seconds (default: 60)
-web-port                  Dashboard port (default: 5050)
-db                        Database file path (default: scanner.db)
-history-retention-days    Storage duration for history (default: 90)
```

---

## 🆘 Troubleshooting

### No devices found?
- Manually specify your range: `scanner.exe -range 192.168.1.0/24`.
- Run with administrator/root privileges.

### Dashboard won't load?
- Ensure the process is running.
- Verify port 5050 is not blocked or in use by another app.

---

## 🎯 Startup Checklist

- [ ] Run the scanner binary.
- [ ] Open the web dashboard.
- [ ] Name your most important devices (e.g., Gateway, NAS).
- [ ] Run a full port scan on a server or computer.
- [ ] Check if any red security badges appear.
- [ ] Try a refined search like `type:computer vendor:apple`.
- [ ] Export a JSON backup of your initial configuration.

---

**Happy scanning! 🚀**

**Version**: 2.0.0  
**Last updated**: 2025-12-28
