package documents

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "markitos-it-svc-documents/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	docsServiceAddr = os.Getenv("DOCS_SERVICE_ADDR")
)

func init() {
	if docsServiceAddr == "" {
		docsServiceAddr = "localhost:8888"
	}
	log.Printf("📡 Documents service address: %s", docsServiceAddr)
}

// GetAllDocumentsFromService fetches all documents from the gRPC service
func GetAllDocumentsFromService() ([]Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		docsServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to documents service: %w", err)
	}
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)
	resp, err := client.GetAllDocuments(ctx, &pb.GetAllDocumentsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch documents: %w", err)
	}

	docs := make([]Document, len(resp.Documents))
	for i, pbDoc := range resp.Documents {
		docs[i] = Document{
			ID:          pbDoc.Id,
			Title:       pbDoc.Title,
			Description: pbDoc.Description,
			Category:    pbDoc.Category,
			Tags:        pbDoc.Tags,
			UpdatedAt:   pbDoc.UpdatedAt.AsTime().Format("2006-01-02"),
			ContentB64:  pbDoc.ContentB64,
			CoverImage:  pbDoc.CoverImage,
		}
	}

	return docs, nil
}

// GetDocumentByIdFromService fetches a single document by ID from the gRPC service
func GetDocumentByIdFromService(id string) (*Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		docsServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to documents service: %w", err)
	}
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)
	resp, err := client.GetDocumentById(ctx, &pb.GetDocumentByIdRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch document: %w", err)
	}

	doc := &Document{
		ID:          resp.Document.Id,
		Title:       resp.Document.Title,
		Description: resp.Document.Description,
		Category:    resp.Document.Category,
		Tags:        resp.Document.Tags,
		UpdatedAt:   resp.Document.UpdatedAt.AsTime().Format("2006-01-02"),
		ContentB64:  resp.Document.ContentB64,
		CoverImage:  resp.Document.CoverImage,
	}

	return doc, nil
}
