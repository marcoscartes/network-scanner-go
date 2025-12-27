# 🚀 Network Scanner - Plan de Implementación Top 5

**Fecha de inicio**: 2025-12-19  
**Objetivo**: Implementar las 5 funcionalidades prioritarias para mejorar el scanner de red

---

## 📋 Resumen Ejecutivo

Este plan detalla la implementación de las 5 funcionalidades más impactantes:

1. **Sistema de Notificaciones** - Alertas proactivas de cambios en la red
2. **Historial y Estadísticas** - Análisis temporal de datos
3. **Gestión de Dispositivos** - Organización y personalización
4. **Dashboard Mejorado** - Mejor UX con búsqueda y tiempo real
5. **Detección de Vulnerabilidades** - Seguridad básica

---

## 🎯 Fase 1: Sistema de Notificaciones (Prioridad 1)

### Objetivos
- Detectar y notificar cambios en la red
- Soportar múltiples canales de notificación
- Sistema configurable y extensible

### Tareas

#### 1.1 Estructura de Datos (1-2 horas)
- [x] Crear modelo `Notification` en `internal/database/models.go`
- [x] Crear modelo `NotificationConfig` para configuración
- [x] Añadir tabla `notifications` a la base de datos
- [x] Añadir tabla `notification_config` a la base de datos

#### 1.2 Detector de Cambios (2-3 horas)
- [x] Crear `internal/notifications/detector.go`
- [x] Función `DetectNewDevice(device *Device) bool`
- [x] Función `DetectDisconnectedDevice(device *Device) bool`
- [x] Función `DetectPortChanges(old, new *Device) []int`
- [x] Función `CompareDeviceStates(old, new []*Device) []Change`

#### 1.3 Sistema de Notificaciones (3-4 horas)
- [x] Crear `internal/notifications/notifier.go`
- [x] Interface `Notifier` con método `Send(notification *Notification) error`
- [x] Implementar `ConsoleNotifier` (logs)
- [x] Implementar `SystemNotifier` (notificaciones del SO - Windows toast)
- [x] Implementar `WebhookNotifier` (POST HTTP)
- [x] Crear `internal/notifications/manager.go`

#### 1.4 Integración con Scanner (1-2 horas)
- [x] Modificar `cmd/scanner/main.go`
- [x] Inicializar NotificationManager
- [x] Comparar estado anterior vs nuevo después de cada scan
- [x] Enviar notificaciones de cambios detectados
- [x] Añadir flags de configuración

#### 1.5 API y Dashboard (2-3 horas)
- [x] Añadir endpoints en `internal/web/server.go`
- [x] `GET /api/notifications` - Listar notificaciones
- [x] `POST /api/notifications/{id}/read` - Marcar como leída
- [x] `DELETE /api/notifications/{id}` - Eliminar
- [x] `GET /api/notifications/config` - Obtener configuración
- [x] `PUT /api/notifications/config` - Actualizar configuración
- [x] Actualizar dashboard
- [x] Badge con contador de notificaciones no leídas
- [x] Panel lateral con lista de notificaciones
- [x] Configuración de notificaciones

**Tiempo estimado**: 9-14 horas  
**Archivos nuevos**: 3 (`detector.go`, `notifier.go`, `manager.go`)  
**Archivos modificados**: 4 (`main.go`, `server.go`, `models.go`, `db.go`)

---

## 📊 Fase 2: Historial y Estadísticas (Prioridad 2)

### Objetivos
- Registrar histórico de todos los cambios
- Generar estadísticas útiles
- Visualizar datos temporales

### Tareas

#### 2.1 Modelo de Datos Históricos (2-3 horas)
- [x] Crear modelo `DeviceHistory` en `internal/database/models.go`
- [x] Crear modelo `NetworkStats` para estadísticas agregadas
- [x] Añadir tabla `device_history` a la base de datos
- [x] Añadir tabla `network_stats` a la base de datos
- [x] Crear índices para consultas rápidas por fecha y MAC

#### 2.2 Sistema de Registro Histórico (2-3 horas)
- [x] Crear `internal/history/recorder.go`
- [x] Función `RecordDeviceState(device *Device, changeType string)`
- [x] Función `RecordNetworkSnapshot(devices []*Device)`
- [x] Función `CalculateDailyStats(date time.Time) *NetworkStats`

#### 2.3 Integración con Scanner (1-2 horas)
- [x] Modificar `cmd/scanner/main.go`
- [x] Registrar snapshot después de cada scan
- [x] Calcular y guardar estadísticas diarias
- [x] Añadir tarea programada para limpieza de datos antiguos

#### 2.4 API de Estadísticas (2-3 horas)
- [x] Añadir endpoints en `internal/web/server.go`
- [x] `GET /api/history/device/{mac}` - Historial de un dispositivo
- [x] `GET /api/history/network` - Historial de la red
- [x] `GET /api/stats/overview` - Estadísticas generales
- [x] `GET /api/stats/trends` - Tendencias temporales
- [x] `GET /api/stats/uptime/{mac}` - Uptime de dispositivo

#### 2.5 Visualización en Dashboard (3-4 horas)
- [x] Crear sección de estadísticas en el dashboard
- [x] Gráfica de dispositivos activos en el tiempo (Chart.js)
- [x] Top 10 dispositivos más activos
- [x] Estadísticas de uptime
- [x] Timeline de eventos importantes
- [x] Añadir vista detallada por dispositivo (En modal de edición)

**Tiempo estimado**: 10-15 horas  
**Archivos nuevos**: 2 (`recorder.go`, `analyzer.go`)  
**Archivos modificados**: 4 (`main.go`, `server.go`, `models.go`, `db.go`)  
**Dependencias nuevas**: Chart.js (frontend)

---

## 🏷️ Fase 3: Gestión de Dispositivos (Prioridad 3)

### Objetivos
- Permitir personalización de dispositivos
- Organizar dispositivos en grupos
- Añadir contexto útil

### Tareas

#### 3.1 Modelo de Datos Extendido (1-2 horas)
- [x] Extender modelo `Device` en `internal/database/models.go`
- [x] Añadir campo `CustomName` (string)
- [x] Añadir campo `Notes` (string)
- [x] Añadir campo `Tags` ([]string - JSON)
- [x] Añadir campo `Group` (string)
- [x] Añadir campo `IsKnown` (bool)
- [x] Migrar base de datos para añadir nuevas columnas

#### 3.2 Lógica de Gestión (2-3 horas)
- [x] Implementar `UpdateDeviceDetails` en DB package
- [x] Soporte para metadatos personalizados
- [x] Sistema de tags funcional

#### 3.3 API de Gestión (2-3 horas)
- [x] Añadir endpoints en `internal/web/server.go`
- [x] `PUT /api/devices/{mac}` - Actualizar dispositivo completo

#### 3.4 UI de Gestión (3-4 horas)
- [x] Actualizar dashboard con funcionalidades de gestión
- [x] Modal de edición de dispositivo (nombre, notas, tags, is_known)
- [x] Badges para tags en la tabla

**Tiempo estimado**: 9-14 horas  
**Archivos nuevos**: 3 (`device_manager.go`, `groups.go`, `import_export.go`)  
**Archivos modificados**: 3 (`server.go`, `models.go`, `db.go`)

---

## 🎨 Fase 4: Dashboard Mejorado (Prioridad 4)

### Objetivos
- Mejorar UX significativamente
- Añadir actualización en tiempo real
- Búsqueda y filtrado avanzado

### Tareas

#### 4.1 WebSockets para Tiempo Real (3-4 horas)
- [x] Añadir dependencia `github.com/gorilla/websocket`
- [x] Crear `internal/web/websocket.go`
- [x] Endpoint `GET /ws` para conexiones WebSocket
- [x] Actualizar frontend para usar WebSocket

#### 4.2 Búsqueda y Filtrado Avanzado (2-3 horas)
- [x] Crear `internal/search/query_parser.go`
- [x] Autocompletado de tags y campos especiales
- [x] Filtros rápidos en UI

#### 4.3 Mejoras Visuales (3-4 horas)
- [x] Implementar tema oscuro con toggle
- [x] Mejorar diseño de la tabla
- [x] Añadir indicadores visuales (Estado online, badges, iconos)
- [x] Animaciones suaves

#### 4.4 Dashboard Widgets (2-3 horas)
- [x] Secciones de widgets informativos
- [x] Total de dispositivos, activos, conocidos, riesgos de seguridad

#### 4.5 Responsive Design (1-2 horas)
- [x] Optimizar para móviles y tablets

**Tiempo estimado**: 11-16 horas  
**Archivos nuevos**: 2 (`websocket.go`, `query_parser.go`)  
**Archivos modificados**: 2 (`server.go`, `index.html`)  
**Dependencias nuevas**: `gorilla/websocket`, Chart.js, posiblemente DataTables.js

---

## 🔒 Fase 5: Detección de Vulnerabilidades (Prioridad 5)

### Objetivos
- Identificar puertos peligrosos
- Detectar servicios desactualizados
- Generar recomendaciones de seguridad

### Tareas

#### 5.1 Base de Conocimiento (2-3 horas)
- [x] Crear `internal/security/vulnerability_db.go`
- [x] Cargar reglas desde `configs/security_rules.json`
- [x] Niveles de severidad implementados

#### 5.2 Motor de Análisis (3-4 horas)
- [x] Implementar `CheckDevice()` para vulnerabilidades
- [x] Sistema de scoring basado en severidad
- [x] Integración en el loop principal de escaneo

#### 5.3 Integración con Scanner (1-2 horas)
- [x] Guardar hallazgos en la base de datos (como JSON en tabla devices)
- [x] Generar notificaciones para hallazgos críticos/altos

#### 5.5 Dashboard de Seguridad (3-4 horas)
- [x] Mostrar riesgos en la tabla de dispositivos (Badge dinámico)
- [x] Health Score visual en la parte superior
- [x] Widget de "Security Risks" (reemplaza Unknown Devices)
- [x] Botón "Rescan" en modal de dispositivo

**Tiempo estimado**: 12-18 horas  
**Archivos nuevos**: 3 (`vulnerability_db.go`, `scanner.go`, `reporter.go`)  
**Archivos modificados**: 4 (`main.go`, `server.go`, `models.go`, `db.go`)  
**Archivos de configuración**: 1 (`security_rules.json`)

---

## 📅 Cronograma Estimado

| Fase | Funcionalidad | Tiempo Estimado | Prioridad |
|------|---------------|-----------------|-----------|
| 1 | Sistema de Notificaciones | 9-14 horas | Alta |
| 2 | Historial y Estadísticas | 10-15 horas | Alta |
| 3 | Gestión de Dispositivos | 9-14 horas | Media |
| 4 | Dashboard Mejorado | 11-16 horas | Media |
| 5 | Detección de Vulnerabilidades | 12-18 horas | Alta |

**Tiempo total estimado**: 51-77 horas (aproximadamente 1.5-2 meses a tiempo parcial)

---

## 🔧 Dependencias Técnicas

### Go Modules
```bash
go get github.com/gorilla/websocket
go get github.com/jung-kurt/gofpdf  # Para reportes PDF (opcional)
```

### Frontend
- Chart.js (gráficas)
- WebSocket API (nativo del navegador)
- Bootstrap 5 (ya incluido)

---

## ✅ Criterios de Éxito

### Fase 1 - Notificaciones
- [x] Se detectan y notifican nuevos dispositivos
- [x] Se detectan dispositivos desconectados
- [x] Se detectan cambios en puertos
- [x] Al menos 2 canales de notificación funcionando

### Fase 2 - Historial
- [x] Se registra el histórico de todos los cambios
- [x] Se pueden consultar estadísticas de los últimos 90 días
- [x] Dashboard muestra gráficas temporales
- [x] Se calcula uptime correctamente

### Fase 3 - Gestión
- [x] Se pueden añadir nombres personalizados
- [x] Sistema de tags funcional
- [x] Grupos con asignación automática
- [x] Exportación/importación funcional

### Fase 4 - Dashboard
- [x] Actualización en tiempo real vía WebSocket
- [x] Búsqueda avanzada funcional
- [x] Tema oscuro implementado
- [x] Responsive en móviles

### Fase 5 - Seguridad
- [x] Se detectan al menos 10 tipos de vulnerabilidades
- [x] Se generan recomendaciones útiles
- [x] Dashboard de seguridad funcional
- [x] Reportes exportables

---

## 🚨 Riesgos y Mitigaciones

| Riesgo | Probabilidad | Impacto | Mitigación |
|--------|--------------|---------|------------|
| Rendimiento con muchos dispositivos | Media | Alto | Implementar paginación y lazy loading |
| WebSockets consumen recursos | Media | Medio | Implementar heartbeat y límite de conexiones |
| Base de datos crece mucho | Alta | Medio | Implementar limpieza automática de datos antiguos |
| Falsos positivos en seguridad | Media | Medio | Permitir marcar como falso positivo |
| Complejidad del código aumenta | Alta | Bajo | Mantener buena documentación y tests |

---

## 📝 Notas de Implementación

### Orden Recomendado
1. **Fase 1** (Notificaciones) - Base para todo lo demás
2. **Fase 2** (Historial) - Necesario para estadísticas
3. **Fase 4** (Dashboard) - Mejora la experiencia de usuario
4. **Fase 3** (Gestión) - Añade personalización
5. **Fase 5** (Seguridad) - Añade valor diferencial

### Testing
- Cada fase debe incluir tests unitarios básicos
- Tests de integración para APIs
- Tests manuales del dashboard

### Documentación
- Actualizar README.md después de cada fase
- Documentar nuevos endpoints en API
- Añadir comentarios en código complejo

---

## 🎯 Siguiente Nivel (Post Top 5)

Después de completar estas 5 fases, consultar el archivo `NEXT_STEPS.md` para las siguientes funcionalidades a implementar.
