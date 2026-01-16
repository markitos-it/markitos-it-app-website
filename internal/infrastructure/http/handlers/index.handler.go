package handlers

import (
	"html/template"
	"log"
	"net/http"

	"markitos-it-app-website/internal/domain/repository"
	"markitos-it-app-website/internal/templates"
)

type IndexModel struct {
	Title string
	Posts []PostView
}

type PostView struct {
	Title   string
	Content string
	Date    string
}

type IndexHandler struct {
	postRepo repository.PostRepository
}

func NewIndexHandler(postRepo repository.PostRepository) *IndexHandler {
	return &IndexHandler{postRepo: postRepo}
}

func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postRepo.GetAll()
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	postViews := make([]PostView, len(posts))
	for i, post := range posts {
		postViews[i] = PostView{
			Title:   post.Title,
			Content: post.Content,
			Date:    post.CreatedAt.Format("2006-01-02"),
		}
	}

	model := IndexModel{
		Title: "Dashboard de Contenido",
		Posts: postViews,
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
