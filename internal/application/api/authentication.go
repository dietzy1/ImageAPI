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

//Generate a new key
func (a Application) AddKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	a.creds.Key = core.GenerateAPIKey()

	err := a.dbauth.StoreKey(ctx, a.creds.Key, a.creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add key to database")
		return
	}
	//Save to db
}

func (a Application) DeleteKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.dbauth.DeleteKey(ctx, vars["key"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete key")
		return
	}
}

func (a Application) AuthenticateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) bool {
	vars := mux.Vars(r)
	_, err := a.session.Get(ctx, vars["key"])
	if err == nil {
		return true
	}
	result, ok := a.dbauth.AuthenticateKey(ctx, vars["key"])
	if ok != true {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to authenticate key")
		return false
	}
	err = a.session.Set(ctx, vars["key"], result)
	return true
}

//Generates a new key on signup
func (a Application) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add credentials while parsing")
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
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
}

func (a Application) Signin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signin called yepperrs")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println("Error parsing form")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to parse the html form")
		return
	}
	creds := core.Credentials{
		Username:     r.Form.Get("username"),
		Passwordhash: r.Form.Get("password"),
	}
	fmt.Println(creds)
	storedCreds, err := a.dbauth.Signin(ctx, creds.Username)
	fmt.Println(storedCreds)

	if a.creds.CompareHash(storedCreds.Passwordhash, creds.Passwordhash) != true {
		fmt.Println("Password incorrect")
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
		/* Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(time.Second * 180), */
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Second * 180),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
}

func (a Application) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to get session cookie")
		return
	}
	err = a.session.Delete(ctx, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete session cookie")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	})
}

func (a Application) ProtectPath(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	sessionToken := cookie.Value

	//Compare with the session token in the redis database
	username, err := a.session.Get(ctx, sessionToken)
	if err != nil || username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Session does not exist in redis db")
		return
	}
	//if err not equal to nil then the session token is valid

	//Set some variable to approved
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
}
