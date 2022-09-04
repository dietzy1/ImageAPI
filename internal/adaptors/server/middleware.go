package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

// API key authentication middleware
func (s *ServerAdapter) authenticateKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if !s.keyauth.AuthenticateKey(ctx, w, r) {
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

// Apply CORS headers //IDK what the fuck this actually does but its needed to load images on javascript front
func (s *ServerAdapter) corsMiddlewareCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://pepe-api.vercel.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

//var cooldown = make(map[string]*rate.Limiter)

var cooldown = make(map[string]*rateLimiting)

type rateLimiting struct {
	c          *rate.Limiter
	mu         sync.Mutex
	expiration time.Time
}

// Rate limiting middleware
func (s *ServerAdapter) rateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]
		rl := rateLimiting{}
		rl.c = rl.ratelimit(key)
		if !rl.c.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			log.Default().Printf("Rate limit exceeded for key %s", key)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (rl *rateLimiting) ratelimit(rateString string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	limiter, exists := cooldown[rateString]
	if !exists {
		limiter = &rateLimiting{
			c:          rate.NewLimiter(1, 5),
			expiration: time.Now().Add(1 * time.Hour),
		}
		cooldown[rateString] = limiter

	}
	return limiter.c
}

/* func (rl *rateLimiting) ratelimit(rateString string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	limiter, exists := cooldown[rateString]
	if !exists {
		limiter = rate.NewLimiter(1, 5) //Still need to configure the exact rate limit
		cooldown[rateString] = limiter
	}
	return limiter
} */

func (s *ServerAdapter) ipRateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl := rateLimiting{}
		ip := getIP(r)
		rl.c = rl.ratelimit(ip)
		if !rl.c.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			log.Default().Printf("Rate limit exceeded for ip %s", ip)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// garbage collector for the rate limiter
func (rl *rateLimiting) gc() {
	var newCooldown = make(map[string]*rateLimiting)
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for k, v := range cooldown {
		if v.expiration.Unix() < time.Now().Unix() {
			delete(cooldown, k)
		} else {
			newCooldown[k] = v
		}
	}
	cooldown = nil
	cooldown = newCooldown
	time.Sleep(1 * time.Hour)
	log.Println("Garbage collector initialized")
	rl.gc()
}
