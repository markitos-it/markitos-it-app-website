package handlers

import (
	"html/template"
	"net/http"

	"markitos-it-app-website/internal/templates"
)

type IndexModel struct {
	Title string
	Posts []Post
}

type Post struct {
	Title       string
	Description string
	Tag         string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	model := IndexModel{
		Title: "Dashboard de Contenido",
		Posts: []Post{
			{
				Title:       "Introducción a JS Moderno",
				Description: "Aprende las bases de ES6+ sin morir en el intento.",
				Tag:         "Tutorial",
			},
			{
				Title:       "CSS Grid vs Flexbox",
				Description: "¿Cuándo usar cada uno? Guía definitiva para 2026.",
				Tag:         "Diseño",
			},
			{
				Title:       "Optimización Web",
				Description: "Cómo lograr un 100 en Lighthouse fácilmente.",
				Tag:         "Performance",
			},
		},
	}

	templatesFS := templates.GetTemplatesFS()

	tmpl, err := template.ParseFS(templatesFS, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
