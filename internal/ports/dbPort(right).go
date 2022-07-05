package ports

import "github.com/dietzy1/imageAPI/internal/application/core"

//Potentially move the ports into the database folder

//implement the mongodb interface methods
type DbPort interface {
	FindImage(uuid string, querytype string) (*core.Image, error)
	FindImages(uuid []string, querytype string) ([]*core.Image, error)
	Store(image *core.Image) error
	Update(uuuid string, image *core.Image) error
	Delete(uuid string) error
}
