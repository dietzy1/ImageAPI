package adapter

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/gorilla/mux"
)

type ServerAdapter struct {
	api    ports.ApiPort
	router http.Handler
	//rl     rateLimiting
}

//helper function
func analyzeInterface(mp ports.ApiPort) {
	fmt.Printf("Interface type: %T\n", mp)
	fmt.Printf("Interface value: %v\n", mp)
	fmt.Printf("Interface is nil: %t\n", mp == nil)
}

//Constructor
func NewServerAdapter(api ports.ApiPort) *ServerAdapter {
	return &ServerAdapter{api: api}
}

//Wrapper for router object
func (s *ServerAdapter) Router() http.Handler {
	return s.router
}

//Routering and paths
func Router(s *ServerAdapter) {
	r := mux.NewRouter()
	//subrouter
	sb := r.PathPrefix("/api/v0").Subrouter()
	//Applies middleware to all subrouters
	sb.Use(s.loggingMiddleware)
	sb.Use(s.corsMiddleware)
	//sb.Use(s.rl.rateLimitingMiddleware)

	sb.HandleFunc("/healthcheck", s.healthcheck)

	//POST, PUT, DELETE ROUTES
	sb.HandleFunc("/image/", s.addImage).Methods(http.MethodPost) //form data
	sb.HandleFunc("/image/", s.updateImage).Queries("uuid", "{uuid}").Methods(http.MethodPut)
	sb.HandleFunc("/image/", s.deleteImage).Queries("uuid", "{uuid}").Methods(http.MethodDelete)

	//GET ROUTES
	sb.HandleFunc("/image/", s.findImage).Queries("tag", "{tag}", "uuid", "{uuid}").Methods(http.MethodGet)
	sb.HandleFunc("/images/", s.findImages).Queries("tags", "{tags}", "quantity", "{quantity}").Methods(http.MethodGet)

	//Fileserver
	fs := http.FileServer(http.Dir("../image-folder"))
	r.PathPrefix("/fileserver/").Handler(http.StripPrefix("/fileserver/", fs)).Methods(http.MethodGet)

	srv := &http.Server{ //&http.Server
		Handler:      r,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	s.router = r

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

//Entry point for http calls
func (s *ServerAdapter) findImage(w http.ResponseWriter, r *http.Request) {
	s.api.FindImage(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) findImages(w http.ResponseWriter, r *http.Request) {
	s.api.FindImages(w, r)

}

//Entry point for http calls
func (s *ServerAdapter) addImage(w http.ResponseWriter, r *http.Request) {
	s.api.AddImage(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) updateImage(w http.ResponseWriter, r *http.Request) {
	s.api.UpdateImage(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) deleteImage(w http.ResponseWriter, r *http.Request) {
	s.api.DeleteImage(w, r)
}

func (s *ServerAdapter) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
