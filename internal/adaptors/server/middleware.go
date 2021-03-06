package server

import (
	"log"
	"net/http"
)

//Authenticate API key
func (s *ServerAdapter) authenticateKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//s.Authentication.AuthenticateKey(w, r)
		next.ServeHTTP(w, r)
	})
}

//Logging middleware
func (s *ServerAdapter) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

//Apply CORS headers //IDK what the fuck this actually does but its needed to load images on javascript front
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

//Rate limiting middleware
/* func (rl *rateLimiting) rateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//key := mux.Vars(r)["key"]
		key := "1"
		rl.c = rl.validateKey(key)
		if rl.c.Allow() != true {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
} */

//Authentication middleware
/* func (s *ServerAdapter) authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
} */

/* var cooldown = make(map[string]*rate.Limiter)
var mu sync.Mutex */

/* type rateLimiting struct {
	c  *rate.Limiter
	mu sync.Mutex
}

func (rl *rateLimiting) validateKey(key string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cd := make(map[string]*rate.Limiter)
	limiter, exists := cd[key]
	if !exists {
		limiter = rate.NewLimiter(1, 1)
		cd[key] = limiter
	}
	return limiter
} */
