@echo off
REM ============================================
REM Network Scanner - Clean Script
REM ============================================
REM Cleans up compiled and temporary files.
REM ============================================

REM Move to the project root (parent of the scripts folder)
cd /d "%~dp0.."

echo.
echo ========================================
echo   Network Scanner - Clean Script
echo ========================================
echo.

set cleaned=0

REM Clean executable
if exist scanner.exe (
    echo [INFO] Checking for running scanner instance...
    tasklist /FI "IMAGENAME eq scanner.exe" 2>nul | find /I "scanner.exe" >nul
    if !errorlevel! equ 0 (
        echo [INFO] Found running scanner instance. Terminating...
        taskkill /F /IM scanner.exe /T >nul 2>nul
        timeout /t 1 >nul
    )
    echo [INFO] Removing scanner.exe...
    del /F /Q scanner.exe 2>nul
    set /a cleaned+=1
    if exist scanner.exe (
        echo [ERROR] Could not remove scanner.exe.
    ) else (
        echo [OK] scanner.exe removed
    )
)

REM Clean database files
echo.
set /p cleandb="Delete database and all data? (y/n): "
if /i "%cleandb%"=="y" (
    if exist scanner.db (
        echo [INFO] Removing database...
        del scanner.db
        set /a cleaned+=1
        echo [OK] scanner.db removed
    )
    if exist scanner.db-shm (
        del scanner.db-shm
        echo [OK] scanner.db-shm removed
    )
    if exist scanner.db-wal (
        del scanner.db-wal
        echo [OK] scanner.db-wal removed
    )
)

REM Clean Go cache
echo.
set /p cleancache="Clean Go cache? (y/n): "
if /i "%cleancache%"=="y" (
    echo [INFO] Cleaning Go cache...
    go clean -cache
    go clean -modcache
    set /a cleaned+=1
    echo [OK] Go cache cleaned
)

REM Summary
echo.
echo ========================================
if %cleaned% gtr 0 (
    echo [OK] Cleanup completed.
) else (
    echo [INFO] No files were found to clean.
)
echo ========================================
echo.

pause
