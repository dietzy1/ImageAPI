package ports

import (
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

//Potentially move the ports into the database folder

//implement the mongodb interface methods
type DbPort interface {
	FindImage(ctx context.Context, querytype string, query string) (*core.Image, error)
	FindImages(ctx context.Context, querytype string, query []string, quantity int) ([]core.Image, error)
	StoreImage(ctx context.Context, image *core.Image) error
	UpdateImage(ctx context.Context, uuid string, image *core.Image) error
	DeleteImage(ctx context.Context, uuid string) error
}

//Need to implement mongodb methods on this interface
type DbAuthenticationPort interface {
	StoreKey(string, string) error
	DeleteKey(string) error
	AuthenticateKey(string) bool
	Signup(core.Credentials) error
	Signin(core.Credentials) error
}
