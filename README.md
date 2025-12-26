# Network Scanner (Go)

A high-performance network scanner written in Go that discovers devices on your local network, scans ports, identifies device types, and provides a web dashboard for monitoring.

## Features

- **Network Discovery**: Automatic device discovery using ping
- **Port Scanning**: Scans common ports and supports full port scans (1-65535)
- **Device Identification**: Identifies device types based on open ports
- **MAC Vendor Lookup**: Resolves device manufacturers
- **Metrics Detection**: Finds Prometheus metrics endpoints
- **Web Dashboard**: Real-time web interface with Bootstrap 5
- **Full Port Scan**: On-demand complete port scanning with progress tracking
- **SQLite Database**: Persistent storage of discovered devices
- **Concurrent Scanning**: Fast parallel scanning using goroutines

## Installation

### Prerequisites

- Go 1.20 or higher
- Windows, Linux, or macOS

### Build

```bash
cd network-scanner-go
go mod tidy
go build -o scanner.exe cmd/scanner/main.go
```

## Usage

### Basic Usage

```bash
# Auto-detect network and start scanning
.\scanner.exe

# Specify network range
.\scanner.exe -range 192.168.1.0/24

# Custom scan interval (default: 60 seconds)
.\scanner.exe -interval 120

# Custom web port (default: 5050)
.\scanner.exe -web-port 8080

# Custom database path
.\scanner.exe -db mydevices.db
```

### Command Line Options

- `-range` - IP range to scan (e.g., 192.168.1.0/24)
- `-interval` - Scan interval in seconds (default: 60)
- `-web-port` - Web interface port (default: 5050)
- `-db` - Database file path (default: scanner.db)

## Web Dashboard

Access the dashboard at: `http://localhost:5050`

### Features

- View all discovered devices
- Sort by any column
- Click port numbers to open services
- Full port scan with progress tracking
- Real-time updates

### Full Port Scan

1. Click the **+** button next to any device
2. Confirm the scan (takes 2-5 minutes)
3. Watch real-time progress
4. View all discovered ports

## Architecture

```
network-scanner-go/
├── cmd/scanner/          # Main application
├── internal/
│   ├── database/        # SQLite operations
│   ├── scanner/         # Network scanning logic
│   ├── vendor/          # MAC vendor lookup
│   └── web/             # HTTP server & templates
└── scanner.db           # SQLite database (created on first run)
```

## Performance

- Concurrent scanning with goroutine pools
- Efficient port scanning with configurable timeouts
- Low memory footprint
- Single binary deployment

## Comparison with Python Version

**Advantages:**
- ✅ 5-10x faster scanning
- ✅ Single binary (no dependencies)
- ✅ Lower memory usage (~20MB vs ~100MB)
- ✅ Native concurrency
- ✅ Cross-platform compilation

**Same Features:**
- ✅ Network discovery
- ✅ Port scanning
- ✅ Device identification
- ✅ Web dashboard
- ✅ Full port scan with progress
- ✅ SQLite database
- ✅ Metrics detection

## 🚀 Roadmap & Development

This project is actively being enhanced with professional-grade features. Check out our development plans:

- **[QUICK_START.md](QUICK_START.md)** - Quick reference guide with checklists
- **[IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md)** - Detailed plan for Top 5 priority features
- **[NEXT_STEPS.md](NEXT_STEPS.md)** - Future enhancements roadmap

### 🎯 Top 5 Upcoming Features

1. **Notification System** - Proactive alerts for network changes
2. **Historical Analytics** - Track device behavior over time
3. **Device Management** - Tags, groups, and custom organization
4. **Enhanced Dashboard** - Real-time updates via WebSocket, dark mode
5. **Vulnerability Detection** - Security scanning and recommendations

**Total estimated development time**: 51-77 hours

See `IMPLEMENTATION_PLAN.md` for detailed breakdown and `NEXT_STEPS.md` for 20+ additional features planned.

## License

MIT License

