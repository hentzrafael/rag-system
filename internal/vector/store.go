package vector


import (
	"context"
	"github.com/hentzrafael/rag-system/pkg/types"
)

type Store interface {
	Upsert(ctx context.Context, chunks []types.Chunk) error
	Search(ctx context.Context, vector []float32, k int) ([]types.Chunk, error)
}