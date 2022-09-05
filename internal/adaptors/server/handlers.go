package server

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//This file contains the entrypoint handlers for the http server.
//Implements methods application interfaces.
//Type APiPort interface
//Type AccAuthPort interface
//Type KeyAuthPOrt interface

// Entry point for following GET route:
// api/v0/image/{uuid}
func (s *ServerAdapter) findImageUuid(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	vars := mux.Vars(r)
	query := vars["uuid"]

	s.api.FindImage(ctx, w, r, query, "uuid")
}

// Entry point for following GET route:
// api/v0/image/random/
func (s *ServerAdapter) findImageRandom(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.FindImage(ctx, w, r, "random", "random")
}

// Entry point for following GET route:
// api/v0/images/tags{}
func (s *ServerAdapter) findImagesTags(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	//query := strings.Split(strings.ReplaceAll(vars["tags"], " ", ""), ",")
	query := strings.Split(strings.ReplaceAll(strings.ToLower(vars["tags"]), " ", ""), ",")

	q := r.URL.Query()
	quantity, err := strconv.Atoi(strings.Join(q["quantity"], ""))
	if err != nil || quantity <= 0 { //<= 0 is a hack to allow for a default value
		quantity = 10
	}

	s.api.FindImages(ctx, w, r, query, "tags", quantity)
}

// Entry point for following GET route:
// api/v0/images/random/
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

// Entry point for following POST route:
// api/v0/image/
func (s *ServerAdapter) addImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.AddImage(ctx, w, r)
}

// Entry point for following PUT route:
// api/v0/image/
func (s *ServerAdapter) updateImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.UpdateImage(ctx, w, r)
}

// Entry point for following DELETE route:
// api/v0/image/{uuid}
func (s *ServerAdapter) deleteImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.api.DeleteImage(ctx, w, r)
}

// Entry point for following healthcheck route:
// /healthcheck
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

// Entry point for following GET route:
// /auth/generatekey/
func (s *ServerAdapter) generateAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.UpdateKey(ctx, w, r)
}

// Idk what this is doing here
func (s *ServerAdapter) generateAdminAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.UpdateKey(ctx, w, r)
}

// Entry point for following POST route:
// /auth/deletekey/
func (s *ServerAdapter) deleteAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.DeleteKey(ctx, w, r)
}

// Entry point for following GET route:
// /auth/showkey/
func (s *ServerAdapter) showAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.ShowKey(ctx, w, r)
}

// Entry point for following POST route:
// /auth/signup/
func (s *ServerAdapter) signup(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Signup(ctx, w, r)
}

// Entry point for following POST route:
// /auth/signin/
func (s *ServerAdapter) signin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Signin(ctx, w, r)
}

// Entry point for following POST route:
// /auth/signout/
func (s *ServerAdapter) signout(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Signout(ctx, w, r)
}

// Entry point for following POST route:
// /auth/refresh/
func (s *ServerAdapter) refresh(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Refresh(ctx, w, r)
}

// Entry point for following POST route:
// /auth/deleteaccount//
func (s *ServerAdapter) deleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.DeleteAccount(ctx, w, r)
}
