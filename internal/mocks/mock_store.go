package mocks

import (
	"context"
	"github.com/hentzrafael/rag-system/pkg/types"
)

type InMemoryStore struct {
	data []types.Chunk
}

func (s *InMemoryStore) Upsert(ctx context.Context, chunks []types.Chunk) error {
	s.data = append(s.data, chunks...)
	return nil
}

func (s *InMemoryStore) Search(ctx context.Context, vector []float32, k int) ([]types.Chunk, error) {
	if len(s.data) < k {
		return s.data, nil
	}
	return s.data[:k], nil
}
