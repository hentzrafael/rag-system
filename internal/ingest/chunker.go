package ingest

import (
	"strings"
	"github.com/hentzrafael/rag-system/pkg/types"
)

type Chunker struct {
	ChunkSize int
}

func (c *Chunker) Split(doc types.Document) ([]types.Chunk, error) {
	words := strings.Fields(doc.Content)

	var chunks []types.Chunk
	for i:= 0; i < len(words); i += c.ChunkSize {
		end := i + c.ChunkSize

		if end > len(words) {
			end = len(words)
		}

		content := strings.Join(words[i:end]," ")

		chunks = append(chunks, types.Chunk{
			ID: "random-uuid",
			DocumentID: doc.ID,
			Content: content,
			Metadata: doc.Metadata,
		})
	}
	return chunks, nil
}