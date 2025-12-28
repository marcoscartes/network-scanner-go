@echo off
REM ============================================
REM Network Scanner - Reorganize Repository
REM ============================================
REM Este script reorganiza el repositorio
REM moviendo archivos a carpetas apropiadas
REM ============================================

echo.
echo ========================================
echo   Reorganizacion del Repositorio
echo ========================================
echo.
echo Este script movera archivos a carpetas
echo para una mejor organizacion.
echo.
echo Estructura propuesta:
echo   /scripts       - Scripts batch
echo   /planning      - Documentos de planificacion
echo   /docs          - Documentacion (ya existe)
echo   /cmd           - Codigo fuente (ya existe)
echo   /internal      - Codigo fuente (ya existe)
echo   /configs       - Configuraciones (ya existe)
echo.
set /p confirm="Continuar? (s/n): "
if /i not "%confirm%"=="s" (
    echo Operacion cancelada
    pause
    exit /b 0
)

echo.
echo [INFO] Iniciando reorganizacion...
echo.

REM ============================================
REM 1. Crear carpetas necesarias
REM ============================================

echo [1/4] Creando carpetas...

if not exist "scripts" mkdir scripts
echo [OK] Carpeta scripts/ creada

if not exist "planning" mkdir planning
echo [OK] Carpeta planning/ creada

echo.

REM ============================================
REM 2. Mover scripts batch
REM ============================================

echo [2/4] Moviendo scripts batch...

if exist "build.bat" (
    move "build.bat" "scripts\build.bat" >nul
    echo [OK] build.bat movido a scripts/
)

if exist "build-advanced.bat" (
    move "build-advanced.bat" "scripts\build-advanced.bat" >nul
    echo [OK] build-advanced.bat movido a scripts/
)

if exist "run.bat" (
    move "run.bat" "scripts\run.bat" >nul
    echo [OK] run.bat movido a scripts/
)

if exist "clean.bat" (
    move "clean.bat" "scripts\clean.bat" >nul
    echo [OK] clean.bat movido a scripts/
)

if exist "menu.bat" (
    move "menu.bat" "scripts\menu.bat" >nul
    echo [OK] menu.bat movido a scripts/
)

if exist "SCRIPTS_README.md" (
    move "SCRIPTS_README.md" "scripts\README.md" >nul
    echo [OK] SCRIPTS_README.md movido a scripts/README.md
)

echo.

REM ============================================
REM 3. Mover documentos de planificacion
REM ============================================

echo [3/4] Moviendo documentos de planificacion...

if exist "IMPLEMENTATION_PLAN.md" (
    move "IMPLEMENTATION_PLAN.md" "planning\IMPLEMENTATION_PLAN.md" >nul
    echo [OK] IMPLEMENTATION_PLAN.md movido a planning/
)

if exist "PROGRESS.md" (
    move "PROGRESS.md" "planning\PROGRESS.md" >nul
    echo [OK] PROGRESS.md movido a planning/
)

if exist "NEXT_STEPS.md" (
    move "NEXT_STEPS.md" "planning\NEXT_STEPS.md" >nul
    echo [OK] NEXT_STEPS.md movido a planning/
)

if exist "ROADMAP_VISUAL.md" (
    move "ROADMAP_VISUAL.md" "planning\ROADMAP_VISUAL.md" >nul
    echo [OK] ROADMAP_VISUAL.md movido a planning/
)

if exist "INDEX.md" (
    move "INDEX.md" "planning\INDEX.md" >nul
    echo [OK] INDEX.md movido a planning/
)

echo.

REM ============================================
REM 4. Crear README en planning
REM ============================================

echo [4/4] Creando README en planning...

echo # Planning Documents > "planning\README.md"
echo. >> "planning\README.md"
echo This folder contains project planning and progress tracking documents. >> "planning\README.md"
echo. >> "planning\README.md"
echo - **IMPLEMENTATION_PLAN.md** - Detailed implementation plan for the top 5 features >> "planning\README.md"
echo - **PROGRESS.md** - Progress tracking for all phases >> "planning\README.md"
echo - **NEXT_STEPS.md** - Future enhancements roadmap >> "planning\README.md"
echo - **ROADMAP_VISUAL.md** - Visual roadmap of the project >> "planning\README.md"
echo - **INDEX.md** - Index of all planning documents >> "planning\README.md"

echo [OK] README.md creado en planning/

echo.

REM ============================================
REM 5. Resumen
REM ============================================

echo ========================================
echo   Reorganizacion Completada!
echo ========================================
echo.
echo Estructura actual:
echo.
echo   network-scanner-go/
echo   ├── README.md              (Principal)
echo   ├── CHANGELOG.md           (Historial de cambios)
echo   ├── QUICK_START.md         (Inicio rapido)
echo   ├── START_HERE.md          (Punto de entrada)
echo   ├── VERSION                (Version del proyecto)
echo   ├── .gitignore             (Git ignore)
echo   ├── go.mod / go.sum        (Dependencias Go)
echo   │
echo   ├── /scripts               (Scripts batch)
echo   │   ├── README.md
echo   │   ├── menu.bat
echo   │   ├── build.bat
echo   │   ├── build-advanced.bat
echo   │   ├── run.bat
echo   │   └── clean.bat
echo   │
echo   ├── /planning              (Planificacion)
echo   │   ├── README.md
echo   │   ├── IMPLEMENTATION_PLAN.md
echo   │   ├── PROGRESS.md
echo   │   ├── NEXT_STEPS.md
echo   │   ├── ROADMAP_VISUAL.md
echo   │   └── INDEX.md
echo   │
echo   ├── /docs                  (Documentacion)
echo   ├── /cmd                   (Codigo fuente)
echo   ├── /internal              (Codigo fuente)
echo   └── /configs               (Configuraciones)
echo.
echo ========================================
echo.
echo NOTA: Los archivos de base de datos (scanner.db*)
echo y el ejecutable (scanner.exe) permanecen en la raiz
echo para facilitar el uso.
echo.
echo ========================================
echo.

pause
