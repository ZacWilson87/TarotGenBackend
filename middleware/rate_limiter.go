package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimiterConfig holds the configuration for rate limiting
type RateLimiterConfig struct {
	Rate       int           // Number of requests
	Period     time.Duration // Time period
	Identifier string        // Identifier for the rate limiter (e.g., "ip", "user")
}

// NewRateLimiter creates a new rate limiter middleware
func NewRateLimiter(config RateLimiterConfig) mux.MiddlewareFunc {
	// Create a rate limiter with memory store
	store := memory.NewStore()
	rate := limiter.Rate{
		Period: config.Period,
		Limit:  int64(config.Rate),
	}
	instance := limiter.New(store, rate)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the identifier (IP address by default)
			identifier := r.RemoteAddr
			if config.Identifier == "user" {
				// TODO: Implement user-based identification when auth is added
				identifier = "anonymous"
			}

			// Get the context for the identifier
			context, err := instance.Get(r.Context(), identifier)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Set rate limit headers
			w.Header().Set("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
			w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
			w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

			// Check if the request is allowed
			if context.Reached {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// DefaultRateLimiter returns a rate limiter with default settings
func DefaultRateLimiter() mux.MiddlewareFunc {
	config := RateLimiterConfig{
		Rate:       100,       // 100 requests
		Period:     time.Hour, // per hour
		Identifier: "ip",      // limit by IP address
	}
	return NewRateLimiter(config)
}
