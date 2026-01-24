package main

import (
	"log"
	"net/http"

	"markitos-it-app-website/internal/infrastructure/http/handlers"
)

func main() {
	homeHandler, err := handlers.NewHomeHandler()
	if err != nil {
		log.Fatalf("Failed to create home handler: %v", err)
	}

	docsHandler, err := handlers.NewDocsHandler()
	if err != nil {
		log.Fatalf("Failed to create docs handler: %v", err)
	}

	mux := http.NewServeMux()

	// Rutas
	mux.HandleFunc("/", homeHandler.Index)
	mux.HandleFunc("/docs", docsHandler.Index)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := "0.0.0.0:8080"
	log.Printf("🚀 Server starting on http://%s", addr)
	log.Printf("   Home: http://localhost:8080/")
	log.Printf("   Docs: http://localhost:8080/docs")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
