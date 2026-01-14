package main

import (
	"log"
	"net/http"

	"markitos-it-app-website/internal/assets"
	"markitos-it-app-website/internal/infrastructure/http/handlers"
)

func main() {
	mux := http.NewServeMux()

	assetsFS := assets.GetAssetsFS()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFS))))

	mux.HandleFunc("/", handlers.IndexHandler)

	addr := "0.0.0.0:8080"
	log.Printf("Server starting on http://%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
