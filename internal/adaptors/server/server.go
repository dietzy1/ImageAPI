package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/gorilla/mux"
)

type ServerAdapter struct {
	api     ports.ApiPort
	router  http.Handler
	accauth ports.AccAuthPort
	keyauth ports.KeyAuthPort
}

// Constructor
func NewServerAdapter(api ports.ApiPort, accauth ports.AccAuthPort, keyauth ports.KeyAuthPort) *ServerAdapter {
	return &ServerAdapter{api: api, accauth: accauth, keyauth: keyauth}
}

// Wrapper for router object
func (s *ServerAdapter) Router() http.Handler {
	return s.router
}

// Routering and paths
func Router(s *ServerAdapter) {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", s.healthcheck)

	//Authentication subrouter routes
	au := r.PathPrefix("/auth").Subrouter()
	au.Use(s.loggingMiddleware)
	au.Use(s.corsMiddlewareCookie)
	au.Use(s.ipRateLimitingMiddleware)

	//Login logout ROUTES
	au.HandleFunc("/signin/", s.signin).Methods(http.MethodPost)
	au.HandleFunc("/signup/", s.signup).Methods(http.MethodPost)
	au.HandleFunc("/signout/", s.signout).Methods(http.MethodPost)
	au.HandleFunc("/refresh/", s.refresh).Methods(http.MethodPost)
	au.HandleFunc("/deletekey/", s.deleteAPIKey).Methods(http.MethodPost)
	au.HandleFunc("/deleteaccount/", s.deleteAccount).Methods(http.MethodPost)

	//Key generation routes
	au.HandleFunc("/generatekey/", s.generateAPIKey).Methods(http.MethodGet)
	au.HandleFunc("/generateadminkey", s.generateAdminAPIKey).Methods(http.MethodGet)
	au.HandleFunc("/showkey/", s.showAPIKey).Methods(http.MethodGet)

	//Image subrouter routes
	sb := r.PathPrefix("/api/v0").Subrouter()
	sb.Use(s.loggingMiddleware)
	sb.Use(s.corsMiddleware)
	sb.Use(s.rateLimitingMiddleware)
	sb.Use(s.authenticateKey)

	//POST, PUT, DELETE ROUTES
	sb.HandleFunc("/image/", s.addImage).Methods(http.MethodPost)   //form data
	sb.HandleFunc("/image/", s.updateImage).Methods(http.MethodPut) //form data
	sb.HandleFunc("/image/{uuid}", s.deleteImage).Queries("key", "{key}").Methods(http.MethodDelete)

	//GET ROUTES // NEW ENDPOINTS // TODO
	//get multiple photos by tags -- Query --tags && quantity && key
	sb.HandleFunc("/images/tags/{tags}", s.findImagesTags).Queries("key", "{key}").Methods(http.MethodGet)
	//"quantity", "{quantity}", is optional

	//Get multiple random images -- Query -- quantity && key
	sb.HandleFunc("/images/random/", s.findImagesRandom).Queries("key", "{key}").Methods(http.MethodGet)
	//"quantity", "{quantity}", is optional

	// Get single image by ID	-- Query -- uuid && key
	sb.HandleFunc("/image/uuid/{uuid}", s.findImageUuid).Queries("key", "{key}").Methods(http.MethodGet)

	//Get single random image -- Query -- key -- confirmed working
	sb.HandleFunc("/image/random/", s.findImageRandom).Queries("key", "{key}").Methods(http.MethodGet)

	srv := &http.Server{ //&http.Server
		Handler:      r,
		Addr:         "0.0.0.0:" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	//rl := rateLimiting{}
	//Instantiate garbage collection of the cooldowns map
	//go rl.gc()
	log.Fatal(srv.ListenAndServe())

	s.router = r
}
