package ports

import "github.com/dietzy1/imageAPI/internal/application/core"

//Potentially move the ports into the database folder

//implement the mongodb interface methods
type DbPort interface {
	FindImage(querytype string, query string) (*core.Image, error)
	FindImages(querytype string, query []string, quantity int) ([]core.Image, error)
	Store(image *core.Image) error
	Update(uuid string, image *core.Image) error
	Delete(uuid string) error
}
