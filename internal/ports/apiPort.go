package ports

import (
	"net/http"
)

//Potentially move the ports into the api layer

//implement http server interface methods
type ApiPort interface {
	FindImage(w http.ResponseWriter, r *http.Request)
	FindImages(w http.ResponseWriter, r *http.Request)
	AddImage(w http.ResponseWriter, r *http.Request)
	DeleteImage(w http.ResponseWriter, r *http.Request)
	UpdateImage(w http.ResponseWriter, r *http.Request)
}

type ApiKeyPort interface {
	AddKey(w http.ResponseWriter, r *http.Request)
	DeleteKey(w http.ResponseWriter, r *http.Request)
	AuthenticateKey(w http.ResponseWriter, r *http.Request)
}
