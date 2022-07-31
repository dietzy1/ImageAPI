package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/gorilla/mux"
)

type ServerAdapter struct {
	api            ports.ApiPort
	router         http.Handler
	authentication ports.AuthenticationPort
}

//helper function
func analyzeInterface(mp ports.ApiPort) {
	fmt.Printf("Interface type: %T\n", mp)
	fmt.Printf("Interface value: %v\n", mp)
	fmt.Printf("Interface is nil: %t\n", mp == nil)
}

//Constructor
func NewServerAdapter(api ports.ApiPort, authentication ports.AuthenticationPort) *ServerAdapter {
	return &ServerAdapter{api: api, authentication: authentication}
}

//Wrapper for router object
func (s *ServerAdapter) Router() http.Handler {
	return s.router
}

//Routering and paths
func Router(s *ServerAdapter) {
	r := mux.NewRouter()
	//subrouter

	//API key generation ROUTES
	r.HandleFunc("/generatekey/", s.generateAPIKey).Methods(http.MethodGet)
	r.HandleFunc("/generateadminkey", s.generateAdminAPIKey).Methods(http.MethodGet)

	//Attach a subrouter to the main router for authentication paths
	//au := r.PathPrefix("/auth").Subrouter()
	//Login logout ROUTES
	r.HandleFunc("/signin/", s.signin).Methods(http.MethodPost)
	r.HandleFunc("/signup/", s.signup).Methods(http.MethodPost)

	sb := r.PathPrefix("/api/v0").Subrouter()
	//Applies middleware to all subrouters
	sb.Use(s.loggingMiddleware)
	sb.Use(s.corsMiddleware)
	sb.Use(s.rateLimitingMiddleware)
	//sb.Use(s.authenticateKey)

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

//Entry point for http calls
func (s *ServerAdapter) findImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.FindImage(ctx, w, r)
}

//Entry point for http calls
func (s *ServerAdapter) findImages(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.FindImages(ctx, w, r)
}

//Entry point for http calls
func (s *ServerAdapter) addImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.AddImage(ctx, w, r)
}

//Entry point for http calls
func (s *ServerAdapter) updateImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.UpdateImage(ctx, w, r)
}

//Entry point for http calls
func (s *ServerAdapter) deleteImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.DeleteImage(ctx, w, r)
}

func (s *ServerAdapter) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func (s *ServerAdapter) generateAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.AddKey(ctx, w, r)
}

func (s *ServerAdapter) generateAdminAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.AddKey(ctx, w, r)
}

func (s *ServerAdapter) deleteAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.DeleteKey(ctx, w, r)
}

func (s *ServerAdapter) signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup called")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.Signup(ctx, w, r)
}

func (s *ServerAdapter) signin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.Signin(ctx, w, r)
}
