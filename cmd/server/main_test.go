package main

import (
	"net/http"
	"testing"
	"time"
)

// TestHealthCheck verifica que el servidor responda 200 OK
func TestHealthCheck(t *testing.T) {
	// Le damos un momento al servidor para estar listo si es necesario
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	// Probamos la ruta principal
	resp, err := client.Get("http://localhost:9000/")
	if err != nil {
		t.Fatalf("El servidor no responde en el puerto 9000. ¿Está encendido? Error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba código 200, pero se obtuvo: %d", resp.StatusCode)
	}
}
