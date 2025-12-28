@echo off
REM ============================================
REM Network Scanner - Build Script
REM ============================================
REM This script compiles the Network Scanner
REM and generates the scanner.exe executable.
REM ============================================

REM Move to the project root (parent of the scripts folder)
cd /d "%~dp0.."

echo.
echo ========================================
echo   Network Scanner - Build Script
echo ========================================
echo.

REM Verify Go is installed
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Go is not installed or not in your PATH.
    echo.
    echo Please install Go from: https://golang.org/dl/
    echo.
    pause
    exit /b 1
)

REM Show Go version
echo [INFO] Verifying Go version...
go version
echo.

REM Clean existing executable if it exists
if exist scanner.exe (
    echo [INFO] Checking for running scanner instance...
    tasklist /FI "IMAGENAME eq scanner.exe" 2>nul | find /I "scanner.exe" >nul
    if !errorlevel! equ 0 (
        echo [INFO] Found running scanner instance. Terminating...
        taskkill /F /IM scanner.exe /T >nul 2>nul
        timeout /t 1 >nul
    )

    echo [INFO] Removing existing executable...
    del /F /Q scanner.exe 2>nul
    if exist scanner.exe (
        echo [ERROR] Could not remove scanner.exe. Please close the application and try again.
        pause
        exit /b 1
    )
    echo.
)

REM Download dependencies
echo [INFO] Tying up dependencies...
go mod tidy
if %errorlevel% neq 0 (
    echo.
    echo [ERROR] Failed to download dependencies.
    pause
    exit /b 1
)
echo.

REM Compile the project
echo [INFO] Compiling Network Scanner...
echo [INFO] This might take a few seconds...
echo.

go build -o scanner.exe cmd\scanner\main.go

if %errorlevel% neq 0 (
    echo.
    echo [ERROR] Error during compilation.
    echo.
    pause
    exit /b 1
)

REM Verify the executable was created
if not exist scanner.exe (
    echo.
    echo [ERROR] Executable was not created correctly.
    pause
    exit /b 1
)

REM Show executable information
echo.
echo ========================================
echo   Build Successful!
echo ========================================
echo.
echo [OK] Executable created: scanner.exe

REM Get file size
for %%A in (scanner.exe) do (
    set size=%%~zA
)

REM Convert bytes to MB (Note: batch math is integer only)
set /a sizeMB=%size% / 1048576

echo [OK] Size: %sizeMB% MB
echo.
echo ========================================
echo.
echo To run the scanner:
echo   scanner.exe
echo.
echo To see options:
echo   scanner.exe -help
echo.
echo To access the dashboard:
echo   http://localhost:5050
echo.
echo ========================================
echo.

pause
