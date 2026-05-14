package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Servir archivos estáticos (CSS/JS)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Ruta principal: El Relato
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/static/css/retro.css">
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <title>Gothic 80s Interface</title>
</head>
<body>
    <div class="container">
        <h1 class="glitch">SISTEMA OPERATIVO Z</h1>
        <div id="story-terminal">
            <p>La laptop de 2014 zumba con esfuerzo. En la pantalla, una <b>mujer de facciones pálidas y curvas imponentes</b> te sonríe, revelando colmillos afilados. Su gato zombie mastica un cable de red.</p>
            <button hx-post="/next-chapter" 
                    hx-target="#story-terminal" 
                    hx-swap="outerHTML"
                    class="btn-retro">
                ESTABLECER VÍNCULO
            </button>
        </div>
    </div>
</body>
</html>
`))
	})

	// Ruta para el siguiente capítulo (HTMX)
	mux.HandleFunc("/next-chapter", func(w http.ResponseWriter, r *http.Request) {
		// Validación extrema: Solo aceptamos peticiones POST que vengan de HTMX
		if r.Method != http.MethodPost {
			http.Error(w, "Acceso Denegado: Protocolo Humano Detectado", http.StatusForbidden)
			return
		}

		// Simulamos un pequeño retraso para que se sienta "procesando"
		time.Sleep(500 * time.Millisecond)

		w.Write([]byte(`
	        <div id="story-terminal" class="glitch-vibration">
	            <p style="color: #ff00ff;">[NIVEL DE ACCESO: VAMPIRO]</p>
	            <p>La mujer gorda sonríe y el gato zombie salta sobre tu teclado virtual. Has entrado al sistema central de 1986.</p>
	            <button class="btn-retro" onclick="location.reload()">DESCONECTARSE</button>
	        </div>
	    `))
	})

	server := &http.Server{
		Addr:         ":9000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Servidor iniciado en http://localhost:9000")
	log.Fatal(server.ListenAndServe())
}
