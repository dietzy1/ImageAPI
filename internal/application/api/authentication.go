package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

//Communicates with mongodb database via the interface DbKeyPort

//Generate a key
func (a Application) AddKey(w http.ResponseWriter, r *http.Request) {
	a.creds.Key = core.GenerateAPIKey()

	err := a.dbauth.StoreKey(a.creds.Key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add key to database")
		return
	}

	//Save to db

}

func (a Application) DeleteKey(w http.ResponseWriter, r *http.Request) {

}

func (a Application) AuthenticateKey(w http.ResponseWriter, r *http.Request) {
	//Perform a check vs the database if the provided key exists in the database

}

func (a Application) Signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
	creds := core.Credentials{
		Username:     r.Form.Get("username"),
		Passwordhash: r.Form.Get("password"),
		Key:          core.GenerateAPIKey(),
		Created:      time.Now(),
		Role:         3,
	}
	//Need to validate the fields here

	err = a.dbauth.Signup(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
}

func (a Application) Signin(w http.ResponseWriter, r *http.Request) {

}

type Authentication interface {
	AddKey(w http.ResponseWriter, r *http.Request)
	DeleteKey(w http.ResponseWriter, r *http.Request)
	AuthenticateKey(w http.ResponseWriter, r *http.Request)
	signup(w http.ResponseWriter, r *http.Request)
	signin(w http.ResponseWriter, r *http.Request)
}
