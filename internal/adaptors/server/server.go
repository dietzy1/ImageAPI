package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/gorilla/mux"
)

//This file contains http routes and router configurations
//See handlers for interface method implementation.

// This is the main struct that contains all api && accauth && keyauth related methods
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

	//Account auth routes
	au.HandleFunc("/signin/", s.signin).Methods(http.MethodPost)
	au.HandleFunc("/signup/", s.signup).Methods(http.MethodPost)
	au.HandleFunc("/signout/", s.signout).Methods(http.MethodPost)
	au.HandleFunc("/refresh/", s.refresh).Methods(http.MethodPost)
	au.HandleFunc("/deleteaccount/", s.deleteAccount).Methods(http.MethodPost)

	//API key routes
	au.HandleFunc("/generatekey/", s.generateAPIKey).Methods(http.MethodGet)
	au.HandleFunc("/showkey/", s.showAPIKey).Methods(http.MethodGet)
	au.HandleFunc("/deletekey/", s.deleteAPIKey).Methods(http.MethodPost)

	//Image subrouter routes
	sb := r.PathPrefix("/api/v0").Subrouter()
	sb.Use(s.loggingMiddleware)
	sb.Use(s.corsMiddleware)
	sb.Use(s.rateLimitingMiddleware)
	sb.Use(s.authenticateKey)

	//Admin paths //Not accesible to average user
	//POST PATCH DELETE
	sb.HandleFunc("/image/", s.addImage).Methods(http.MethodPost)   //form data format
	sb.HandleFunc("/image/", s.updateImage).Methods(http.MethodPut) //form data format
	sb.HandleFunc("/image/{uuid}", s.deleteImage).Queries("key", "{key}").Methods(http.MethodDelete)

	//Normal user paths //Accessible to the average user

	//GET multiple images by tags // The API key should be appended as a query parameter //Quantity can be appended as an optional query parameter.
	sb.HandleFunc("/images/tags/{tags}", s.findImagesTags).Queries("key", "{key}").Methods(http.MethodGet)

	//Get multiple random images -- Query -- quantity && key
	//GET multiple random images The API key should be appended as a query parameter //Quantity can be appended as an optional query parameter.
	sb.HandleFunc("/images/random/", s.findImagesRandom).Queries("key", "{key}").Methods(http.MethodGet)

	//GET single image by ID // The API key should be appended as a query parameter
	sb.HandleFunc("/image/uuid/{uuid}", s.findImageUuid).Queries("key", "{key}").Methods(http.MethodGet)

	//Get single random image // The API key should be appended as a query parameter
	sb.HandleFunc("/image/random/", s.findImageRandom).Queries("key", "{key}").Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		//Addr:         "0.0.0.0:" + os.Getenv("PORT"), //Required addr for railway.app deployment.
		Addr: os.Getenv("ADDR") + os.Getenv("PORT"), //Required addr for railway.app deployment.
		//Addr:         "localhost:8000",
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
