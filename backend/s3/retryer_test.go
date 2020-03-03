package s3

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/rclone/rclone/lib/pacer"
)

func TestRetryerShouldRetry(t *testing.T) {
	retryer := retryer{client.DefaultRetryer{NumMaxRetries: 3}}
	conf := aws.NewConfig()
	op := &request.Operation{
		Name:       "GetCredentials",
		HTTPMethod: "GET",
	}
	retryable := true

	t.Run("retry", func(t *testing.T) {
		r := request.New(*conf, metadata.ClientInfo{}, request.Handlers{}, retryer, op, nil, nil)
		r.Retryable = &retryable
		if !retryer.ShouldRetry(r) {
			t.Fatal("ShouldRetry = false, expected true")
		}
	})

	t.Run("no retry", func(t *testing.T) {
		r := request.New(*conf, metadata.ClientInfo{}, request.Handlers{}, retryer, op, nil, nil)
		r.Retryable = &retryable
		r.SetContext(pacer.WithRetries(context.Background(), 0))
		if retryer.ShouldRetry(r) {
			t.Fatal("ShouldRetry = true, expected false")
		}
	})
}
