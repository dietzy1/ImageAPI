package ports

import (
	"context"
	"net/http"
)

//Potentially move the ports into the api layer

// implement http server interface methods
type ApiPort interface {
	FindImage(ctx context.Context, w http.ResponseWriter, r *http.Request, query string, querytype string)
	FindImages(ctx context.Context, w http.ResponseWriter, r *http.Request, query []string, querytype string, quantity int)
	AddImage(ctx context.Context, w http.ResponseWriter, r *http.Request, ownerUuid string)
	DeleteImage(ctx context.Context, w http.ResponseWriter, r *http.Request, ownerUuid string)
	UpdateImage(ctx context.Context, w http.ResponseWriter, r *http.Request, ownerUuid string)
}

type AccAuthPort interface {
	Signup(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signin(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signout(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Refresh(ctx context.Context, w http.ResponseWriter, r *http.Request)
	DeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type KeyAuthPort interface {
	UpdateKey(ctx context.Context, w http.ResponseWriter, r *http.Request)
	DeleteKey(ctx context.Context, w http.ResponseWriter, r *http.Request)
	AuthenticateKey(ctx context.Context, w http.ResponseWriter, r *http.Request) bool
	ShowKey(ctx context.Context, w http.ResponseWriter, r *http.Request)
	FindOwner(ctx context.Context, w http.ResponseWriter, r *http.Request) (string, error)
}

type EloSystemPort interface {
	RequestMatch(ctx context.Context, w http.ResponseWriter, r *http.Request)
	MatchResult(ctx context.Context, w http.ResponseWriter, r *http.Request)
	GetLeaderboard(ctx context.Context, w http.ResponseWriter, r *http.Request)
}
