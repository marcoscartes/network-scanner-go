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
- [ ] Crear modelo `Notification` en `internal/database/models.go`
  - Campos: ID, Type, DeviceIP, Message, Timestamp, Read, Severity
- [ ] Crear modelo `NotificationConfig` para configuración
  - Campos: EnabledChannels, EmailConfig, TelegramConfig, WebhookURL
- [ ] Añadir tabla `notifications` a la base de datos
- [ ] Añadir tabla `notification_config` a la base de datos

#### 1.2 Detector de Cambios (2-3 horas)
- [ ] Crear `internal/notifications/detector.go`
  - Función `DetectNewDevice(device *Device) bool`
  - Función `DetectDisconnectedDevice(device *Device) bool`
  - Función `DetectPortChanges(old, new *Device) []int`
  - Función `CompareDeviceStates(old, new []*Device) []Change`

#### 1.3 Sistema de Notificaciones (3-4 horas)
- [ ] Crear `internal/notifications/notifier.go`
  - Interface `Notifier` con método `Send(notification *Notification) error`
  - Implementar `ConsoleNotifier` (logs)
  - Implementar `SystemNotifier` (notificaciones del SO - Windows toast)
  - Implementar `WebhookNotifier` (POST HTTP)
- [ ] Crear `internal/notifications/manager.go`
  - Gestionar múltiples notificadores
  - Cola de notificaciones
  - Rate limiting para evitar spam

#### 1.4 Integración con Scanner (1-2 horas)
- [ ] Modificar `cmd/scanner/main.go`
  - Inicializar NotificationManager
  - Comparar estado anterior vs nuevo después de cada scan
  - Enviar notificaciones de cambios detectados
- [ ] Añadir flags de configuración
  - `-notify-new-devices` (bool)
  - `-notify-disconnected` (bool)
  - `-notify-port-changes` (bool)
  - `-webhook-url` (string)

#### 1.5 API y Dashboard (2-3 horas)
- [ ] Añadir endpoints en `internal/web/server.go`
  - `GET /api/notifications` - Listar notificaciones
  - `POST /api/notifications/{id}/read` - Marcar como leída
  - `DELETE /api/notifications/{id}` - Eliminar
  - `GET /api/notifications/config` - Obtener configuración
  - `PUT /api/notifications/config` - Actualizar configuración
- [ ] Actualizar dashboard
  - Badge con contador de notificaciones no leídas
  - Panel lateral con lista de notificaciones
  - Configuración de notificaciones

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
- [ ] Crear modelo `DeviceHistory` en `internal/database/models.go`
  - Campos: ID, DeviceMAC, IP, Hostname, OpenPorts, Timestamp, ChangeType
- [ ] Crear modelo `NetworkStats` para estadísticas agregadas
  - Campos: Date, TotalDevices, NewDevices, DisconnectedDevices, TotalPorts
- [ ] Añadir tabla `device_history` a la base de datos
- [ ] Añadir tabla `network_stats` a la base de datos
- [ ] Crear índices para consultas rápidas por fecha y MAC

#### 2.2 Sistema de Registro Histórico (2-3 horas)
- [ ] Crear `internal/history/recorder.go`
  - Función `RecordDeviceState(device *Device, changeType string)`
  - Función `RecordNetworkSnapshot(devices []*Device)`
  - Función `CalculateDailyStats(date time.Time) *NetworkStats`
- [ ] Crear `internal/history/analyzer.go`
  - Función `GetDeviceUptime(mac string, period time.Duration) float64`
  - Función `GetDeviceHistory(mac string, from, to time.Time) []*DeviceHistory`
  - Función `GetNetworkTrends(days int) []*NetworkStats`
  - Función `GetMostActiveDevices(limit int) []*Device`

#### 2.3 Integración con Scanner (1-2 horas)
- [ ] Modificar `cmd/scanner/main.go`
  - Registrar snapshot después de cada scan
  - Calcular y guardar estadísticas diarias
- [ ] Añadir tarea programada para limpieza de datos antiguos
  - Flag `-history-retention-days` (default: 90)

#### 2.4 API de Estadísticas (2-3 horas)
- [ ] Añadir endpoints en `internal/web/server.go`
  - `GET /api/history/device/{mac}` - Historial de un dispositivo
  - `GET /api/history/network` - Historial de la red
  - `GET /api/stats/overview` - Estadísticas generales
  - `GET /api/stats/trends` - Tendencias temporales
  - `GET /api/stats/uptime/{mac}` - Uptime de dispositivo

#### 2.5 Visualización en Dashboard (3-4 horas)
- [ ] Crear sección de estadísticas en el dashboard
  - Gráfica de dispositivos activos en el tiempo (Chart.js)
  - Top 10 dispositivos más activos
  - Estadísticas de uptime
  - Timeline de eventos importantes
- [ ] Añadir vista detallada por dispositivo
  - Historial completo de cambios
  - Gráfica de disponibilidad
  - Cambios de puertos en el tiempo

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
- [ ] Extender modelo `Device` en `internal/database/models.go`
  - Añadir campo `CustomName` (string)
  - Añadir campo `Notes` (string)
  - Añadir campo `Tags` ([]string - JSON)
  - Añadir campo `Group` (string)
  - Añadir campo `IsFavorite` (bool)
  - Añadir campo `CustomIcon` (string)
- [ ] Migrar base de datos para añadir nuevas columnas
- [ ] Crear modelo `DeviceGroup` para grupos predefinidos

#### 3.2 Lógica de Gestión (2-3 horas)
- [ ] Crear `internal/management/device_manager.go`
  - Función `UpdateDeviceMetadata(mac string, metadata *DeviceMetadata)`
  - Función `AddTag(mac string, tag string)`
  - Función `RemoveTag(mac string, tag string)`
  - Función `SetGroup(mac string, group string)`
  - Función `ToggleFavorite(mac string)`
- [ ] Crear `internal/management/groups.go`
  - Función `CreateGroup(name, description string)`
  - Función `GetAllGroups() []*DeviceGroup`
  - Función `GetDevicesByGroup(group string) []*Device`
  - Función `AutoAssignGroups(devices []*Device)` - Asignación automática

#### 3.3 API de Gestión (2-3 horas)
- [ ] Añadir endpoints en `internal/web/server.go`
  - `PUT /api/devices/{mac}/name` - Cambiar nombre
  - `PUT /api/devices/{mac}/notes` - Actualizar notas
  - `POST /api/devices/{mac}/tags` - Añadir tag
  - `DELETE /api/devices/{mac}/tags/{tag}` - Eliminar tag
  - `PUT /api/devices/{mac}/group` - Cambiar grupo
  - `POST /api/devices/{mac}/favorite` - Toggle favorito
  - `GET /api/groups` - Listar grupos
  - `POST /api/groups` - Crear grupo

#### 3.4 UI de Gestión (3-4 horas)
- [ ] Actualizar dashboard con funcionalidades de gestión
  - Modal de edición de dispositivo (nombre, notas, tags, grupo)
  - Selector de iconos personalizados
  - Filtros por tags, grupos, favoritos
  - Vista de grupos con drag & drop
  - Indicador visual de favoritos (estrella)
- [ ] Añadir búsqueda avanzada
  - Por nombre, IP, MAC, tags, grupo
  - Autocompletado de tags

#### 3.5 Importación/Exportación (1-2 horas)
- [ ] Crear `internal/management/import_export.go`
  - Función `ExportDevices(format string) ([]byte, error)` - JSON/CSV
  - Función `ImportDevices(data []byte, format string) error`
- [ ] Añadir endpoints
  - `GET /api/export?format=json|csv`
  - `POST /api/import`

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
- [ ] Añadir dependencia `github.com/gorilla/websocket`
- [ ] Crear `internal/web/websocket.go`
  - Gestionar conexiones WebSocket
  - Broadcast de actualizaciones a clientes conectados
  - Heartbeat para mantener conexiones vivas
- [ ] Modificar `internal/web/server.go`
  - Endpoint `GET /ws` para conexiones WebSocket
  - Enviar actualizaciones cuando hay cambios en dispositivos
- [ ] Actualizar frontend
  - Conectar a WebSocket
  - Actualizar tabla automáticamente sin recargar

#### 4.2 Búsqueda y Filtrado Avanzado (2-3 horas)
- [ ] Crear `internal/search/query_parser.go`
  - Parser de queries tipo "ip:192.168.1.* AND ports:80,443"
  - Soporte para operadores: AND, OR, NOT
  - Campos: ip, mac, hostname, type, ports, tags, group
- [ ] Añadir endpoint `GET /api/search?q=query`
- [ ] Actualizar UI
  - Barra de búsqueda con sintaxis avanzada
  - Sugerencias de autocompletado
  - Filtros rápidos (botones para favoritos, tipos comunes, etc.)

#### 4.3 Mejoras Visuales (3-4 horas)
- [ ] Implementar tema oscuro
  - CSS variables para colores
  - Toggle en el dashboard
  - Guardar preferencia en localStorage
- [ ] Mejorar diseño de la tabla
  - Paginación
  - Ordenamiento por múltiples columnas
  - Columnas redimensionables
  - Exportar vista actual (CSV/JSON)
- [ ] Añadir indicadores visuales
  - Estado online/offline con colores
  - Badges para tags
  - Iconos para tipos de dispositivos
  - Animaciones suaves

#### 4.4 Dashboard Widgets (2-3 horas)
- [ ] Crear sección de widgets en la parte superior
  - Total de dispositivos activos
  - Nuevos dispositivos (últimas 24h)
  - Dispositivos offline
  - Puertos abiertos totales
  - Última actualización
- [ ] Añadir mini-gráficas
  - Sparklines de actividad
  - Distribución de tipos de dispositivos (pie chart)

#### 4.5 Responsive Design (1-2 horas)
- [ ] Optimizar para móviles
  - Tabla responsive con scroll horizontal
  - Menú hamburguesa
  - Touch-friendly buttons
- [ ] Optimizar para tablets
  - Layout adaptativo

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
- [ ] Crear `internal/security/vulnerability_db.go`
  - Mapa de puertos peligrosos con descripciones
  - Lista de servicios comunes con versiones vulnerables
  - Recomendaciones de seguridad por tipo de hallazgo
- [ ] Crear archivo de configuración `security_rules.json`
  - Reglas personalizables
  - Niveles de severidad (Critical, High, Medium, Low, Info)

#### 5.2 Motor de Análisis (3-4 horas)
- [ ] Crear `internal/security/scanner.go`
  - Función `AnalyzeDevice(device *Device) []*SecurityFinding`
  - Función `CheckDangerousPorts(ports []int) []*SecurityFinding`
  - Función `CheckServiceVersions(services map[int]string) []*SecurityFinding`
  - Función `CheckCommonVulnerabilities(device *Device) []*SecurityFinding`
- [ ] Crear modelo `SecurityFinding`
  - Campos: DeviceMAC, Type, Severity, Port, Service, Description, Recommendation, DetectedAt

#### 5.3 Integración con Scanner (1-2 horas)
- [ ] Modificar `cmd/scanner/main.go`
  - Ejecutar análisis de seguridad después de identificar dispositivo
  - Guardar findings en base de datos
  - Generar notificaciones para hallazgos críticos
- [ ] Añadir flag `-security-scan` (bool, default: true)

#### 5.4 Base de Datos de Findings (1-2 horas)
- [ ] Crear tabla `security_findings` en base de datos
- [ ] Añadir funciones CRUD en `internal/database/db.go`
  - `SaveSecurityFinding(finding *SecurityFinding)`
  - `GetFindingsByDevice(mac string) []*SecurityFinding`
  - `GetAllFindings() []*SecurityFinding`
  - `GetFindingsBySeverity(severity string) []*SecurityFinding`
  - `MarkFindingAsResolved(id int)`

#### 5.5 API y Dashboard de Seguridad (3-4 horas)
- [ ] Añadir endpoints en `internal/web/server.go`
  - `GET /api/security/findings` - Todos los hallazgos
  - `GET /api/security/findings/{mac}` - Por dispositivo
  - `GET /api/security/summary` - Resumen de seguridad
  - `POST /api/security/findings/{id}/resolve` - Marcar como resuelto
  - `POST /api/security/scan/{mac}` - Forzar scan de seguridad
- [ ] Crear sección de seguridad en dashboard
  - Panel de resumen (críticos, altos, medios, bajos)
  - Lista de hallazgos con filtros por severidad
  - Vista detallada por dispositivo
  - Recomendaciones de remediación
  - Indicador de "Security Score" por dispositivo

#### 5.6 Reportes de Seguridad (2-3 horas)
- [ ] Crear `internal/security/reporter.go`
  - Función `GenerateSecurityReport(format string) ([]byte, error)`
  - Soporte para PDF y HTML
  - Incluir gráficas y estadísticas
- [ ] Añadir endpoint `GET /api/security/report?format=pdf|html`
- [ ] Añadir botón de descarga en dashboard

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
- [ ] Se detectan y notifican nuevos dispositivos
- [ ] Se detectan dispositivos desconectados
- [ ] Se detectan cambios en puertos
- [ ] Al menos 2 canales de notificación funcionando

### Fase 2 - Historial
- [ ] Se registra el histórico de todos los cambios
- [ ] Se pueden consultar estadísticas de los últimos 90 días
- [ ] Dashboard muestra gráficas temporales
- [ ] Se calcula uptime correctamente

### Fase 3 - Gestión
- [ ] Se pueden añadir nombres personalizados
- [ ] Sistema de tags funcional
- [ ] Grupos con asignación automática
- [ ] Exportación/importación funcional

### Fase 4 - Dashboard
- [ ] Actualización en tiempo real vía WebSocket
- [ ] Búsqueda avanzada funcional
- [ ] Tema oscuro implementado
- [ ] Responsive en móviles

### Fase 5 - Seguridad
- [ ] Se detectan al menos 10 tipos de vulnerabilidades
- [ ] Se generan recomendaciones útiles
- [ ] Dashboard de seguridad funcional
- [ ] Reportes exportables

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
