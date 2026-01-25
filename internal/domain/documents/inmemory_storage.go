package documents

import (
	"encoding/base64"
	"io"
	"markitos-it-app-website/internal/templates"
)

// GetAllDocuments retorna todos los documentos con su contenido en base64
// Simula cómo vendrían los datos desde una API
func GetAllDocuments() ([]Document, error) {
	docsMetadata := []struct {
		ID          string
		Title       string
		Description string
		Category    string
		Tags        []string
		UpdatedAt   string
		FilePath    string
		CoverImage  string
	}{
		{
			ID:          "getting-started-keptn",
			Title:       "Getting Started with Keptn",
			Description: "Learn the basics of Keptn and how to set up your first project",
			Category:    "Keptn Integrations",
			Tags:        []string{"beginner", "setup", "tutorial"},
			UpdatedAt:   "2026-01-20",
			FilePath:    "docs/getting-started-keptn.md",
			CoverImage:  "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=1200&h=400&fit=crop",
		},
		{
			ID:          "youtube-api-integration",
			Title:       "YouTube API Integration Guide",
			Description: "Learn how to integrate YouTube's Data API v3 into your applications",
			Category:    "API Integration",
			Tags:        []string{"youtube", "api", "video"},
			UpdatedAt:   "2026-01-22",
			FilePath:    "docs/youtube-api-integration.md",
			CoverImage:  "https://images.unsplash.com/photo-1611162616305-c69b3fa7fbe0?w=1200&h=400&fit=crop",
		},
		{
			ID:          "helm-chart-best-practices",
			Title:       "Helm Chart Best Practices",
			Description: "Create production-ready Helm charts that are maintainable and secure",
			Category:    "Helm Charts",
			Tags:        []string{"helm", "kubernetes", "best-practices"},
			UpdatedAt:   "2026-01-18",
			FilePath:    "docs/helm-chart-best-practices.md",
			CoverImage:  "https://images.unsplash.com/photo-1605745341075-1a6e8b9e7b8e?w=1200&h=400&fit=crop",
		},
		{
			ID:          "docker-optimization",
			Title:       "Docker Image Optimization",
			Description: "Best practices for creating smaller, faster, and more secure Docker images",
			Category:    "Container Images",
			Tags:        []string{"docker", "optimization", "security"},
			UpdatedAt:   "2026-01-21",
			FilePath:    "docs/docker-optimization.md",
			CoverImage:  "https://images.unsplash.com/photo-1605745341112-85968b19335b?w=1200&h=400&fit=crop",
		},
		{
			ID:          "kubernetes-networking",
			Title:       "Kubernetes Networking Deep Dive",
			Description: "Understanding Kubernetes networking model, services, and policies",
			Category:    "Kubernetes",
			Tags:        []string{"kubernetes", "networking", "advanced"},
			UpdatedAt:   "2026-01-19",
			FilePath:    "docs/kubernetes-networking.md",
			CoverImage:  "https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1200&h=400&fit=crop",
		},
		{
			ID:          "ci-cd-pipelines",
			Title:       "Modern CI/CD Pipelines",
			Description: "Build automated CI/CD pipelines with GitHub Actions, GitLab CI, and Jenkins",
			Category:    "DevOps",
			Tags:        []string{"cicd", "automation", "deployment"},
			UpdatedAt:   "2026-01-23",
			FilePath:    "docs/ci-cd-pipelines.md",
			CoverImage:  "https://images.unsplash.com/photo-1667372393119-3d4c48d07fc9?w=1200&h=400&fit=crop",
		},
		{
			ID:          "video-streaming-architecture",
			Title:       "Video Streaming Architecture",
			Description: "Building a scalable video streaming platform like YouTube or Netflix",
			Category:    "Architecture",
			Tags:        []string{"video", "streaming", "architecture"},
			UpdatedAt:   "2026-01-24",
			FilePath:    "docs/video-streaming-architecture.md",
			CoverImage:  "https://images.unsplash.com/photo-1574717024653-61fd2cf4d44d?w=1200&h=400&fit=crop",
		},
		{
			ID:          "microservices-patterns",
			Title:       "Microservices Design Patterns",
			Description: "Essential patterns for building resilient distributed systems",
			Category:    "Architecture",
			Tags:        []string{"microservices", "patterns", "distributed-systems"},
			UpdatedAt:   "2026-01-17",
			FilePath:    "docs/microservices-patterns.md",
			CoverImage:  "https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1200&h=400&fit=crop",
		},
		{
			ID:          "monitoring-observability",
			Title:       "Monitoring & Observability",
			Description: "Implement comprehensive monitoring with Prometheus, Grafana, and OpenTelemetry",
			Category:    "DevOps",
			Tags:        []string{"monitoring", "observability", "prometheus"},
			UpdatedAt:   "2026-01-16",
			FilePath:    "docs/monitoring-observability.md",
			CoverImage:  "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=1200&h=400&fit=crop",
		},
		{
			ID:          "content-delivery-networks",
			Title:       "Content Delivery Networks (CDN)",
			Description: "Optimize global content delivery with CDN strategies and best practices",
			Category:    "Infrastructure",
			Tags:        []string{"cdn", "performance", "caching"},
			UpdatedAt:   "2026-01-25",
			FilePath:    "docs/content-delivery-networks.md",
			CoverImage:  "https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1200&h=400&fit=crop",
		},
	}

	docs := make([]Document, 0, len(docsMetadata))

	for _, meta := range docsMetadata {
		contentB64, err := loadDocumentContent(meta.FilePath)
		if err != nil {
			continue
		}

		docs = append(docs, Document{
			ID:          meta.ID,
			Title:       meta.Title,
			Description: meta.Description,
			Category:    meta.Category,
			Tags:        meta.Tags,
			UpdatedAt:   meta.UpdatedAt,
			ContentB64:  contentB64,
			CoverImage:  meta.CoverImage,
		})
	}

	return docs, nil
}

func loadDocumentContent(filePath string) (string, error) {
	file, err := templates.FS().Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(content), nil
}
