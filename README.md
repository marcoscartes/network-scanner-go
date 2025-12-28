# Network Scanner (Go)

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/Version-2.0.0-blue.svg?style=flat-square)](https://github.com/marcoscartes/network-scanner-go/releases)
[![AI Generated](https://img.shields.io/badge/AI%20Generated-100%25-blueviolet?style=flat-square&logo=openai)](https://openai.com/)

> [!IMPORTANT]
> **100% AI-Generated Code**: Every line of code, documentation, and logic in this repository was generated and architected by an Artificial Intelligence agent (Antigravity).

A high-performance network scanner written in Go that discovers devices on your local network, scans ports, identifies device types, and provides a web dashboard for monitoring.

## ✨ Features

- **Network Discovery**: Automatic device discovery using ping
- **Port Scanning**: Scans common ports and supports full port scans (1-65535)
- **Device Identification**: Identifies device types based on open ports
- **MAC Vendor Lookup**: Resolves device manufacturers
- **Metrics Detection**: Finds Prometheus metrics endpoints
- **Security Auditing**: Detects vulnerable services with remediation guides and official documentation links
- **Port Information Tooltips**: Interactive tooltips showing detailed information about 80+ common ports
- **Web Dashboard**: Modern real-time interface with dark/light mode and WebSocket updates
- **Historical Trends**: Tracks device uptime and network growth with interactive charts
- **Device Management**: Advanced organization with custom naming, tags, and groups
- **Data Portability**: Full JSON import/export for device metadata
- **Full Port Scan**: On-demand complete port scanning with real-time progress tracking
- **SQLite Database**: Core persistent storage with WAL mode for high concurrency
- **Concurrent Scanning**: Fast parallel scanning using goroutine pools

---

## 🚀 Quick Start

### 1. Build

```bash
# Using build script (Windows)
scripts\build.bat

# Or manually
go mod tidy
go build -o scanner.exe cmd/scanner/main.go
```

### 2. Run

```bash
# Auto-detect network and start scanning
.\scanner.exe

# Or use the interactive menu
scripts\menu.bat
```

### 3. Access Dashboard

Open your browser: `http://localhost:5050`

**For detailed instructions**, see [QUICK_START.md](QUICK_START.md)

---

## 📚 Documentation

### For Users
- **[Quick Start Guide](QUICK_START.md)** - Get started in 5 minutes
- **[User Guide](docs/USER_GUIDE.md)** - Complete user manual
- **[FAQ](docs/FAQ.md)** - Frequently asked questions (54+)

### For Developers
- **[Architecture](docs/ARCHITECTURE.md)** - System design and components
- **[API Reference](docs/API_REFERENCE.md)** - REST API documentation
- **[Documentation Index](docs/INDEX.md)** - All documentation

### Build & Run
- **[Scripts Guide](scripts/README.md)** - Build and run scripts
- **[Repository Structure](REPOSITORY_STRUCTURE.md)** - Project organization

### Planning
- **[Implementation Plan](planning/IMPLEMENTATION_PLAN.md)** - Completed features (Phases 1-5)
- **[Progress Tracking](planning/PROGRESS.md)** - Development progress
- **[Next Steps](planning/NEXT_STEPS.md)** - Future enhancements (20+ features)
- **[Visual Roadmap](planning/ROADMAP_VISUAL.md)** - Project roadmap

---

## 💻 Installation

### Prerequisites

- Go 1.20 or higher
- Windows, Linux, or macOS

### Build Options

**Option 1: Using Scripts (Recommended for Windows)**
```bash
# Interactive menu with all options
scripts\menu.bat

# Quick build
scripts\build.bat

# Advanced build (optimized, debug, etc.)
scripts\build-advanced.bat
```

**Option 2: Manual Build**
```bash
cd network-scanner-go
go mod tidy
go build -o scanner.exe cmd/scanner/main.go
```

**Option 3: Optimized Build (Production)**
```bash
go build -ldflags="-s -w" -o scanner.exe cmd/scanner/main.go
```

---

## 🎮 Usage

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
- `-history-retention-days` - Days to keep history (default: 90)

### Using Scripts

```bash
# Interactive menu with all options
scripts\menu.bat

# Run with custom options
scripts\run.bat

# Clean compiled files
scripts\clean.bat
```

---

## 🖥️ Web Dashboard

Access the dashboard at: `http://localhost:5050`

### Features

- View all discovered devices
- Sort by any column
- Click port numbers to open services
- **Hover over ports** to see detailed service information
- Full port scan with progress tracking
- **Interactive Security**: Click vulnerability badges to see "How to Fix"
- **Historical Charts**: Network activity and device distribution
- **Real-time**: Instant updates via WebSockets and Live indicators
- **Search**: Advanced search with filters (type:, port:, group:, tag:)
- **Manage**: Custom names, tags, groups, and notes
- **Export/Import**: Backup and restore device data

### Full Port Scan

1. Click the **+** button next to any device
2. Confirm the scan (takes 2-5 minutes)
3. Watch real-time progress
4. View all discovered ports

---

## 📁 Project Structure

```
network-scanner-go/
├── README.md                   # This file
├── QUICK_START.md              # Quick start guide
├── CHANGELOG.md                # Version history
│
├── scripts/                    # Build and run scripts
│   ├── menu.bat                # Interactive menu
│   ├── build.bat               # Simple build
│   ├── build-advanced.bat      # Advanced build options
│   ├── run.bat                 # Run with options
│   └── clean.bat               # Clean files
│
├── planning/                   # Project planning
│   ├── IMPLEMENTATION_PLAN.md  # Completed features
│   ├── PROGRESS.md             # Progress tracking
│   ├── NEXT_STEPS.md           # Future features
│   └── ROADMAP_VISUAL.md       # Visual roadmap
│
├── docs/                       # Documentation
│   ├── USER_GUIDE.md           # Complete user guide
│   ├── API_REFERENCE.md        # API documentation
│   ├── ARCHITECTURE.md         # System architecture
│   └── FAQ.md                  # Frequently asked questions
│
├── cmd/scanner/                # Main application
├── internal/                   # Internal packages
│   ├── database/               # SQLite operations
│   ├── scanner/                # Network scanning
│   ├── web/                    # HTTP server
│   ├── notifications/          # Notification system
│   ├── security/               # Security scanning
│   ├── history/                # Historical analytics
│   └── ...
│
├── configs/                    # Configuration files
└── scanner.db                  # SQLite database (created on first run)
```

For detailed structure, see [REPOSITORY_STRUCTURE.md](REPOSITORY_STRUCTURE.md)

---

## ⚡ Performance

- Concurrent scanning with goroutine pools
- Efficient port scanning with configurable timeouts
- Low memory footprint (~20-50 MB)
- Single binary deployment
- SQLite with WAL mode for high concurrency

---

## 🆚 Comparison with Python Version

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

---

## 🎯 Development Status

### ✅ Completed Features (Phases 1-5)

1. **Notification System** - ✅ Complete
2. **Historical Analytics** - ✅ Complete
3. **Device Management** - ✅ Complete
4. **Enhanced Dashboard** - ✅ Complete
5. **Vulnerability Detection** - ✅ Complete

**Total development time**: 56 hours

### 🔮 Upcoming Features (Phase 6+)

See [planning/NEXT_STEPS.md](planning/NEXT_STEPS.md) for 20+ planned features including:
- Authentication & Authorization
- Email notifications
- Prometheus/Grafana integration
- Network topology mapping
- And much more...

---

## 🤝 Contributing

Contributions are welcome! Please see:
- [planning/IMPLEMENTATION_PLAN.md](planning/IMPLEMENTATION_PLAN.md) - Completed work
- [planning/NEXT_STEPS.md](planning/NEXT_STEPS.md) - Future features
- [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) - System design

---

## 📄 License

MIT License - See [LICENSE](LICENSE) for details

---

## 🆘 Support

- **Documentation**: [docs/](docs/)
- **FAQ**: [docs/FAQ.md](docs/FAQ.md)
- **Issues**: [GitHub Issues](https://github.com/your-username/network-scanner-go/issues)

---

## 🌟 Acknowledgments

Built with Go, SQLite, Bootstrap, and Chart.js

**Version**: 2.0.0  
**Last Updated**: 2025-12-28
