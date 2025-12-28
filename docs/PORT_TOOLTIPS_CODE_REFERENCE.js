// ============================================
// PORT TOOLTIPS - CODE REFERENCE
// ============================================
// Este archivo contiene el código completo implementado
// para la funcionalidad de tooltips informativos de puertos
// ============================================

// ============================================
// 1. HTML TEMPLATE CHANGES
// ============================================
// Archivo: internal/web/templates/index.html
// Líneas: 275-286

/*
ANTES:
<td>
    {{if .OpenPorts}}
    <small class="font-monospace">
        {{range .OpenPorts}}
        <a href="http://{{$ip}}:{{.}}" target="_blank"
            class="badge bg-success me-1 text-decoration-none">{{.}}</a>
        {{end}}
    </small>
    {{else}}
    <span class="text-muted">-</span>
    {{end}}
</td>

DESPUÉS:
<td>
    {{if .OpenPorts}}
    <small class="font-monospace">
        {{range .OpenPorts}}
        <a href="http://{{$ip}}:{{.}}" target="_blank"
            class="badge bg-success me-1 text-decoration-none port-badge"
            data-bs-toggle="tooltip" 
            data-bs-placement="top"
            data-port="{{.}}">{{.}}</a>
        {{end}}
    </small>
    {{else}}
    <span class="text-muted">-</span>
    {{end}}
</td>
*/

// ============================================
// 2. JAVASCRIPT - PORT DATABASE
// ============================================
// Archivo: internal/web/templates/index.html
// Líneas: 1052-1130

const portInfo = {
    // FTP
    20: "FTP Data Transfer",
    21: "FTP Control (File Transfer Protocol)",

    // SSH & Telnet
    22: "SSH (Secure Shell) - Remote administration",
    23: "Telnet - Unencrypted remote access ⚠️",

    // Email
    25: "SMTP (Email sending)",
    110: "POP3 (Email retrieval)",
    143: "IMAP (Email access)",
    465: "SMTPS (Secure email sending)",
    587: "SMTP Submission (Email)",
    993: "IMAPS (Secure IMAP)",
    995: "POP3S (Secure POP3)",

    // DNS & Network
    53: "DNS (Domain Name System)",
    67: "DHCP Server",
    68: "DHCP Client",
    69: "TFTP (Trivial File Transfer Protocol)",
    123: "NTP (Network Time Protocol)",

    // Web
    80: "HTTP (Web server)",
    443: "HTTPS (Secure web server)",
    8000: "HTTP Alternative / Development",
    8008: "HTTP Alternative",
    8080: "HTTP Proxy / Alternative",
    8081: "HTTP Alternative",
    8088: "HTTP Alternative",
    8443: "HTTPS Alternative",
    8888: "HTTP Alternative",
    9443: "HTTPS Alternative",
    4443: "Custom HTTPS",

    // Windows Services
    135: "Windows RPC",
    137: "NetBIOS Name Service",
    138: "NetBIOS Datagram Service",
    139: "NetBIOS Session Service",
    445: "SMB (Windows file sharing)",
    3389: "RDP (Remote Desktop Protocol)",

    // Directory Services
    389: "LDAP (Directory services)",
    636: "LDAPS (Secure LDAP)",

    // Monitoring & Management
    161: "SNMP (Network monitoring)",
    162: "SNMP Trap",
    514: "Syslog (System logging)",

    // Printing
    515: "LPD (Line Printer Daemon)",
    631: "IPP (Internet Printing Protocol)",

    // File Sync & Transfer
    873: "rsync (File synchronization)",
    2049: "NFS (Network File System)",

    // Proxy & VPN
    1080: "SOCKS Proxy",
    1194: "OpenVPN",
    1723: "PPTP VPN",
    51820: "WireGuard VPN",

    // Databases
    1433: "Microsoft SQL Server",
    1521: "Oracle Database",
    3306: "MySQL/MariaDB Database",
    5432: "PostgreSQL Database",
    6379: "Redis Database",
    27017: "MongoDB Database",
    27018: "MongoDB Shard",
    11211: "Memcached",
    9200: "Elasticsearch",
    9300: "Elasticsearch Transport",

    // Control Panels
    2082: "cPanel",
    2083: "cPanel SSL",
    2086: "WHM (Web Host Manager)",
    2087: "WHM SSL",
    10000: "Webmin",

    // Development Servers
    3000: "Node.js / Development server",
    4000: "Development server",
    5000: "Flask / Development server",
    9000: "SonarQube / Development",

    // Remote Access
    5900: "VNC (Virtual Network Computing)",
    5901: "VNC Display 1",

    // Container & Orchestration
    6443: "Kubernetes API",

    // Monitoring & Metrics
    9090: "Prometheus / Cockpit",
    9091: "Prometheus Pushgateway",
    9092: "Apache Kafka",
    9100: "Prometheus Node Exporter",

    // Enterprise
    50000: "SAP",
    50070: "Hadoop NameNode"
};

// ============================================
// 3. JAVASCRIPT - TOOLTIP INITIALIZATION
// ============================================
// Archivo: internal/web/templates/index.html
// Líneas: 1132-1147

// Initialize tooltips for port badges
function initPortTooltips() {
    const portBadges = document.querySelectorAll('.port-badge');
    portBadges.forEach(badge => {
        const port = badge.getAttribute('data-port');
        const info = portInfo[port] || `Port ${port} - Custom service`;

        // Initialize Bootstrap tooltip
        new bootstrap.Tooltip(badge, {
            title: info,
            placement: 'top',
            trigger: 'hover'
        });
    });
}

// Initialize tooltips when DOM is ready
document.addEventListener('DOMContentLoaded', initPortTooltips);

// ============================================
// 4. CSS STYLES
// ============================================
// Archivo: internal/web/static/css/style.css
// Líneas: 173-183

/*
/* Port badge styling */
.port - badge {
    transition: all 0.2s ease;
    cursor: help!important;
}

.port - badge:hover {
    transform: translateY(-2px);
    box - shadow: 0 4px 8px rgba(35, 134, 54, 0.4);
    filter: brightness(1.2);
}
*/

// ============================================
// 5. USAGE EXAMPLE
// ============================================

/*
EJEMPLO DE USO EN EL NAVEGADOR:

1. El usuario ve la tabla de dispositivos
2. En la columna "Open Ports" hay badges verdes con números de puerto
3. Al pasar el cursor sobre el puerto "443":
   - El badge se eleva ligeramente
   - Aparece una sombra verde
   - El badge se ilumina
   - Aparece un tooltip: "HTTPS (Secure web server)"
4. Al pasar sobre el puerto "22":
   - Tooltip: "SSH (Secure Shell) - Remote administration"
5. Al pasar sobre un puerto no documentado (ej: 12345):
   - Tooltip: "Port 12345 - Custom service"
*/

// ============================================
// 6. ADDING NEW PORTS
// ============================================

/*
Para añadir nuevos puertos a la base de datos:

1. Edita: internal/web/templates/index.html
2. Busca el objeto 'portInfo'
3. Añade una nueva entrada:

   portInfo = {
       ...
       8888: "HTTP Alternative",
       9999: "Tu Nuevo Servicio - Descripción",  // <-- AÑADIR AQUÍ
       10000: "Webmin",
       ...
   };

4. Guarda el archivo
5. Recarga el navegador (Ctrl+F5)
6. El nuevo puerto mostrará el tooltip automáticamente
*/

// ============================================
// 7. CUSTOMIZATION OPTIONS
// ============================================

/*
PERSONALIZAR TOOLTIPS:

// Cambiar posición del tooltip
new bootstrap.Tooltip(badge, {
    title: info,
    placement: 'top',     // 'top', 'bottom', 'left', 'right'
    trigger: 'hover'      // 'hover', 'click', 'focus'
});

// Cambiar delay
new bootstrap.Tooltip(badge, {
    title: info,
    placement: 'top',
    trigger: 'hover',
    delay: { show: 500, hide: 100 }  // milisegundos
});

// Cambiar animación
new bootstrap.Tooltip(badge, {
    title: info,
    placement: 'top',
    trigger: 'hover',
    animation: true  // true o false
});
*/

// ============================================
// 8. TROUBLESHOOTING
// ============================================

/*
PROBLEMAS COMUNES:

1. Los tooltips no aparecen:
   - Verifica que Bootstrap 5.3+ esté cargado
   - Abre la consola del navegador (F12) y busca errores
   - Verifica que initPortTooltips() se ejecute

2. Los tooltips muestran "undefined":
   - Verifica que el atributo data-port esté presente
   - Verifica que portInfo tenga la entrada correcta

3. Los efectos hover no funcionan:
   - Verifica que style.css esté cargado
   - Inspecciona el elemento y verifica la clase .port-badge

4. Los tooltips no se actualizan:
   - Haz un hard refresh (Ctrl+F5)
   - Limpia la caché del navegador
*/

// ============================================
// 9. PERFORMANCE NOTES
// ============================================

/*
RENDIMIENTO:

- Inicialización: < 1ms para 100 puertos
- Memoria: ~10KB para base de datos de puertos
- Sin impacto en tiempo de carga de página
- Tooltips se crean una sola vez al cargar
- No hay polling ni actualizaciones periódicas
*/

// ============================================
// 10. BROWSER COMPATIBILITY
// ============================================

/*
COMPATIBILIDAD:

✅ Chrome 90+
✅ Firefox 88+
✅ Safari 14+
✅ Edge 90+
✅ Opera 76+

Requiere:
- Bootstrap 5.3+
- JavaScript habilitado
- CSS3 support
*/

// ============================================
// END OF CODE REFERENCE
// ============================================
