package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/google/uuid"
)

//I need to create 3 different roles 1 and 2 where 1 is admin priviledge and 2 is a normal user
//And then routes needs to be restricted based on the role someone has
//Only one who should be able to generate an admin account is an admin

//1 == root generate other admin accounts
//2 == admin acc has access to all routes except root
//3 == normal user only has access to get endpoints and basic authentication.

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
		Created_At:   a.image.SetTime(),
		Role:         3,
	}
	err = creds.Validate(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to validate credentials")
		return
	}
	//issue with the databasedriver
	err = a.dbAccAuth.Signup(ctx, creds)
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
	storedCreds, err := a.dbAccAuth.Signin(ctx, creds.Username)
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

// Take the session token from the cookie and use that
func (a Application) DeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	username, err := a.session.Get(ctx, cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Session does not exist in redis db")
		return
	}
	err = a.dbAccAuth.DeleteAccount(ctx, username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("Unable to delete account in the database")
		return
	}
	_ = json.NewEncoder(w).Encode("Account details has succesfully been deleted")
}
