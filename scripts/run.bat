@echo off
REM ============================================
REM Network Scanner - Run Script
REM ============================================
REM Executes the scanner with common options.
REM ============================================

REM Move to the project root (parent of the scripts folder)
cd /d "%~dp0.."

echo.
echo ========================================
echo   Network Scanner - Run Script
echo ========================================
echo.

REM Verify executable exists
if not exist scanner.exe (
    echo [ERROR] scanner.exe not found.
    echo.
    echo Please compile first using build.bat.
    echo.
    pause
    exit /b 1
)

REM Show options menu
echo Select how to run the scanner:
echo.
echo 1. Normal Mode (Auto-detect network)
echo 2. Specify Network Range
echo 3. Change Web Port
echo 4. Debug Mode (Short interval)
echo 5. Show Help
echo 6. Exit
echo.
set /p choice="Option (1-6): "

if "%choice%"=="1" goto run_normal
if "%choice%"=="2" goto run_custom_range
if "%choice%"=="3" goto run_custom_port
if "%choice%"=="4" goto run_debug
if "%choice%"=="5" goto show_help
if "%choice%"=="6" goto end

echo [ERROR] Invalid option.
pause
exit /b 1

:run_normal
echo.
echo [INFO] Starting scanner in normal mode...
echo [INFO] Dashboard: http://localhost:5050
echo.
echo Press Ctrl+C to stop.
echo.
scanner.exe
goto end

:run_custom_range
echo.
set /p range="Enter network range (e.g., 192.168.1.0/24): "
echo.
echo [INFO] Starting scanner with range: %range%
echo [INFO] Dashboard: http://localhost:5050
echo.
echo Press Ctrl+C to stop.
echo.
scanner.exe -range %range%
goto end

:run_custom_port
echo.
set /p port="Enter web port (e.g., 8080): "
echo.
echo [INFO] Starting scanner on port: %port%
echo [INFO] Dashboard: http://localhost:%port%
echo.
echo Press Ctrl+C to stop.
echo.
scanner.exe -web-port %port%
goto end

:run_debug
echo.
echo [INFO] Starting scanner in debug mode...
echo [INFO] Scan interval: 30 seconds
echo [INFO] Dashboard: http://localhost:5050
echo.
echo Press Ctrl+C to stop.
echo.
scanner.exe -interval 30
goto end

:show_help
echo.
scanner.exe -help
echo.
pause
goto end

:end
exit /b 0
