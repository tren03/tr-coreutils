package ratelimiter

import (
	"fmt"

	"github.com/tren03/tr-coreutils/tr-ratelimiter/algorithms"
)

/*
We will not be implementing this for now
X-RateLimit-Limit - Max request client can make within a window
X-RateLimit-Remaining - Number of request remaining in current window
X-RateLimit-Reset - Rate limit reset time in epoch

We will be implementing this for now
Retry-After - Number of seconds client has to wait, to retry again - Only present when the user hits the rate limit
X-RateLimit-Remaining - Number of request remaining in current window
*/

type RatelimterResponse struct {
	Allow      bool
	RetryAfter *int64
	Reminaing  int
}

func (r RatelimterResponse) String() string {
	retry := "N/A"
	if r.RetryAfter != nil {
		retry = fmt.Sprintf("%ds", *r.RetryAfter)
	}

	return fmt.Sprintf(
		"Rate Limit Status:\n"+
			"  Allowed:   %t\n"+
			"  Remaining: %d\n"+
			"  Retry In:  %s",
		r.Allow, r.Reminaing, retry,
	)
}

type IRatelimiter interface {
	Allow(key string) RatelimterResponse
}

type Ratelimiter struct {
	alg algorithms.IAlgorithm
}

func NewRateLimiter(alg algorithms.IAlgorithm) *Ratelimiter {
	return &Ratelimiter{
		alg: alg,
	}
}

func (r *Ratelimiter) Allow(key string) RatelimterResponse {
	response := r.alg.Execute(key)

	return RatelimterResponse{
		Allow:      response.Allow,
		RetryAfter: response.RetryAfter,
		Reminaing:  response.Remaining,
	}
}
