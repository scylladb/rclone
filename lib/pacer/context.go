package pacer

import "context"

// ctxt is a context key type.
type ctxt byte

// ctxt enumeration.
const (
	ctxNoRetry ctxt = iota
)

func IsNoRetryCtx(ctx context.Context) bool {
	if _, ok := ctx.Value(ctxNoRetry).(bool); ok {
		return true
	}
	return false
}

// WithNoRetry disables retries.
func WithNoRetry(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxNoRetry, true)
}
