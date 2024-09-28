package retry

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Retry retry an operation which returns an error using an exponential backoff strategy
func Retry(o func() error) error {
	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.InitialInterval = 500 * time.Millisecond
	expBackoff.MaxInterval = 2 * time.Second
	expBackoff.MaxElapsedTime = 5 * time.Second

	return backoff.Retry(o, expBackoff)
}
