package main

import (
	"github.com/hentzrafael/rag-system/internal/ingest"
	"github.com/hentzrafael/rag-system/internal/mocks"
	"github.com/hentzrafael/rag-system/pkg/types"
	"context"
	"log"
)
func main(){
	log.Println("Starting RAG system")
	chunker := &ingest.Chunker{ChunkSize: 100}
	embedder := &mocks.MockEmbedder{}
	store := &mocks.InMemoryStore{}

	service := ingest.NewService(chunker, embedder, store)

	doc := types.Document{
		ID: "1",
		Content: "This is a test document",
		Metadata: map[string]string{
			"source": "test",
		},
	}

	ctx := context.Background()

	err := service.Ingest(ctx, doc)
	if err != nil {
		log.Fatalf("Error ingesting document: %v", err)
	}

	log.Println("Document ingested successfully")
}