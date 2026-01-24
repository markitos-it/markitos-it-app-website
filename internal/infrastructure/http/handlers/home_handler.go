package handlers

import (
	"html/template"
	"io"
	"markitos-it-app-website/internal/templates"
	"net/http"
)

type HomeHandler struct {
	tmpl *template.Template
}

func NewHomeHandler() (*HomeHandler, error) {
	tmpl, err := template.New("base.html").ParseFS(
		templates.FS(),
		"shared/base.html",
		"shared/head.html",
		"shared/navbar.html",
		"shared/sidebar.html",
		"shared/scripts.html",
		"shared/styles.css",
		"shared/common.js",
		"home/index/content.html",
		"home/index/styles.css",
		"home/index/script.js",
	)
	if err != nil {
		return nil, err
	}
	return &HomeHandler{tmpl: tmpl}, nil
}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Leer CSS y JS compartidos
	sharedCSS, _ := templates.FS().Open("shared/styles.css")
	sharedCSSBytes, _ := io.ReadAll(sharedCSS)
	sharedCSS.Close()

	sharedJS, _ := templates.FS().Open("shared/common.js")
	sharedJSBytes, _ := io.ReadAll(sharedJS)
	sharedJS.Close()

	// Leer CSS y JS específicos de la página
	pageCSS, _ := templates.FS().Open("home/index/styles.css")
	pageCSSBytes, _ := io.ReadAll(pageCSS)
	pageCSS.Close()

	pageJS, _ := templates.FS().Open("home/index/script.js")
	pageJSBytes, _ := io.ReadAll(pageJS)
	pageJS.Close()

	data := map[string]interface{}{
		"PageClass":     "home-page",
		"Title":         "Home",
		"ActiveSection": "home",
		"ResultsCount":  "2,450",
		"Packages": []map[string]string{
			{
				"ID":          "prometheus",
				"Name":        "Prometheus",
				"Version":     "v2.45.0",
				"Description": "The official Prometheus monitoring system for Kubernetes clusters.",
				"Icon":        "H",
				"IconClass":   "",
				"Badge":       "Official",
				"Stars":       "4.5k",
				"UpdatedAgo":  "2d ago",
			},
			{
				"ID":          "gatekeeper",
				"Name":        "Gatekeeper",
				"Version":     "v3.13.0",
				"Description": "Policy Controller for Kubernetes using Open Policy Agent.",
				"Icon":        "O",
				"IconClass":   "purple",
				"Badge":       "Verified",
				"Stars":       "1.2k",
				"UpdatedAgo":  "5h ago",
			},
		},
		"SharedStyles": template.CSS(string(sharedCSSBytes)),
		"SharedScript": template.JS(string(sharedJSBytes)),
		"PageStyles":   template.CSS(string(pageCSSBytes)),
		"PageScript":   template.JS(string(pageJSBytes)),
	}
	if err := h.tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
