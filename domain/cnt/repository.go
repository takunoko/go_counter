package cnt

import "context"

type CntRepository interface {
	Set(ctx context.Context, key string, val int) error
	Get(ctx context.Context, key string) (int, error)
	CntUp(ctx context.Context, key string) (int, error)
}
