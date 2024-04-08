package transactor

import "context"

type (
	Key struct{}

	Func func(ctx context.Context) (interface{}, error)
)
