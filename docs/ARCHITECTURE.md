# 🏗️ System Architecture - Network Scanner

**Version**: 2.0.0  
**Last updated**: 2025-12-28

---

## 📋 Table of Contents

1. [Overview](#overview)
2. [High-Level Architecture](#high-level-architecture)
3. [Main Components](#main-components)
4. [Data Flow](#data-flow)
5. [Database](#database)
6. [Concurrency](#concurrency)
7. [Real-Time Communication](#real-time-communication)
8. [Security](#security)
9. [Design Decisions](#design-decisions)

---

## 🎯 Overview

Network Scanner is a monolithic application written in Go that combines:
- **Backend**: HTTP Server + WebSocket.
- **Scanner**: Network discovery and scanning engine.
- **Database**: SQLite with WAL (Write-Ahead Logging) mode.
- **Frontend**: HTML/CSS/JavaScript with Bootstrap 5.

### Architectural Features

- ✅ **Modular Monolith**: Easy deployment with well-separated components.
- ✅ **Native Concurrency**: Goroutines for parallel scanning.
- ✅ **Real-time**: WebSockets for instant dashboard updates.
- ✅ **Persistence**: Embedded SQLite, no external dependencies.
- ✅ **Single Binary**: Statically compiled for easy distribution.

---

## 🏛️ High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        USER                                 │
│                    (Web Browser)                            │
└────────────┬────────────────────────────────────────────────┘
             │
             │ HTTP/WebSocket
             │
┌────────────▼────────────────────────────────────────────────┐
│                    WEB SERVER (Go)                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │   HTTP       │  │  WebSocket   │  │   Static     │     │
│  │   Handlers   │  │   Hub        │  │   Files      │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
└────────────┬────────────────────────────────────────────────┘
             │
             │ Function Calls
             │
┌────────────▼────────────────────────────────────────────────┐
│                   BUSINESS LOGIC                            │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐      │
│  │ Scanner  │ │  Notif.  │ │ Security │ │  History │      │
│  │  Engine  │ │  Manager │ │  Checker │ │ Recorder │      │
│  │  └───────┘ └──────────┘ └──────────┘ └──────────┘      │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐                   │
│  │  Device  │ │  Search  │ │  Import/ │                   │
│  │  Manager │ │  Parser  │ │  Export  │                   │
│  └──────────┘ └──────────┘ └──────────┘                   │
└────────────┬────────────────────────────────────────────────┘
             │
             │ SQL Queries
             │
┌────────────▼────────────────────────────────────────────────┐
│                    DATABASE LAYER                           │
│  ┌──────────────────────────────────────────────────────┐  │
│  │              SQLite (WAL Mode)                       │  │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐            │  │
│  │  │ Devices  │ │ History  │ │  Notif.  │            │  │
│  │  └──────────┘ └──────────┘ └──────────┘            │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

---

## 🧩 Main Components

### 1. Main Application (`cmd/scanner/main.go`)

**Responsibilities**:
- System initialization.
- CLI flag configuration.
- Component orchestration.
- Main scanning loop.

**Startup Flow**:
1. Parse CLI flags.
2. Initialize database.
3. Initialize notification manager.
4. Initialize history recorder.
5. Start web server (goroutine).
104. Start scan loop (goroutine).
105. Wait for signals (graceful shutdown).

---

### 2. Scanner Engine (`internal/scanner/`)

**Files**:
- `scanner.go`: Primary scanning logic.
- `port_scanner.go`: Port scanning logic.
- `identifier.go`: Device type identification.

**Core Functions**:

#### `DiscoverDevices(network string) []Device`
- Generates a list of IPs in the specified range.
- Parallel pinging using goroutines.
- Hostname resolution.
- MAC address lookup.
- Hardware vendor (OUI) lookup.

#### `ScanCommonPorts(ip string) []int`
- Scans a predefined list of common ports.
- Configurable timeout (default: 500ms).
- Leverages concurrency with goroutines.

#### `ScanAllPorts(ip string) []int`
- Scans all ports from 1 to 65535.
- Worker pool of 100 goroutines for speed.
- Progress reporting via channels.
- Support for task cancellation.

---

### 3. Web Server (`internal/web/`)

**Files**:
- `server.go`: HTTP server configuration.
- `handlers.go`: REST API request handlers.
- `websocket.go`: WebSocket lifecycle management.
- `templates/`: HTML templates.

**Endpoints** (details in [API_REFERENCE.md](API_REFERENCE.md)):
- `/`: Main Dashboard.
- `/api/devices`: Device CRUD operations.
- `/api/scan-all-ports/:ip`: Full port scan trigger.
- `/api/notifications`: Alert management.
- `/api/stats/*`: Dashboard statistics.
- `/ws`: WebSocket connection point.

---

### 4. Database Layer (`internal/database/`)

**Responsibilities**:
- Connection management.
- CRUD operations.

**SQLite Optimization Configuration**:
- `PRAGMA journal_mode=WAL`: High concurrency support.
- `PRAGMA busy_timeout=5000`: Handle concurrent write access.
- `PRAGMA synchronous=NORMAL`: Balance speed and persistence.

---

### 5. Notification System (`internal/notifications/`)

**Workflow**:
1. Detector compares current network state vs. the previous state.
2. Changes identified (e.g., new device, offline status, port changes).
3. Notifications created in the DB.
4. Notifier delivers alerts to configured channels (Console, System, Webhook).
5. WebSocket broadcast sends alerts to active browser clients.

---

### 6. Security Scanner (`internal/security/`)

**Process**:
1. Retrieve open ports for a specific device.
2. Compare findings against a security rules database (`internal/security/vulnerability_db.go`).
3. Generate a list of potential risks.
4. Calculate a security score.
5. Store results in the database and notify for critical findings.

---

### 7. History Recorder (`internal/history/`)

**Strategy**:
- **Device Snapshots**: Captured every hour.
- **Network Stats**: Aggregate daily summaries.
- **Retention**: Default is 90 days.
- **Cleanup**: Daily automatic process.

---

## 🔄 Data Flow

### Discovery Cycle

1. **Timer**: Triggers every 60 seconds.
2. **Discovery**: Parallel pings + hardware identification.
3. **Quick Port Scan**: Checks common services.
4. **Identification**: Heuristics categorize device types.
5. **Security Check**: Vulnerability analysis.
6. **Persistence**: Database upsert.
7. **Change Detection**: Compares against prior state.
8. **History Recording**: Snapshot captured.
9. **Broadcast**: WebSocket informs clients that discovery is complete.

---

## 💾 Database Schema Summary

**Devices**: Main table for discovered devices and metadata.
**Device History**: Historical IPs, status, and service snapshots.
**Network Stats**: Daily counts of total/active devices and vulnerabilities.
**Notifications**: Persistent alert logs.

---

## ⚡ Concurrency Model

**Main Goroutines**:
- Web Server listener.
- WebSocket Hub manager.
- Periodic Scan Loop.
- History Recorder worker.
- Active Full Scans (one per requested device).

---

## 🎯 Design Decisions

### why Go?
- Native high-concurrency support.
- Compiled as a single binary for zero-dependency deployment.
- High performance for network-intensive tasks.

### Why SQLite?
- No external process management required.
- Embedded storage is lightweight.
- WAL mode provides sufficient concurrency for small-to-medium networks.

### Why WebSockets?
- Provides a "live" feel to the dashboard.
- More efficient than continuous HTTP polling.

---

## 📊 Performance Metrics

- **Ping Sweep (254 IPs)**: ~5-10 seconds.
- **Quick Port Scan**: ~1 second per device.
- **Full Port Scan (65k ports)**: ~2-5 minutes per device.
- **Memory Footprint**: ~30MB (typical base usage).

---

**Last Review**: 2025-12-28
