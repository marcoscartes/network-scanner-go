# 🔧 Build and Execution Scripts

This folder contains several batch scripts (.bat) to facilitate the compilation and execution of the Network Scanner.

---

## 📋 Available Scripts

### 🎯 **menu.bat** - Main Menu (RECOMMENDED)

Primary script with all options presented in an interactive menu.

**Usage**:
```batch
menu.bat
```

**Options**:
- Compilation (Normal, Optimized, Advanced)
- Execution (Normal, with Options)
- Maintenance (Clean, Logs, Backup)
- Utilities (Dashboard, Docs, Dependencies)

---

### 🔨 **build.bat** - Simple Build

Compiles the project quickly and easily.

**Usage**:
```batch
build.bat
```

**Features**:
- ✅ Verifies Go installation
- ✅ Automatically downloads dependencies
- ✅ Compiles the executable
- ✅ Displays result information
- ✅ Error handling

**Result**: Generates `scanner.exe` (~20 MB)

---

### ⚙️ **build-advanced.bat** - Advanced Build

A menu with multiple compilation options.

**Usage**:
```batch
build-advanced.bat
```

**Options**:
1. **Normal Build**: Standard compilation for development.
2. **Optimized Build**: Reduces size (-ldflags="-s -w").
3. **Debug Build**: Includes debug symbols.
4. **Clean**: Removes compiled files.
5. **Build + Run**: Automatically compiles and launches.

**When to use each**:
- **Normal**: Daily development.
- **Optimized**: User distribution (reduces size by ~30%).
- **Debug**: Debugging with tools like Delve.

---

### ▶️ **run.bat** - Run Scanner

Executes the scanner with different configurations.

**Usage**:
```batch
run.bat
```

**Options**:
1. **Normal Mode**: Auto-detects the network.
2. **Custom Range**: Specify range (e.g., 192.168.1.0/24).
3. **Custom Port**: Change web port (e.g., 8080).
4. **Debug Mode**: Short scan interval (30s).
5. **View Help**: Shows all available flags.

---

### 🧹 **clean.bat** - Cleanup

Cleans up compiled and temporary files.

**Usage**:
```batch
clean.bat
```

**Deletes**:
- `scanner.exe` (executable)
- `scanner.db` (database - optional)
- `scanner.db-shm` and `scanner.db-wal` (WAL files)
- Go cache (optional)

**⚠️ Warning**: Deleting the database is irreversible. Create a backup first.

---

## 🚀 Quick Start Guide

### First Time

1. **Build**:
   ```batch
   build.bat
   ```

2. **Run**:
   ```batch
   scanner.exe
   ```

3. **Open Dashboard**:
   ```
   http://localhost:5050
   ```

### Daily Use

**Option 1: Main Menu (Recommended)**
```batch
menu.bat
```

**Option 2: Individual Scripts**
```batch
# Re-build
build.bat

# Run
run.bat
```

---

## 📊 Build Comparison

| Type | Size | Speed | Debug | Use Case |
|------|--------|-----------|-------|-----|
| **Normal** | ~20 MB | Fast | ❌ | Development |
| **Optimized** | ~14 MB | Very Fast | ❌ | Production |
| **Debug** | ~25 MB | Slow | ✅ | Debugging |

---

## 🔧 Command Line Options

Once compiled, you can run the scanner with various options:

```batch
# View all options
scanner.exe -help

# Specify network range
scanner.exe -range 192.168.1.0/24

# Change scan interval (seconds)
scanner.exe -interval 120

# Change web port
scanner.exe -web-port 8080

# Change database location
scanner.exe -db C:\data\scanner.db

# Change history retention (days)
scanner.exe -history-retention-days 30

# Combine options
scanner.exe -range 192.168.1.0/24 -interval 120 -web-port 8080
```

---

## 🛠️ Troubleshooting

### Error: "Go is not installed"
**Solution**:
1. Install Go from: https://golang.org/dl/
2. Verify installation: `go version`
3. Ensure Go is in your PATH.

### Error: "Error downloading dependencies"
**Solution**:
```batch
# Clear Go mod cache
go clean -modcache

# Try again
build.bat
```

### Scanner not finding devices
**Solution**:
- Run as administrator (Right-click `run.bat` → "Run as administrator").
- Manually specify the range: `scanner.exe -range 192.168.1.0/24`.

---

## 📝 Important Notes

### Permissions
Some systems may require administrator privileges for:
- Sending pings (ICMP).
- Accessing the ARP table.
- Scanning privileged ports (< 1024).

### Antivirus
Some antivirus software may flag the scanner as suspicious due to its network scanning activities.
**Solution**: Add an exception for `scanner.exe`.

---

## 🎓 Additional Resources

- **[README.md](../README.md)** - Project introduction.
- **[docs/USER_GUIDE.md](../docs/USER_GUIDE.md)** - Complete user manual.
- **[docs/FAQ.md](../docs/FAQ.md)** - Frequently asked questions.

---

**Last updated**: 2025-12-28  
**Version**: 2.0.0
