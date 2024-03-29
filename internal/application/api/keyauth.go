package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/gorilla/mux"
)

//This file is responcible for auth of API keys.
//Implements methods on the type KeyPortAuth interface.
//The methods are called from the handlers layer.

// Keys are autogenerated upon acc creation. The session token is used for verification to generate a fresh key.
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

	err = a.dbKeyAuth.StoreKey(ctx, newKey, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to add key to database")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Uses the session token for verification and sets the users key field to empty string.
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

	err = a.dbKeyAuth.DeleteKey(ctx, username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete key")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Uses the appended key query parameter and checks if the key exist in the caching layer or database.
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
	result, ok := a.dbKeyAuth.AuthenticateKey(ctx, key)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	err = a.session.Set(ctx, key, result)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return false
	}
	return true
}

// returns the user uuid of the API key. Used for image uploading and updating/deletion.
func (a Application) FindOwner(ctx context.Context, w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	key := vars["key"]
	if vars["key"] == "" {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode("Unable to parse the html form")
			return "", err
		}
		key = r.Form.Get("key")
	}

	_, err := a.session.Get(ctx, key)
	if err == nil {
		return "", err
	}

	return a.dbKeyAuth.GetUserUUID(ctx, key)
}

// Returns the API key for the user.
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
	key, err := a.dbKeyAuth.GetKey(ctx, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to get key from database")
		return
	}
	//return the key to the user
	_ = json.NewEncoder(w).Encode(key)
}
