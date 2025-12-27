# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.5.0] - 2025-12-27

### Added
- **Vulnerability Intelligence**: Added actionable "Fix" instructions and "Learn More" external links for all detected vulnerabilities.
- **Interactive Security UI**: Vulnerability badges are now clickable, opening a detailed modal with remediation steps.
- **Improved Code Structure**: Separated CSS into a standalone file (`internal/web/static/css/style.css`) for better maintainability and performance.

### Changed
- **Enhanced Icons**: Added bi-icons for various UI elements and improved contrast in dark mode.
- **Better UX**: Replaced inline styles with a structured CSS class system.
- **Security Knowledge Base**: Enriched `configs/security_rules.json` with detailed severity descriptions and remediation URLs.

### Fixed
- Fixed broken HTML tags in device tag rendering.
- Improved responsiveness of Chart.js elements.

## [0.4.0] - 2025-12-26

### Added
- **Device Management System**: Complete functionality to manage device details (Custom Name, Type, Tags, Notes, Trusted status).
- **Interactive UI**: New Edit Device modal and improved dashboard table displaying custom metadata.
- **Dashboard Controls**: Added navbar controls for Auto-Refresh toggle and interval selection (persisted key).
- **Smart Refresh**: Prevents page reload while user is interacting with modals.
- **Notification Management**: Added "Mark All Read" & "Clear All" buttons; automatic cleanup of old notifications.
- **CLI Flags**: Added `-notification-retention` flag (default: 7 days).

### Changed
- **Database**: Automatic schema migration for new columns (`custom_name`, `tags`, etc.).
- **API**: New endpoints for `PUT /api/devices/{mac}` and bulk notification actions.

### Fixed
- Fixed UI refresh logic to be less intrusive during user interaction.

## [0.3.0] - 2025-12-26

### Added
- **Dashboard**: Integrated Chart.js for visualizing network trends and device distribution.
- **Scanner**: Automated historical snapshot recording (hourly) and daily statistics aggregation.
- **UI**: Added "Network Activity" line chart and "Device Types" doughnut chart.
- **Integration**: Linked scanner main loop with history recording module.

### Changed
- Optimized historical data collection to prevent database bloat (snapshots limited to hourly).
- Updated `PROGRESS.md` to reflect Phase 2 completion.

## [0.2.1] - 2025-12-26

### Added
- **Branding**: Added a new professional application icon.

### Changed
- Updated dashboard HTML to include favicon link.

## [0.2.0] - 2025-12-26

### Added
- Historical data models (DeviceHistory, NetworkStats)
- History recording system
- Statistics analyzer with uptime calculations
- API endpoints for device history and network statistics:
  - `GET /api/history/device/{mac}` - Device history
  - `GET /api/history/network` - Network history
  - `GET /api/stats/overview` - Statistics overview
  - `GET /api/stats/trends` - Network trends
  - `GET /api/stats/uptime/{mac}` - Device uptime
- Database functions for historical data queries
- Network growth analysis
- Most active devices tracking

### Changed
- Extended database schema with history tables and indices

## [0.1.0] - 2025-12-19

### Added
- Complete notification system
- Multi-channel notifications (Console, System, Webhook)
- Network change detection (new devices, disconnections, port changes)
- Notification management API:
  - `GET /api/notifications` - List all notifications
  - `POST /api/notifications/{id}/read` - Mark as read
  - `DELETE /api/notifications/{id}` - Delete notification
  - `GET /api/notifications/config` - Get configuration
  - `PUT /api/notifications/config` - Update configuration
- Modern dashboard with notification sidebar
- Real-time notification updates
- Rate limiting to prevent notification spam
- Notification persistence in database

### Changed
- Enhanced database schema with notifications tables
- Updated web dashboard with notification UI

## [0.0.1] - 2025-12-03

### Added
- Initial Go implementation of network scanner
- Device discovery on local network
- Port scanning capabilities
- MAC address vendor lookup
- Device type identification
- Web dashboard for visualization
- SQLite database for device storage
- API endpoints:
  - `POST /api/scan-all-ports/{ip}` - Full port scan
  - `GET /api/scan-progress/{ip}` - Scan progress
- Automatic network range detection
- Concurrent scanning for better performance

### Features
- Network device discovery
- Port scanning (quick and full)
- Vendor identification via MAC address
- Device type detection
- Web-based dashboard
- RESTful API
- Persistent storage

[Unreleased]: https://github.com/marcoscartes/network-scanner-go/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/marcoscartes/network-scanner-go/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/marcoscartes/network-scanner-go/compare/v0.2.1...v0.3.0
[0.2.1]: https://github.com/marcoscartes/network-scanner-go/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/marcoscartes/network-scanner-go/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/gigas/network-scanner-go/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/gigas/network-scanner-go/releases/tag/v0.0.1
