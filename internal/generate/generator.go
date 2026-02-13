package generate

import "context"

type Generator interface {
	Stream(ctx context.Context, prompt string) (<-chan string, error)
}