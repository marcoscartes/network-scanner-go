@echo off
REM ============================================
REM Network Scanner - Advanced Build Script
REM ============================================
REM Advanced script with various compilation options.
REM ============================================

REM Move to the project root (parent of the scripts folder)
cd /d "%~dp0.."

setlocal enabledelayedexpansion

echo.
echo ========================================
echo   Network Scanner - Advanced Build
echo ========================================
echo.

REM Verify Go is installed
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Go is not installed or not in your PATH.
    echo.
    pause
    exit /b 1
)

REM Show options menu
:menu
echo.
echo Select an option:
echo.
echo 1. Normal Build (Development)
echo 2. Optimized Build (Production)
echo 3. Debug Build (with debug info)
echo 4. Clean Compiled Files
echo 5. Build + Run
echo 6. Exit
echo.
set /p choice="Option (1-6): "

if "%choice%"=="1" goto build_normal
if "%choice%"=="2" goto build_optimized
if "%choice%"=="3" goto build_debug
if "%choice%"=="4" goto clean
if "%choice%"=="5" goto build_and_run
if "%choice%"=="6" goto end

echo [ERROR] Invalid option
goto menu

:build_normal
echo.
echo [INFO] Compiling in normal mode...
echo.
go mod tidy
go build -o scanner.exe cmd\scanner\main.go
if %errorlevel% equ 0 (
    echo [OK] Compilation successful!
    call :show_info
) else (
    echo [ERROR] Error during compilation.
)
goto menu

:build_optimized
echo.
echo [INFO] Compiling in optimized mode (Production)...
echo [INFO] This reduces binary size and improves performance.
echo.
go mod tidy
go build -ldflags="-s -w" -o scanner.exe cmd\scanner\main.go
if %errorlevel% equ 0 (
    echo [OK] Optimized compilation successful!
    call :show_info
    echo [INFO] Optimized executable (no debug symbols).
) else (
    echo [ERROR] Error during compilation.
)
goto menu

:build_debug
echo.
echo [INFO] Compiling with debug information...
echo.
go mod tidy
go build -gcflags="all=-N -l" -o scanner.exe cmd\scanner\main.go
if %errorlevel% equ 0 (
    echo [OK] Debug compilation successful!
    call :show_info
    echo [INFO] Executable includes debug symbols (for debugging).
) else (
    echo [ERROR] Error during compilation.
)
goto menu

:clean
echo.
echo [INFO] Cleaning compiled files...
if exist scanner.exe (
    echo [INFO] Checking for running scanner instance...
    tasklist /FI "IMAGENAME eq scanner.exe" 2>nul | find /I "scanner.exe" >nul
    if !errorlevel! equ 0 (
        echo [INFO] Found running scanner instance. Terminating...
        taskkill /F /IM scanner.exe /T >nul 2>nul
        timeout /t 1 >nul
    )
    del /F /Q scanner.exe 2>nul
    if exist scanner.exe (
        echo [ERROR] Could not remove scanner.exe.
    ) else (
        echo [OK] scanner.exe removed
    )
)
if exist scanner.db (
    set /p confirm="Do you want to delete the database (scanner.db)? (y/n): "
    if /i "!confirm!"=="y" (
        del scanner.db
        if exist scanner.db-shm del scanner.db-shm
        if exist scanner.db-wal del scanner.db-wal
        echo [OK] Database deleted
    )
)
echo [OK] Cleanup completed
goto menu

:build_and_run
echo.
echo [INFO] Compiling and running...
echo.
go mod tidy
go build -o scanner.exe cmd\scanner\main.go
if %errorlevel% equ 0 (
    echo [OK] Compilation successful!
    echo.
    echo [INFO] Starting scanner...
    echo.
    echo ========================================
    echo.
    start scanner.exe
    timeout /t 2 >nul
    echo [INFO] Scanner started in the background.
    echo [INFO] Dashboard: http://localhost:5050
    echo.
    echo Press Ctrl+C in the scanner window to stop it.
    echo.
) else (
    echo [ERROR] Error during compilation.
)
goto menu

:show_info
if exist scanner.exe (
    echo.
    for %%A in (scanner.exe) do (
        set size=%%~zA
        set /a sizeMB=!size! / 1048576
        echo [INFO] Size: !sizeMB! MB ^(!size! bytes^)
    )
    echo [INFO] Location: %cd%\scanner.exe
)
exit /b 0

:end
echo.
echo Exiting...
echo.
endlocal
exit /b 0
