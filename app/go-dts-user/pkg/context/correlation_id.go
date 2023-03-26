package context

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string

const (
	corrId ContextKey = "X-Correlation-ID"
)

func (c ContextKey) String() string {
	return string(c)
}

func GetCorrelationID(ctx context.Context) (context.Context, string) {
	val := ctx.Value(corrId)
	if val != nil {
		vals, ok := val.(string)
		if !ok || vals == "" {
			vals = uuid.New().String()
			ctx = context.WithValue(ctx, corrId, vals)
		}
		return ctx, vals
	}
	vals := uuid.New().String()
	ctx = context.WithValue(ctx, corrId, vals)

	return ctx, vals
}
