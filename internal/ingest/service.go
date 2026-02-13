package ingest

import (
	"context"

	"github.com/hentzrafael/rag-system/internal/llm"
	"github.com/hentzrafael/rag-system/internal/vector"
	"github.com/hentzrafael/rag-system/pkg/types"
)

type Service struct {
	Chunker *Chunker
	Embedder llm.Embedder
	VectorStore vector.Store
}

func NewService(
	chunker *Chunker,
	embedder llm.Embedder,
	store vector.Store,
) *Service {
	return &Service{
		Chunker:  chunker,
		Embedder: embedder,
		VectorStore:    store,
	}
}

func (s *Service) Ingest(ctx context.Context, doc types.Document) error {
	chunks, err := s.Chunker.Split(doc)

	if err != nil {
		return err
	}

	texts := make([]string, len(chunks))
	for i := range chunks {
		texts[i] = chunks[i].Content
	}

	vectors, err := s.Embedder.Embed(ctx, texts)
	if err != nil {
		return err
	}

	for i := range chunks {
		chunks[i].Vector = vectors[i]
	}

	return s.VectorStore.Upsert(ctx, chunks)
}