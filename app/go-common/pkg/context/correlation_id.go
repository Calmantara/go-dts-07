package context

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type ContextKey string

const (
	CorrID ContextKey = "X-Correlation-ID"
)

func (c ContextKey) String() string {
	return string(c)
}

func GetCorrelationID(ctxIn context.Context) (ctx context.Context, vals string) {
	vals, ok := ctxIn.Value(CorrID.String()).(string)
	if !ok || vals == "" {
		vals = ulid.Make().String()
		ctx = context.WithValue(ctxIn, CorrID, vals)
		return ctx, vals
	}

	vals, ok = ctxIn.Value(CorrID.String()).(string)
	if !ok || vals == "" {
		vals = ulid.Make().String()
		ctx = context.WithValue(ctxIn, CorrID, vals)
	}

	return ctx, vals
}
