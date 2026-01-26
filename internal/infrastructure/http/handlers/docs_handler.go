package handlers

import (
	"encoding/base64"
	"html/template"
	"io"
	"markitos-it-app-website/internal/domain/documents"
	"markitos-it-app-website/internal/templates"
	"net/http"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type DocsHandler struct {
	indexTmpl *template.Template
	viewTmpl  *template.Template
	markdown  goldmark.Markdown
}

func NewDocsHandler() (*DocsHandler, error) {
	indexTmpl, err := template.New("base.html").ParseFS(
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

	viewTmpl, err := template.New("base.html").ParseFS(
		templates.FS(),
		"shared/base.html",
		"shared/head.html",
		"shared/navbar.html",
		"shared/sidebar.html",
		"shared/scripts.html",
		"shared/styles.css",
		"shared/common.js",
		"docs/view/content.html",
		"docs/view/styles.css",
		"docs/view/script.js",
	)
	if err != nil {
		return nil, err
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.TaskList,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(), // allow raw HTML (iframe) so the YouTube player renders
		),
	)

	return &DocsHandler{
		indexTmpl: indexTmpl,
		viewTmpl:  viewTmpl,
		markdown:  md,
	}, nil
}

func (h *DocsHandler) Index(w http.ResponseWriter, r *http.Request) {
	sharedCSS, _ := templates.FS().Open("shared/styles.css")
	sharedCSSBytes, _ := io.ReadAll(sharedCSS)
	sharedCSS.Close()

	sharedJS, _ := templates.FS().Open("shared/common.js")
	sharedJSBytes, _ := io.ReadAll(sharedJS)
	sharedJS.Close()

	pageCSS, _ := templates.FS().Open("docs/index/styles.css")
	pageCSSBytes, _ := io.ReadAll(pageCSS)
	pageCSS.Close()

	pageJS, _ := templates.FS().Open("docs/index/script.js")
	pageJSBytes, _ := io.ReadAll(pageJS)
	pageJS.Close()

	docs, err := documents.GetAllDocuments()
	if err != nil {
		http.Error(w, "Error loading documents", http.StatusInternalServerError)
		return
	}

	categories := []string{"All"}
	categoryMap := make(map[string]bool)
	for _, doc := range docs {
		if !categoryMap[doc.Category] {
			categories = append(categories, doc.Category)
			categoryMap[doc.Category] = true
		}
	}

	docsInterface := make([]map[string]interface{}, len(docs))
	for i, doc := range docs {
		docsInterface[i] = map[string]interface{}{
			"ID":          doc.ID,
			"Title":       doc.Title,
			"Description": doc.Description,
			"Category":    doc.Category,
			"Tags":        doc.Tags,
			"UpdatedAt":   doc.UpdatedAt,
			"ContentB64":  doc.ContentB64,
			"CoverImage":  doc.CoverImage,
		}
	}

	data := map[string]interface{}{
		"PageClass":     "docs-page",
		"Title":         "Documentation Dashboard",
		"ActiveSection": "docs",
		"Documents":     docsInterface,
		"Categories":    categories,
		"SharedStyles":  template.CSS(string(sharedCSSBytes)),
		"SharedScript":  template.JS(string(sharedJSBytes)),
		"PageStyles":    template.CSS(string(pageCSSBytes)),
		"PageScript":    template.JS(string(pageJSBytes)),
	}

	if err := h.indexTmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *DocsHandler) View(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	docID := strings.TrimPrefix(path, "/docs/")

	doc, err := documents.GetDocumentById(docID)
	if err != nil {
		http.Error(w, "Error loading document", http.StatusInternalServerError)
		return
	}

	if doc == nil {
		http.NotFound(w, r)
		return
	}

	contentMarkdown, err := base64.StdEncoding.DecodeString(doc.ContentB64)
	if err != nil {
		http.Error(w, "Error decoding document", http.StatusInternalServerError)
		return
	}

	var htmlContent strings.Builder
	if err := h.markdown.Convert(contentMarkdown, &htmlContent); err != nil {
		http.Error(w, "Error converting markdown", http.StatusInternalServerError)
		return
	}

	sharedCSS, _ := templates.FS().Open("shared/styles.css")
	sharedCSSBytes, _ := io.ReadAll(sharedCSS)
	sharedCSS.Close()

	sharedJS, _ := templates.FS().Open("shared/common.js")
	sharedJSBytes, _ := io.ReadAll(sharedJS)
	sharedJS.Close()

	pageCSS, _ := templates.FS().Open("docs/view/styles.css")
	pageCSSBytes, _ := io.ReadAll(pageCSS)
	pageCSS.Close()

	pageJS, _ := templates.FS().Open("docs/view/script.js")
	pageJSBytes, _ := io.ReadAll(pageJS)
	pageJS.Close()

	data := map[string]interface{}{
		"PageClass":     "docs-view-page",
		"Title":         doc.Title,
		"ActiveSection": "docs",
		"ID":            doc.ID,
		"Category":      doc.Category,
		"Description":   doc.Description,
		"Tags":          doc.Tags,
		"UpdatedAt":     doc.UpdatedAt,
		"CoverImage":    doc.CoverImage,
		"Content":       template.HTML(htmlContent.String()),
		"SharedStyles":  template.CSS(string(sharedCSSBytes)),
		"SharedScript":  template.JS(string(sharedJSBytes)),
		"PageStyles":    template.CSS(string(pageCSSBytes)),
		"PageScript":    template.JS(string(pageJSBytes)),
	}

	if err := h.viewTmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
