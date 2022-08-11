package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// API key authentication middleware
func (s *ServerAdapter) authenticateKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if !s.authentication.AuthenticateKey(ctx, w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Logging middleware
func (s *ServerAdapter) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Should prolly disable this shit later
// Apply CORS headers //IDK what the fuck this actually does but its needed to load images on javascript front
func (s *ServerAdapter) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Should prolly disable this shit later
// Apply CORS headers //IDK what the fuck this actually does but its needed to load images on javascript front
func (s *ServerAdapter) corsMiddlewareCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* 	w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		} */
		//w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Origin", "localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

var cooldown = make(map[string]*rate.Limiter)

type rateLimiting struct {
	c  *rate.Limiter
	mu sync.Mutex
}

// Rate limiting middleware
func (s *ServerAdapter) rateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//key := mux.Vars(r)["key"]
		key := "1"
		rl := rateLimiting{}
		rl.c = rl.ratelimitKey(key)
		if !rl.c.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			log.Default().Printf("Rate limit exceeded for key %s", key)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (rl *rateLimiting) ratelimitKey(key string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	limiter, exists := cooldown[key]
	if !exists {
		limiter = rate.NewLimiter(1%2, 1) //Still need to configure the exact rate limit
		cooldown[key] = limiter
	}
	return limiter
}
