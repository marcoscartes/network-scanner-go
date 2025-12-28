# 📁 Estructura del Repositorio

Esta es la estructura organizada del proyecto Network Scanner.

---

## 📂 Estructura de Carpetas

```
network-scanner-go/
│
├── 📄 README.md                    # Introducción al proyecto
├── 📄 CHANGELOG.md                 # Historial de cambios
├── 📄 QUICK_START.md               # Guía de inicio rápido
├── 📄 START_HERE.md                # Punto de entrada para nuevos usuarios
├── 📄 VERSION                      # Versión actual del proyecto
├── 📄 LICENSE                      # Licencia MIT
├── 📄 .gitignore                   # Archivos ignorados por Git
├── 📄 go.mod                       # Dependencias de Go
├── 📄 go.sum                       # Checksums de dependencias
│
├── 📁 scripts/                     # Scripts de compilación y ejecución
│   ├── README.md                   # Guía de uso de scripts
│   ├── menu.bat                    # Menú principal interactivo
│   ├── build.bat                   # Compilación simple
│   ├── build-advanced.bat          # Compilación avanzada
│   ├── run.bat                     # Ejecutar con opciones
│   └── clean.bat                   # Limpieza de archivos
│
├── 📁 planning/                    # Documentos de planificación
│   ├── README.md                   # Índice de documentos de planificación
│   ├── IMPLEMENTATION_PLAN.md      # Plan de implementación (Fases 1-5)
│   ├── PROGRESS.md                 # Seguimiento de progreso
│   ├── NEXT_STEPS.md               # Próximas funcionalidades (Fase 6+)
│   ├── ROADMAP_VISUAL.md           # Roadmap visual del proyecto
│   └── INDEX.md                    # Índice completo de planificación
│
├── 📁 docs/                        # Documentación completa
│   ├── README.md                   # Índice de documentación
│   ├── INDEX.md                    # Índice completo
│   ├── USER_GUIDE.md               # Guía del usuario
│   ├── QUICK_START_GUIDE.md        # Inicio rápido detallado
│   ├── API_REFERENCE.md            # Referencia de API REST
│   ├── ARCHITECTURE.md             # Arquitectura del sistema
│   ├── FAQ.md                      # Preguntas frecuentes
│   ├── PORT_TOOLTIPS_FEATURE.md    # Documentación de tooltips
│   ├── PORT_TOOLTIPS_CODE_REFERENCE.js
│   ├── IMPLEMENTATION_SUMMARY_PORT_TOOLTIPS.md
│   └── DOCUMENTATION_SUMMARY.md    # Resumen de documentación
│
├── 📁 cmd/                         # Aplicaciones ejecutables
│   └── scanner/
│       └── main.go                 # Punto de entrada principal
│
├── 📁 internal/                    # Código fuente interno
│   ├── database/                   # Capa de base de datos
│   │   ├── db.go
│   │   └── models.go
│   ├── scanner/                    # Motor de escaneo
│   │   ├── scanner.go
│   │   ├── port_scanner.go
│   │   └── identifier.go
│   ├── web/                        # Servidor web
│   │   ├── server.go
│   │   ├── handlers.go
│   │   ├── websocket.go
│   │   ├── static/                 # Archivos estáticos
│   │   │   ├── css/
│   │   │   └── images/
│   │   └── templates/              # Plantillas HTML
│   │       └── index.html
│   ├── notifications/              # Sistema de notificaciones
│   │   ├── manager.go
│   │   ├── detector.go
│   │   └── notifier.go
│   ├── security/                   # Escaneo de seguridad
│   │   ├── vulnerability_db.go
│   │   └── checker.go
│   ├── history/                    # Análisis histórico
│   │   └── recorder.go
│   ├── management/                 # Gestión de dispositivos
│   │   └── import_export.go
│   ├── search/                     # Búsqueda avanzada
│   │   └── query_parser.go
│   └── vendor/                     # Lookup de vendors MAC
│
├── 📁 configs/                     # Archivos de configuración
│   └── security_rules.json         # Reglas de seguridad
│
├── 📁 .github/                     # Configuración de GitHub (futuro)
│   ├── workflows/                  # GitHub Actions
│   ├── ISSUE_TEMPLATE/             # Templates de issues
│   └── PULL_REQUEST_TEMPLATE.md    # Template de PRs
│
├── 📄 scanner.exe                  # Ejecutable compilado
├── 📄 scanner.db                   # Base de datos SQLite
├── 📄 scanner.db-shm               # Shared memory (SQLite WAL)
└── 📄 scanner.db-wal               # Write-Ahead Log (SQLite WAL)
```

---

## 📚 Guía de Navegación

### Para Usuarios Nuevos

1. **Empieza aquí**: [START_HERE.md](START_HERE.md)
2. **Inicio rápido**: [QUICK_START.md](QUICK_START.md)
3. **Documentación completa**: [docs/USER_GUIDE.md](docs/USER_GUIDE.md)

### Para Desarrolladores

1. **Arquitectura**: [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)
2. **API Reference**: [docs/API_REFERENCE.md](docs/API_REFERENCE.md)
3. **Código fuente**: `cmd/` e `internal/`

### Para Compilar y Ejecutar

1. **Scripts**: [scripts/README.md](scripts/README.md)
2. **Menú interactivo**: `scripts/menu.bat`
3. **Build simple**: `scripts/build.bat`

### Para Contribuir

1. **Planificación**: [planning/](planning/)
2. **Progreso**: [planning/PROGRESS.md](planning/PROGRESS.md)
3. **Próximos pasos**: [planning/NEXT_STEPS.md](planning/NEXT_STEPS.md)

---

## 🎯 Archivos Principales en la Raíz

| Archivo | Descripción |
|---------|-------------|
| **README.md** | Introducción al proyecto |
| **CHANGELOG.md** | Historial de cambios por versión |
| **QUICK_START.md** | Guía de inicio rápido |
| **START_HERE.md** | Punto de entrada para nuevos usuarios |
| **VERSION** | Versión actual (2.0.0) |
| **go.mod** | Dependencias de Go |
| **scanner.exe** | Ejecutable compilado |
| **scanner.db** | Base de datos SQLite |

---

## 📁 Descripción de Carpetas

### `/scripts`
Scripts batch para compilación y ejecución en Windows.
- **Uso**: Facilita el desarrollo y deployment
- **Documentación**: [scripts/README.md](scripts/README.md)

### `/planning`
Documentos de planificación y seguimiento del proyecto.
- **Uso**: Roadmap, progreso, próximas features
- **Documentación**: [planning/README.md](planning/README.md)

### `/docs`
Documentación completa del proyecto.
- **Uso**: Guías de usuario, API, arquitectura
- **Documentación**: [docs/README.md](docs/README.md)

### `/cmd`
Aplicaciones ejecutables (punto de entrada).
- **Uso**: Código principal del scanner
- **Archivo**: `cmd/scanner/main.go`

### `/internal`
Código fuente interno del proyecto.
- **Uso**: Lógica de negocio, componentes
- **Paquetes**: database, scanner, web, notifications, etc.

### `/configs`
Archivos de configuración.
- **Uso**: Reglas de seguridad, configuraciones
- **Archivo**: `configs/security_rules.json`

---

## 🔄 Reorganización

Si necesitas reorganizar el repositorio, ejecuta:

```batch
reorganize.bat
```

Este script moverá automáticamente los archivos a sus carpetas correspondientes.

---

## 📊 Estadísticas del Proyecto

- **Líneas de código**: ~10,000+
- **Archivos de código**: 20+
- **Documentación**: ~60,000 palabras
- **Scripts**: 5
- **Funcionalidades**: 13+

---

## 🎓 Recursos Adicionales

- **GitHub**: [Repositorio](https://github.com/tu-usuario/network-scanner-go)
- **Issues**: [Reportar problemas](https://github.com/tu-usuario/network-scanner-go/issues)
- **Wiki**: [Documentación adicional](https://github.com/tu-usuario/network-scanner-go/wiki)

---

**Última actualización**: 2025-12-28  
**Versión**: 2.0.0
