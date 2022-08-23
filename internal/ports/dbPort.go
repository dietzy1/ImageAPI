package ports

import (
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

//Potentially move the ports into the database folder

// implement the mongodb interface methods
type DbPort interface {
	FindImage(ctx context.Context, querytype string, query string) (*core.Image, error)
	FindImages(ctx context.Context, querytype string, query []string, quantity int) ([]core.Image, error)
	StoreImage(ctx context.Context, image *core.Image) error
	UpdateImage(ctx context.Context, image *core.Image) error
	DeleteImage(ctx context.Context, uuid string) error
}

// Need to implement mongodb methods on this interface
type DbAuthenticationPort interface {
	StoreKey(ctx context.Context, newKey string, username string) error
	DeleteKey(ctx context.Context, username string) error
	AuthenticateKey(ctx context.Context, key string) (string, bool)
	GetKey(ctx context.Context, userrname string) (string, error)
	Signup(ctx context.Context, creds core.Credentials) error
	Signin(ctx context.Context, username string) (core.Credentials, error)
}

type SessionPort interface {
	Set(ctx context.Context, key string, session interface{}) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Update(ctx context.Context, key string) error
}
