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
	//query := strings.Split(strings.ReplaceAll(vars["tags"], " ", ""), ",")
	query := strings.Split(strings.ReplaceAll(strings.ToLower(vars["tags"]), " ", ""), ",")

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
	s.keyauth.UpdateKey(ctx, w, r)
}

func (s *ServerAdapter) generateAdminAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.UpdateKey(ctx, w, r)
}

func (s *ServerAdapter) deleteAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.DeleteKey(ctx, w, r)
}

func (s *ServerAdapter) showAPIKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.keyauth.ShowKey(ctx, w, r)
}

func (s *ServerAdapter) signup(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Signup(ctx, w, r)
}

func (s *ServerAdapter) signin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Signin(ctx, w, r)
}

func (s *ServerAdapter) signout(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Signout(ctx, w, r)
}

func (s *ServerAdapter) refresh(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.Refresh(ctx, w, r)
}

func (s *ServerAdapter) deleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.accauth.DeleteAccount(ctx, w, r)
}
