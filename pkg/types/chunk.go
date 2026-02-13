package types

type Chunk struct {
	ID string
	DocumentID string
	Content string
	Vector []float32
	Metadata map[string]string
}