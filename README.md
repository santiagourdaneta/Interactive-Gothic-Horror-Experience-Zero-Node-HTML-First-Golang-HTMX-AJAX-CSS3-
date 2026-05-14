# 🧛‍♂️ PROTOCOLO ZOMBIE 1986
### *Interactive Gothic Horror Experience | High-Performance Engine*

Una plataforma de narrativa envolvente e interactiva con estética **Retro-80s CRT**, construida bajo la filosofía **Zero-Node** y **HTML-First**. Diseñada para volar en hardware limitado (probada en laptops de 2014) manteniendo una calidad técnica extrema y seguridad absoluta.

---

## 🛠️ Stack Tecnológico (The "Pure Metal" Way)

Este proyecto rechaza la sobrecarga de los frameworks modernos para abrazar el rendimiento puro:

*   **Lenguaje:** [Go (Golang)](https://go.dev/) - El motor de alto rendimiento y tipado fuerte.
*   **Interactividad:** [HTMX](https://htmx.org/) - AJAX moderno sin escribir una línea de JavaScript pesado.
*   **Estética:** CSS3 puro con efectos de fósforo, scanlines y animaciones de glitch.
*   **Calidad:** `Staticcheck` (Linter), `go fmt` (Formatter) y `httptest` (E2E Testing).
*   **Infraestructura:** Despliegue mediante binarios estáticos (Docker multi-stage).

---

## 🚀 Rendimiento Probado (Benchmarks)

Ejecutado en hardware vintage (CPU 2014) con las siguientes métricas de estrés:

| Métrica | Resultado |
| :--- | :--- |
| **Requests por Segundo** | ~2,028 req/sec |
| **Tiempo de Respuesta (Avg)** | 48ms |
| **Total Peticiones (Stress Test)** | 10,000 en 4.9s |
| **Consumo de RAM** | < 20MB |

---

## 📂 Estructura del Proyecto

```text
.
├── cmd/server          # Punto de entrada (main.go) y Tests E2E
├── internal/           # Lógica de negocio y validaciones extremas
├── static/             # Assets (CSS retro, JS ligero)
│   ├── css/            # Estilos CRT y animaciones neón
│   └── js/             # HTMX para interactividad
├── views/              # Componentes HTML-First (Templates)
├── Makefile            # Comandos de orquestación (Build, Test, Lint)
└── .git/hooks          # Blindaje pre-commit (Fmt + Lint + Test)

🔧 Instalación y Desarrollo
Sigue estos pasos para levantar el sistema en tu entorno local:

1. Clonar el repositorio:

git clone [https://github.com/santiagourdaneta/Interactive-Gothic-Horror-Experience-Zero-Node-HTML-First-Golang-HTMX-AJAX-CSS3-](https://github.com/santiagourdaneta/Interactive-Gothic-Horror-Experience-Zero-Node-HTML-First-Golang-HTMX-AJAX-CSS3-)
cd Interactive-Gothic-Horror-Experience-Zero-Node-HTML-First-Golang-HTMX-AJAX-CSS3-

2.  **Validar calidad (Linter & Tests):**
    ```bash
    # Ejecuta el análisis estático
    staticcheck ./...
    # Corre las pruebas unitarias y de integración
    go test ./... -v
    ```

3.  **Encender el sistema:**
    ```bash
    go run cmd/server/main.go

🛡️ Protocolos de Seguridad (CI/CD Local)
Este repositorio incluye un Pre-commit Hook que garantiza que el código sea "Perfectamente Seguro" antes de cada commit:

Auto-Formatting: El código siempre sigue el estándar oficial de Go.

Static Analysis: El linter bloquea cualquier commit con posibles errores de memoria o lógica.

E2E Guard: Las pruebas de integración aseguran que el flujo del relato no se rompa (validación de rutas y respuestas HTMX).

🌑 El Relato Envolvente
"En un Surco sumido en la oscuridad, una mujer vampiro aguarda tras una terminal de 1986. No busca sangre, busca código limpio. Su gato zombie custodia el buffer de entrada. ¿Estás listo para establecer el vínculo?"

Desarrollado por Santiago Urdaneta Anton

Ingeniería Minimalista | Rendimiento Extremo | Estética Vintage    