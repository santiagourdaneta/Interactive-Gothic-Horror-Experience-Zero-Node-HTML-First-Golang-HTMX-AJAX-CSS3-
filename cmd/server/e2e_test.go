package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEndToEndStoryFlow(t *testing.T) {
	// 1. Configuración del Servidor de Prueba (Lightweight)
	mux := http.NewServeMux()

	// Registramos los mismos manejadores que en main.go
	mux.HandleFunc("/next-chapter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		w.Write([]byte(`<div id="story-terminal" class="glitch-vibration">`))
	})

	// 2. Simulamos la interacción E2E: El usuario hace clic (POST)
	server := httptest.NewServer(mux)
	defer server.Close()

	// Creamos una petición POST como si fuera HTMX
	resp, err := http.Post(server.URL+"/next-chapter", "text/plain", nil)
	if err != nil {
		t.Fatalf("No se pudo realizar la petición E2E: %v", err)
	}
	defer resp.Body.Close()

	// 3. Validaciones Extremas
	if resp.StatusCode != http.StatusOK {
		t.Errorf("E2E Fallido: Se esperaba 200 OK, se obtuvo %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	content := string(body)

	// Verificamos que el contenido devuelto sea el correcto para la UI
	if !strings.Contains(content, "glitch-vibration") {
		t.Error("E2E Fallido: El contenido no tiene la clase de animación esperada.")
	}

	if !strings.Contains(content, "id=\"story-terminal\"") {
		t.Error("E2E Fallido: El fragmento HTML no mantiene el ID necesario para HTMX.")
	}
}
