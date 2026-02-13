package mocks

import (
	"context"
)
type MockEmbedder struct{}

func (m *MockEmbedder) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	vectors := make([][]float32, len(texts))
	for i := range texts {
		vectors[i] = []float32{0.1, 0.2, 0.3} // fake vector
	}
	return vectors, nil
}
