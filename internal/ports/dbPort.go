package ports

import "github.com/dietzy1/imageAPI/internal/application/core"

//Potentially move the ports into the database folder

//implement the mongodb interface methods
type DbPort interface {
	FindImage(querytype string, query string) (*core.Image, error)
	FindImages(querytype string, query []string, quantity int) ([]core.Image, error)
	StoreImage(image *core.Image) error
	UpdateImage(uuid string, image *core.Image) error
	DeleteImage(uuid string) error
}

//Need to implement mongodb methods on this interface
type DbKeyPort interface {
	StoreKey(string) error
	DeleteKey(string) error
	AuthenticateKey(string) bool
}
