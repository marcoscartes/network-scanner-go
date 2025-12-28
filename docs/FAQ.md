# ❓ Frequently Asked Questions (FAQ)

**Version**: 2.0.0  
**Last updated**: 2025-12-28

---

## 📋 Table of Contents

1. [General](#general)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Scanning](#scanning)
5. [Dashboard](#dashboard)
6. [Notifications](#notifications)
7. [Security](#security)
8. [Performance](#performance)
9. [Common Issues](#common-issues)
10. [Advanced](#advanced)

---

## 🌐 General

### What is Network Scanner?

Network Scanner is a professional network discovery and monitoring tool written in Go. it allows you to discover devices, scan ports, detect vulnerabilities, and receive notifications about changes in your local network.

### Is it free?

Yes, Network Scanner is completely free and open source under the MIT license.

### How does it differ from other tools?

- ✅ **Single binary**: No dependencies required for execution.
- ✅ **Modern Dashboard**: Intuitive web interface.
- ✅ **Real-time**: Instant updates via WebSockets.
- ✅ **Advanced Management**: Tags, groups, and advanced search.
- ✅ **Integrated Security**: Automatic vulnerability detection.
- ✅ **History**: Temporal analysis of your network activity.

### Which operating systems are supported?

- ✅ Windows 10/11
- ✅ Linux (Ubuntu, Debian, CentOS, etc.)
- ✅ macOS

### Do I need administrator privileges?

For basic functionality, no. However, some specific features require elevated permissions:
- **Ping (ICMP)**: May require admin/root on some systems to send raw packets.
- **MAC Resolution**: Requires access to the ARP table.
- **Ports < 1024**: Requires root/admin privileges to scan on Linux/Mac.

---

## 💻 Installation

### How do I install Network Scanner?

**Option 1: Pre-compiled Binary**
1. Download the binary for your OS.
2. Place it in a folder.
3. Run it from the terminal.

**Option 2: Build from Source**
```bash
git clone https://github.com/your-user/network-scanner-go.git
cd network-scanner-go
go mod tidy
go build -o scanner.exe cmd/scanner/main.go
```

### What dependencies do I need?

**To run the binary**: None.

**To build from source**:
- Go 1.20 or higher.
- Git (to clone the repository).

### Where is the data stored?

By default, in `scanner.db` within the same directory as the executable. You can change this location using the `-db` flag:

```bash
./scanner.exe -db /custom/path/scanner.db
```

### Can I run it on a headless server?

Yes, the scanner works perfectly on headless servers. You just need to access the dashboard from another computer on the network:

```bash
# On the server
./scanner.exe -web-port 5050 -range 192.168.1.0/24

# From another computer
http://SERVER_IP:5050
```

---

## 🎯 Usage

### How do I start the scanner?

```bash
# Basic usage (auto-detects network)
./scanner.exe

# Specify network range
./scanner.exe -range 192.168.1.0/24

# Change web port
./scanner.exe -web-port 8080
```

### How do I access the dashboard?

Open your browser and go to: `http://localhost:5050`

### How do I stop the scanner?

Press `Ctrl+C` in the terminal. The scanner will close gracefully, saving all data.

### Can I run it in the background?

**Linux/Mac**:
```bash
nohup ./scanner &
```

**Windows**:
- Use a Windows Service wrapper.
- Or run it via Task Scheduler.

### How do I change the scan interval?

Use the `-interval` flag (in seconds):

```bash
# Scan every 2 minutes
./scanner.exe -interval 120

# Scan every 5 minutes
./scanner.exe -interval 300
```

---

## 🔍 Scanning

### Which ports are scanned by default?

Commonly used ports including:
- **Web**: 80, 443, 8080, 8443
- **SSH/Telnet**: 22, 23
- **FTP**: 21
- **Email**: 25, 110, 143, 465, 587, 993, 995
- **Databases**: 3306, 5432, 27017, 6379
- **Windows Services**: 135, 139, 445, 3389
- And many more...

### How do I scan all ports?

1. In the dashboard, click the **+** button next to the device.
2. Confirm the scan request.
3. Wait 2-5 minutes.
4. The results will be updated automatically.

### Why does the full scan take so long?

Scanning 65,535 ports takes time. The scanner uses:
- **100 parallel workers** for speed.
- **500ms timeouts** for reliability.
- **Result**: Approx. 2-5 minutes per device.

### Can I scan multiple networks?

Currently, only one network range per instance. To scan multiple networks:
- Run multiple instances of the scanner.
- Use different web ports and database files for each.

```bash
# Network 1
./scanner.exe -range 192.168.1.0/24 -web-port 5050 -db net1.db

# Network 2
./scanner.exe -range 192.168.2.0/24 -web-port 5051 -db net2.db
```

### Does the scan affect my network performance?

The impact is minimal:
- **Ping**: Small ICMP packets.
- **Port scan**: Brief TCP connections.
- **Bandwidth**: < 1 Mbps during active scanning.
- **Devices**: Should not be adversely affected.

---

## 🖥️ Dashboard

### How do I switch between dark and light themes?

Click the sun/moon icon in the top header. Your preference is saved in your browser's local storage.

### How do I search for a specific device?

Use the search bar:
- **By IP**: `192.168.1.100`
- **By Name**: `router`
- **By Type**: `type:router`
- **By Port**: `port:80`
- **By Group**: `group:office`

See the [Advanced Search](USER_GUIDE.md#advanced-search) section in the User Guide for more syntax.

### How do I organize my devices?

1. Click the **Edit** (pencil) icon next to the device.
2. You can add:
   - **Custom Name**: "Main Router"
   - **Tags**: "critical, network"
   - **Group**: "Infrastructure"
   - **Notes**: Additional information.

### Are changes lost after a restart?

No, all changes are saved in the SQLite database and persist across restarts.

### Can I export my data?

Yes:
1. Click on "Data" → "Export JSON".
2. A JSON file containing all devices and metadata will be downloaded.

---

## 🔔 Notifications

### What types of notifications will I receive?

- **new_device** (Info): New device detected on the network.
- **device_offline** (Warning): A previously active device is no longer reachable.
- **port_change** (Warning): Changes detected in open ports.
- **vulnerability_detected** (Critical): A new security risk has been identified.

### How can I view notifications?

Click the bell icon (🔔) in the top header.

### Why am I not receiving notifications?

Possible causes:
1. **No changes**: Notifications only trigger when something changes.
2. **Rate limiting**: Maximum 1 notification every 30s per device.
3. **WebSocket disconnected**: Check the "Live" indicator (it should be green).

---

## 🔒 Security

### What vulnerabilities does it detect?

- **Insecure Services**: Telnet, unencrypted FTP.
- **Dangerous Ports**: Exposed SMB, RDP.
- **Exposed Databases**: MySQL, PostgreSQL, MongoDB accessible from the network.
- **Weak Configurations**: Services that typically lack authentication.

### How do I fix a vulnerability?

1. Click the red vulnerability badge in the dashboard.
2. A modal will open with details.
3. Read the **Description** of the issue.
4. Follow the recommended **Solution**.
5. Click **Learn More** for official documentation.

### Does the scanner attempt to exploit vulnerabilities?

**No**. The scanner only:
- Detects open ports.
- Compares findings against a database of known security risks.
- Reports potential issues.

It does not perform any attacks or exploitations.

### Is it safe to use the scanner?

Yes, the scanner is safe:
- ✅ **Read-only**: It does not modify device configurations.
- ✅ **Local-first**: It does not send your data to the internet.
- ⚠️ **Note**: The dashboard does not currently feature authentication. Anyone on your local network can access it.

### How can I protect the dashboard?

Since there is no built-in auth yet, you could use:
1. **Firewall**: Block port 5050 except for trusted IPs.
2. **VPN**: Access it only via a secure VPN connection.
3. **Reverse Proxy**: Use Nginx or Apache with Basic Auth.

---

## ⚡ Performance

### How many devices can it handle?

- **Optimal**: 1-100 devices.
- **Good**: 100-500 devices.
- **Possible**: 500-1000 devices.

### Does it consume many resources?

No:
- **CPU**: < 5% idle, ~20% during active scanning.
- **RAM**: ~20-50 MB.
- **Disk**: 5-50 MB (database size).

---

## 🔧 Common Issues

### It's not finding any devices

**Solutions**:
1. Check the network range: `./scanner.exe -range 192.168.1.0/24`
2. Run as administrator/root.
3. Ensure no firewalls are blocking ICMP or ARP.
4. Try scanning a known specific IP first to verify connectivity.

### Dashboard doesn't load

**Checks**:
1. Is the scanner process running?
2. Is port 5050 already in use?
3. Try a different port: `./scanner.exe -web-port 8080`

### Error: "database is locked"

**Solution**:
- Close other connections to the database.
- Restart the scanner.
- (The system uses WAL mode to minimize this, but it can still happen with concurrent external access).

---

## 📞 More Help

- **[User Guide](USER_GUIDE.md)**: Complete manual.
- **[API Reference](API_REFERENCE.md)**: API documentation.
- **[Architecture](ARCHITECTURE.md)**: System design details.

---

**Last update**: 2025-12-28
