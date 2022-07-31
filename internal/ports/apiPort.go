package ports

import (
	"context"
	"net/http"
)

//Potentially move the ports into the api layer

//implement http server interface methods
type ApiPort interface {
	FindImage(ctx context.Context, w http.ResponseWriter, r *http.Request)
	FindImages(ctx context.Context, w http.ResponseWriter, r *http.Request)
	AddImage(ctx context.Context, w http.ResponseWriter, r *http.Request)
	DeleteImage(ctx context.Context, w http.ResponseWriter, r *http.Request)
	UpdateImage(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type AuthenticationPort interface {
	AddKey(ctx context.Context, w http.ResponseWriter, r *http.Request)
	DeleteKey(ctx context.Context, w http.ResponseWriter, r *http.Request)
	AuthenticateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) bool
	Signup(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signin(ctx context.Context, w http.ResponseWriter, r *http.Request)
}
