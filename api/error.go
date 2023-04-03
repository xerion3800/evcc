package api

import "errors"

// ErrNotAvailable indicates that a feature is not available
var ErrNotAvailable = errors.New("not available")

// ErrMustRetry indicates that a rate-limited operation should be retried
var ErrMustRetry = errors.New("must retry")

// ErrSponsorRequired indicates that a sponsor token is required
var ErrSponsorRequired = errors.New("sponsorship required, see https://github.com/xerion3800/evcc#sponsorship")

// ErrMissingCredentials indicates that user/password are missing
var ErrMissingCredentials = errors.New("missing credentials")

// ErrOutdated indicates that result is outdated
var ErrOutdated = errors.New("outdated")

// ErrTimeout is the error returned when a timeout happened.
// Modeled after context.DeadlineError
var ErrTimeout error = errTimeoutError{}

type errTimeoutError struct{}

func (errTimeoutError) Error() string   { return "timeout" }
func (errTimeoutError) Timeout() bool   { return true }
func (errTimeoutError) Temporary() bool { return true }
