package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/gorilla/mux"
)

//Communicates with mongodb database via the interface DbKeyPort

// Generate a new key
func (a Application) UpdateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to get session cookie")
		return
	}
	username, err := a.session.Get(ctx, cookie.Value)
	if err != nil || username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to get username from session")
		return
	}
	newKey := core.GenerateAPIKey()

	err = a.dbauth.StoreKey(ctx, newKey, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to add key to database")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a Application) DeleteKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to get session cookie")
		return
	}
	username, err := a.session.Get(ctx, cookie.Value)
	if err != nil || username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to get username from session")
		return
	}

	err = a.dbauth.DeleteKey(ctx, username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete key")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a Application) AuthenticateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) bool {
	vars := mux.Vars(r)
	key := vars["key"]
	if vars["key"] == "" {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode("Unable to parse the html form")
			return false
		}
		key = r.Form.Get("key")
	}

	_, err := a.session.Get(ctx, key)
	if err == nil {
		return true
	}
	result, ok := a.dbauth.AuthenticateKey(ctx, key)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode("Unable to authenticate key")
		return false
	}

	err = a.session.Set(ctx, key, result)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode("Unable to authenticate key")
		return false
	}
	return true
}

func (a Application) ShowKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to get session cookie")
		return
	}
	username, err := a.session.Get(ctx, cookie.Value)
	if err != nil || username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to get username from session")
		return
	}
	key, err := a.dbauth.GetKey(ctx, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to get key from database")
		return
	}
	//return the key to the user
	_ = json.NewEncoder(w).Encode(key)
}
