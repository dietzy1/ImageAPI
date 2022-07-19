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

//TODO need to implement middlewares

//Routering and paths
func Router(s *ServerAdapter) {
	r := mux.NewRouter()
	//subrouter
	sb := r.PathPrefix("/api/v0").Subrouter()
	//Applies middleware to all subrouters
	sb.Use(s.loggingMiddleware)
	sb.Use(s.corsMiddleware)

	sb.HandleFunc("/healthcheck", s.healthcheck)

	//POST, PUT, DELETE ROUTES
	sb.HandleFunc("/image/", s.addImage).Methods(http.MethodPost) //form data
	sb.HandleFunc("/image/{uuid}", s.updateImage).Queries("uuid", "{uuid}").Methods(http.MethodPut)
	sb.HandleFunc("/image/{uuid}", s.deleteImage).Queries("uuid", "{uuid}").Methods(http.MethodDelete)

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
