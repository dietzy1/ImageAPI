package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/gorilla/mux"
)

//Communicates with mongodb database via the interface DbKeyPort

//Generate a new key
func (a Application) AddKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	a.creds.Key = core.GenerateAPIKey()

	err := a.dbauth.StoreKey(a.creds.Key, a.creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add key to database")
		return
	}

	//Save to db

}

func (a Application) DeleteKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.dbauth.DeleteKey(vars["key"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete key")
		return
	}
}

func (a Application) AuthenticateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//Perform a check vs the database if the provided key exists in the database

}

//Generates a new key on signup
func (a Application) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}

	creds := core.Credentials{
		Username:     r.Form.Get("username"),
		Passwordhash: r.Form.Get("password"),
		//Need to implement hashing on the password
		Key:     core.GenerateAPIKey(),
		Created: time.Now(),
		Role:    3,
	}
	err = creds.Validate(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to validate credentials")
		return
	}

	err = a.dbauth.Signup(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
}

func (a Application) Signin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add credentials while parsing")
		return
	}
	creds := core.Credentials{
		Username:     r.Form.Get("username"),
		Passwordhash: r.Form.Get("password"),
	}
	err = a.dbauth.Signin(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to log in with credentials")
		return
	}

}
