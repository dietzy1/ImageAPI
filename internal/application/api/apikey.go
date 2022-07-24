package api

import "net/http"

//Communicates with mongodb database via the interface DbKeyPort

func (a Application) AddKey(w http.ResponseWriter, r *http.Request) {

	return
}

func (a Application) DeleteKey(w http.ResponseWriter, r *http.Request) {
	return
}

func (a Application) AuthenticateKey(w http.ResponseWriter, r *http.Request) {
	//Perform a check vs the database if the provided key exists in the database
	return
}
