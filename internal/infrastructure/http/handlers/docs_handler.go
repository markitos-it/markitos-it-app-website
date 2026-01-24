package handlers

import (
	"html/template"
	"io"
	"markitos-it-app-website/internal/templates"
	"net/http"
)

type DocsHandler struct {
	tmpl *template.Template
}

func NewDocsHandler() (*DocsHandler, error) {
	tmpl, err := template.New("base.html").ParseFS(
		templates.FS(),
		"shared/base.html",
		"shared/head.html",
		"shared/navbar.html",
		"shared/sidebar.html",
		"shared/scripts.html",
		"shared/styles.css",
		"shared/common.js",
		"docs/index/content.html",
		"docs/index/styles.css",
		"docs/index/script.js",
	)
	if err != nil {
		return nil, err
	}
	return &DocsHandler{tmpl: tmpl}, nil
}

func (h *DocsHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Leer CSS y JS compartidos
	sharedCSS, _ := templates.FS().Open("shared/styles.css")
	sharedCSSBytes, _ := io.ReadAll(sharedCSS)
	sharedCSS.Close()

	sharedJS, _ := templates.FS().Open("shared/common.js")
	sharedJSBytes, _ := io.ReadAll(sharedJS)
	sharedJS.Close()

	// Leer CSS y JS específicos de la página
	pageCSS, _ := templates.FS().Open("docs/index/styles.css")
	pageCSSBytes, _ := io.ReadAll(pageCSS)
	pageCSS.Close()

	pageJS, _ := templates.FS().Open("docs/index/script.js")
	pageJSBytes, _ := io.ReadAll(pageJS)
	pageJS.Close()

	data := map[string]interface{}{
		"PageClass":     "docs-page",
		"Title":         "Keptn Docs",
		"ActiveSection": "docs",
		"Sections": []map[string]interface{}{
			{"Title": "Keptn Integrations", "Slug": "keptn-integrations", "Active": true},
			{"Title": "Helm Charts", "Slug": "helm-charts", "Active": false},
			{"Title": "Container Images", "Slug": "container-images", "Active": false},
		},
		"Content": template.HTML(`
			<h1>Keptn Integrations</h1>
			<p>This section explains how to set up the Keptn control plane.</p>
			<div class="code-box">$ artifacthubctl install keptn</div>
			<p>Once installed, you can manage your sequences directly from the dashboard.</p>
		`),
		"TableOfContents": []map[string]string{
			{"Anchor": "installation", "Text": "Installation"},
			{"Anchor": "configuration", "Text": "Configuration"},
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

func (h *DocsHandler) View(w http.ResponseWriter, r *http.Request) {
	h.Index(w, r)
}
