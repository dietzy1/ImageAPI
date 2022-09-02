package server

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/gorilla/mux"
)

type ServerAdapter struct {
	api            ports.ApiPort
	router         http.Handler
	authentication ports.AuthenticationPort
}

// helper function
/* func analyzeInterface(mp ports.ApiPort) {
	fmt.Printf("Interface type: %T\n", mp)
	fmt.Printf("Interface value: %v\n", mp)
	fmt.Printf("Interface is nil: %t\n", mp == nil)
} */

// Constructor
func NewServerAdapter(api ports.ApiPort, authentication ports.AuthenticationPort) *ServerAdapter {
	return &ServerAdapter{api: api, authentication: authentication}
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
		TLSConfig:    &tls.Config{},
	}
	rl := rateLimiting{}
	//Instantiate garbage collection of the cooldowns map
	go rl.gc()
	log.Fatal(srv.ListenAndServe())

	s.router = r
}

// Entry point for http calls
func (s *ServerAdapter) findImageUuid(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	vars := mux.Vars(r)
	query := vars["uuid"]
	querytype := "uuid"

	s.api.FindImage(ctx, w, r, query, querytype)
}

func (s *ServerAdapter) findImageRandom(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.FindImage(ctx, w, r, "random", "random")
}

// Entry point for http calls
func (s *ServerAdapter) findImagesTags(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	query := strings.Split(strings.ReplaceAll(vars["tags"], " ", ""), ",")

	q := r.URL.Query()
	quantity, err := strconv.Atoi(strings.Join(q["quantity"], ""))
	if err != nil || quantity <= 0 { //<= 0 is a hack to allow for a default value
		quantity = 10
	}

	s.api.FindImages(ctx, w, r, query, "tags", quantity)
}

// Entry point for http calls
func (s *ServerAdapter) findImagesRandom(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	q := r.URL.Query()
	quantity, err := strconv.Atoi(strings.Join(q["quantity"], ""))
	if err != nil || quantity <= 0 { //<= 0 is a hack to allow for a default value
		quantity = 10
	}

	s.api.FindImages(ctx, w, r, nil, "random", quantity)
}

// Entry point for http calls
func (s *ServerAdapter) addImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.AddImage(ctx, w, r)
}

// Entry point for http calls
func (s *ServerAdapter) updateImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.UpdateImage(ctx, w, r)
}

// Entry point for http calls
func (s *ServerAdapter) deleteImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.DeleteImage(ctx, w, r)
}

func (s *ServerAdapter) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	if r.Method == "OPTIONS" {
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func (s *ServerAdapter) generateAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.UpdateKey(ctx, w, r)
}

func (s *ServerAdapter) generateAdminAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.UpdateKey(ctx, w, r)
}

func (s *ServerAdapter) deleteAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.DeleteKey(ctx, w, r)
}

func (s *ServerAdapter) showAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.ShowKey(ctx, w, r)
}

func (s *ServerAdapter) signup(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.Signup(ctx, w, r)
}

func (s *ServerAdapter) signin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.Signin(ctx, w, r)
}

func (s *ServerAdapter) signout(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.Signout(ctx, w, r)
}

func (s *ServerAdapter) refresh(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.Refresh(ctx, w, r)
}

func (s *ServerAdapter) deleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.authentication.DeleteAccount(ctx, w, r)
}
