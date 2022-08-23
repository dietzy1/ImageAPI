package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/google/uuid"
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

// Generates a new key on signup
func (a Application) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to parse credentials")
		return
	}
	creds := core.Credentials{
		Username:     r.Form.Get("username"),
		Passwordhash: a.creds.Hash(r.Form.Get("password")),
		Key:          core.GenerateAPIKey(),
		Created:      time.Now(),
		Role:         3,
	}
	err = creds.Validate(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to validate credentials")
		return
	}
	//issue with the databasedriver
	err = a.dbauth.Signup(ctx, creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
}

func (a Application) Signin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to parse the html form")
		return
	}
	creds := core.Credentials{
		Username:     r.Form.Get("username"),
		Passwordhash: r.Form.Get("password"),
	}
	storedCreds, err := a.dbauth.Signin(ctx, creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to authenticate user")
		return
	}

	if !a.creds.CompareHash(storedCreds.Passwordhash, creds.Passwordhash) {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to signin")
		return
	}
	sessionToken := uuid.New().String()
	err = a.session.Set(ctx, sessionToken, creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to store session in redis")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Second * 180),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
}

func (a Application) Signout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	r.Cookies()
	r.Response.Cookies()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to get session cookie")
		return
	}
	fmt.Println(cookie.Value)
	err = a.session.Delete(ctx, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete session cookie")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
		MaxAge:   -1,
	})
}

func (a Application) Refresh(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode("Unable to get session cookie")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.session.Get(ctx, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Session does not exist in redis db")
		return
	}
	//ACCESS THE REDIS DATABASE AND FIND THE USERNAME AND UPDATE THE DELETE TIME
	err = a.session.Update(ctx, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update session cookie")
		return
	}
	//Update the cookie deletion time
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    cookie.Value,
		Expires:  time.Now().Add(time.Second * 180),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
}
