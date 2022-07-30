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

func (a Application) AuthenticateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//Perform a check vs the database if the provided key exists in the database
	return
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

	if a.creds.CompareHash(storedCreds.Passwordhash, creds.Passwordhash) != true {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to signin")
		return
	}
	//User is verified at this point
	//err = a.StoreKey(ctx, creds.Key, creds.Username)

	//Call redis to store the key in a cache

	sessionToken := uuid.New().String() //Need to sign this with a secret
	ExpiresAt := time.Now().Add(time.Second * 180)

	//declare the object to be stored in the session
	session := session{
		username: creds.Username,
		expires:  ExpiresAt,
	}
	//Store the session in the redis cache

	err = a.session.Set(ctx, sessionToken, session)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to store session in redis")
		return
	}
	//fmt.Println(sessions)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: ExpiresAt,
	})
}
