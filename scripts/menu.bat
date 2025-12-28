@echo off
REM ============================================
REM Network Scanner - Main Menu
REM ============================================
REM Primary script with all navigation and control options.
REM ============================================

REM Move to the project root (parent of the scripts folder)
cd /d "%~dp0.."

setlocal enabledelayedexpansion

:main_menu
cls
echo.
echo ========================================
echo   NETWORK SCANNER - MAIN MENU
echo ========================================
echo.
echo   Version: 2.0.0
echo   Project: Network Scanner (Go)
echo.
echo ========================================
echo.

REM Verify executable state
if exist scanner.exe (
    echo [OK] Executable: scanner.exe found
    for %%A in (scanner.exe) do (
        set size=%%~zA
        set /a sizeMB=!size! / 1048576
        echo [OK] Size: !sizeMB! MB
    )
) else (
    echo [!] Executable: Not compiled
)

REM Verify database
if exist scanner.db (
    for %%A in (scanner.db) do (
        set dbsize=%%~zA
        set /a dbsizeMB=!dbsize! / 1048576
        echo [OK] Database: !dbsizeMB! MB
    )
) else (
    echo [!] Database: Not found
)

echo.
echo ========================================
echo   OPTIONS
echo ========================================
echo.
echo   COMPILATION:
echo   1. Normal Build
echo   2. Optimized Build (Production)
echo   3. Advanced Build (Compilation Menu)
echo.
echo   EXECUTION:
echo   4. Run Scanner
echo   5. Run with Options
echo.
echo   MAINTENANCE:
echo   6. Clean Files
echo   7. View Logs
echo   8. Backup Database
echo.
echo   UTILITIES:
echo   9. Open Dashboard in Browser
echo   10. View Documentation
echo   11. Verify Dependencies
echo.
echo   0. Exit
echo.
echo ========================================
echo.
set /p choice="Select an option (0-11): "

if "%choice%"=="1" goto build_normal
if "%choice%"=="2" goto build_optimized
if "%choice%"=="3" goto build_advanced
if "%choice%"=="4" goto run_scanner
if "%choice%"=="5" goto run_options
if "%choice%"=="6" goto clean_files
if "%choice%"=="7" goto view_logs
if "%choice%"=="8" goto backup_db
if "%choice%"=="9" goto open_dashboard
if "%choice%"=="10" goto view_docs
if "%choice%"=="11" goto check_deps
if "%choice%"=="0" goto end

echo.
echo [ERROR] Invalid option
timeout /t 2 >nul
goto main_menu

:build_normal
cls
echo.
echo [NORMAL BUILD]
echo ========================================
echo.
call scripts\build.bat
pause
goto main_menu

:build_optimized
cls
echo.
echo [OPTIMIZED BUILD]
echo ========================================
echo.
echo [INFO] Compiling in production mode...
go mod tidy
go build -ldflags="-s -w" -o scanner.exe cmd\scanner\main.go
if %errorlevel% equ 0 (
    echo.
    echo [OK] Compilation successful!
    for %%A in (scanner.exe) do (
        set size=%%~zA
        set /a sizeMB=!size! / 1048576
        echo [OK] Size: !sizeMB! MB
    )
) else (
    echo [ERROR] Error during compilation.
)
echo.
pause
goto main_menu

:build_advanced
cls
call scripts\build-advanced.bat
goto main_menu

:run_scanner
cls
echo.
echo [RUN SCANNER]
echo ========================================
echo.
if not exist scanner.exe (
    echo [ERROR] Executable not found.
    echo [INFO] Please compile first using option 1 or 2.
    pause
    goto main_menu
)
echo [INFO] Starting Network Scanner...
echo [INFO] Dashboard: http://localhost:5050
echo.
echo Press Ctrl+C in the scanner window to stop.
echo.
start scanner.exe
timeout /t 2 >nul
echo.
echo [OK] Scanner started in the background.
echo.
pause
goto main_menu

:run_options
cls
call scripts\run.bat
goto main_menu

:clean_files
cls
call scripts\clean.bat
goto main_menu

:view_logs
cls
echo.
echo [VIEW LOGS]
echo ========================================
echo.
echo [INFO] Searching for log files...
echo.
if exist scanner.log (
    type scanner.log
) else (
    echo [INFO] No log files found.
)
echo.
pause
goto main_menu

:backup_db
cls
echo.
echo [DATABASE BACKUP]
echo ========================================
echo.
if not exist scanner.db (
    echo [ERROR] Database not found.
    pause
    goto main_menu
)

REM Create backups folder if it doesn't exist
if not exist backups mkdir backups

REM Generate backup name with date and time
for /f "tokens=2-4 delims=/ " %%a in ('date /t') do (set mydate=%%c%%b%%a)
for /f "tokens=1-2 delims=/:" %%a in ('time /t') do (set mytime=%%a%%b)
set mytime=%mytime: =0%

set backupfile=backups\scanner_backup_%mydate%_%mytime%.db

echo [INFO] Creating backup...
copy scanner.db %backupfile% >nul
if %errorlevel% equ 0 (
    echo [OK] Backup created: %backupfile%
    for %%A in (%backupfile%) do (
        set size=%%~zA
        set /a sizeMB=!size! / 1048576
        echo [OK] Size: !sizeMB! MB
    )
) else (
    echo [ERROR] Failed to create backup.
)
echo.
pause
goto main_menu

:open_dashboard
cls
echo.
echo [OPEN DASHBOARD]
echo ========================================
echo.
echo [INFO] Opening dashboard in your default browser...
start http://localhost:5050
timeout /t 2 >nul
echo [OK] Dashboard opened.
echo.
pause
goto main_menu

:view_docs
cls
echo.
echo [DOCUMENTATION]
echo ========================================
echo.
echo Available documentation:
echo.
echo 1. README.md - Introduction
echo 2. docs/USER_GUIDE.md - User Guide
echo 3. docs/API_REFERENCE.md - REST API Reference
echo 4. docs/FAQ.md - Frequently Asked Questions
echo 5. docs/QUICK_START_GUIDE.md - Quick Start
echo.
echo Opening documentation folder...
start explorer docs
timeout /t 2 >nul
echo.
pause
goto main_menu

:check_deps
cls
echo.
echo [VERIFY DEPENDENCIES]
echo ========================================
echo.
echo [INFO] Verifying Go...
where go >nul 2>nul
if %errorlevel% equ 0 (
    go version
    echo [OK] Go is correctly installed.
) else (
    echo [ERROR] Go not found.
)
echo.
echo [INFO] Verifying Go modules...
go mod verify
echo.
echo [INFO] Listing dependencies...
go list -m all
echo.
pause
goto main_menu

:end
cls
echo.
echo ========================================
echo   Thank you for using Network Scanner!
echo ========================================
echo.
echo For more information visit:
echo   - docs/README.md
echo   - docs/USER_GUIDE.md
echo.
timeout /t 3 >nul
endlocal
exit /b 0
