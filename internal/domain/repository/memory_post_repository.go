package repository

import (
	"time"

	"markitos-it-app-website/internal/domain"
)

type MemoryPostRepository struct {
	posts []domain.Post
}

func NewMemoryPostRepository() *MemoryPostRepository {
	return &MemoryPostRepository{
		posts: []domain.Post{
			{
				ID:        1,
				Title:     "Bienvenido a Markitos IT",
				Content:   "Este es el primer post de prueba. Aquí compartiremos contenido sobre tecnología, programación y desarrollo de software.",
				CreatedAt: time.Date(2026, 1, 15, 10, 0, 0, 0, time.UTC),
			},
			{
				ID:        2,
				Title:     "Introducción a Go",
				Content:   "Go es un lenguaje de programación moderno, eficiente y fácil de aprender. Perfecto para construir aplicaciones web y microservicios.",
				CreatedAt: time.Date(2026, 1, 14, 15, 30, 0, 0, time.UTC),
			},
			{
				ID:        3,
				Title:     "Arquitectura Limpia",
				Content:   "La arquitectura limpia nos ayuda a mantener nuestro código organizado, testeable y mantenible a largo plazo.",
				CreatedAt: time.Date(2026, 1, 13, 9, 15, 0, 0, time.UTC),
			},
		},
	}
}

func (r *MemoryPostRepository) GetAll() ([]domain.Post, error) {
	return r.posts, nil
}
