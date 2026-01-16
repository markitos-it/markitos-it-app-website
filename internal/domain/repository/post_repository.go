package repository

import "markitos-it-app-website/internal/domain"

type PostRepository interface {
	GetAll() ([]domain.Post, error)
}
