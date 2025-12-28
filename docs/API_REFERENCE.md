# 🔌 API Reference - Network Scanner

**Version**: 2.0.0  
**Base URL**: `http://localhost:5050`  
**Last updated**: 2025-12-28

---

## 📋 Table of Contents

1. [Introduction](#introduction)
2. [Authentication](#authentication)
3. [Device Endpoints](#device-endpoints)
4. [Scan Endpoints](#scan-endpoints)
5. [Notification Endpoints](#notification-endpoints)
6. [Statistics Endpoints](#statistics-endpoints)
7. [History Endpoints](#history-endpoints)
8. [Management Endpoints](#management-endpoints)
9. [WebSocket](#websocket)
10. [Error Codes](#error-codes)
11. [Usage Examples](#usage-examples)

---

## 🎯 Introduction

The Network Scanner REST API provides programmatic access to all system features. 

### Features
- **Format**: JSON
- **Protocol**: HTTP/1.1
- **Real-time**: WebSocket updates for live events.
- **CORS**: Enabled for development.

### Response Conventions

**Success Response**:
```json
{
  "status": "success",
  "data": { ... }
}
```

**Error Response**:
```json
{
  "status": "error",
  "error": "Description of the error"
}
```

---

## 🔐 Authentication

**Current Status**: No authentication required.

**Forward-looking**: JWT (JSON Web Tokens) will be implemented in Phase 6.

---

## 📱 Device Endpoints

### GET /api/devices

Retrieves a list of all discovered devices.

**Query Parameters**:
- `q` (string, optional): Advanced search query strings.

**Example**:
```bash
curl http://localhost:5050/api/devices
curl http://localhost:5050/api/devices?q=type:router
```

---

### GET /api/devices/:mac

Retrieves detailed information for a specific device.

**Parameters**:
- `mac` (string): The hardware MAC address of the device.

**Example**:
```bash
curl http://localhost:5050/api/devices/aa:bb:cc:dd:ee:ff
```

---

### PUT /api/devices/:mac

Updates metadata for a specific device.

**Body**:
```json
{
  "custom_name": "My Router",
  "custom_type": "Gateway",
  "is_known": true,
  "tags": ["critical", "lan"],
  "group_name": "Infrastructure",
  "notes": "Main entry point"
}
```

---

## 🔍 Scan Endpoints

### POST /api/scan-all-ports/:ip

Triggers a full port scan (1-65535) for a specified IP.

**Example**:
```bash
curl -X POST http://localhost:5050/api/scan-all-ports/192.168.1.100
```

---

### GET /api/scan-progress/:ip

Retrieves the current progress of an active full port scan.

**Example**:
```bash
curl http://localhost:5050/api/scan-progress/192.168.1.100
```

---

## 🔔 Notification Endpoints

### GET /api/notifications

Retrieves all recorded notifications.

**Query Parameters**:
- `unread` (boolean, optional): Filter for only unread notifications.

---

### POST /api/notifications/:id/read

Marks a specific notification as read.

---

### POST /api/notifications/read-all

Marks every existing notification as read.

---

## 📊 Statistics Endpoints

### GET /api/stats/overview

Provides general network health and summary stats.

---

### GET /api/stats/trends

Retrieves historical network trends.

**Query Parameters**:
- `days` (int, optional): The number of days to look back (default: 30).

---

## 📜 History Endpoints

### GET /api/history/device/:mac

Retrieves the scan history for a specific hardware device.

---

## 📦 Management Endpoints

### GET /api/export

Exports all device metadata in JSON format.

---

### POST /api/import

Imports device metadata from a compatible JSON file. Requires a `multipart/form-data` file upload with the field name `file`.

---

## 🔌 WebSocket

**URL**: `ws://localhost:5050/ws`

The WebSocket connection broadcasts events in real-time. Message types include:
- `scan_progress`: Updates during lengthy full port scans.
- `scan_complete`: Triggered when a full scan finishes.
- `discovery_complete`: Sent after each background network discovery pass.
- `notification`: Broadcasts a new system alert.

---

## 📝 Usage Examples

### Example: Start Scan and Monitor Progress

```bash
# Start the full port scan
curl -X POST http://localhost:5050/api/scan-all-ports/192.168.1.5

# Monitor progress via CLI (every 5s)
watch -n 5 curl http://localhost:5050/api/scan-progress/192.168.1.5
```

### Example: Bulk Update via Export/Import

```bash
# 1. Export current data
curl http://localhost:5050/api/export > backup.json

# 2. Modify backup.json (e.g., mass tag update)

# 3. Re-import updated file
curl -X POST -F "file=@backup.json" http://localhost:5050/api/import
```

---

**Last update**: 2025-12-28
