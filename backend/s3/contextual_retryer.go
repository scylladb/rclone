package s3

import (
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/rclone/rclone/lib/pacer"
)

type contextualRetryer struct {
	client.DefaultRetryer
}

// ShouldRetry returns true if the request should be retried.
func (d contextualRetryer) ShouldRetry(r *request.Request) bool {
	if pacer.IsNoRetryCtx(r.Context()) {
		return false
	}

	// ShouldRetry returns false if number of max retries is 0.
	if d.NumMaxRetries == 0 {
		return false
	}

	// If one of the other handlers already set the retry state
	// we don't want to override it based on the service's state
	if r.Retryable != nil {
		return *r.Retryable
	}
	return r.IsErrorRetryable() || r.IsErrorThrottle()
}
