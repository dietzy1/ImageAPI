package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/dietzy1/imageAPI/internal/ports"

	"github.com/gorilla/mux"
)

//Implements the Api port methods
type Application struct {
	db     ports.DbPort
	image  core.Image
	creds  core.Credentials
	file   ports.FilePort
	dbauth ports.DbAuthenticationPort
}

//Constructor
func NewApplication(db ports.DbPort, file ports.FilePort) *Application {
	return &Application{db: db, file: file}
}

//Implements methods on the APi port
func (a Application) FindImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	querytype := "uuid"
	query := vars["uuid"]
	if vars["uuid"] == "" {
		querytype = "tags"
		query = vars["tag"]
	}

	image, err := a.db.FindImage(querytype, query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	image.Filepath = "http://localhost:8000/fileserver/" + image.Uuid + ".jpg"
	w.Header().Set("Content-Type", "application/json")
	if image == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(image)
}

//Implements methods on the APi port
func (a Application) FindImages(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	querytype := "tags"
	quantity, err := strconv.Atoi(strings.Join(q["quantity"], ""))
	if err != nil || quantity <= 0 { //<= 0 is a hack to allow for a default value
		quantity = 1
	}
	query := []string{}
	tags := strings.Join(q["tags"], "")
	query = strings.Split(tags, ", ")

	images, err := a.db.FindImages(querytype, query, quantity)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for i := 0; i < len(images); i++ {
		images[i].Filepath = "http://localhost:8000/fileserver/" + images[i].Uuid + ".jpg"
		w.Header().Set("Content-Type", "application/json")
	}
	if images == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(images)
}

func (a Application) AddImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
	a.image.Name = r.Form.Get("name")
	tags := strings.Join(r.Form["tags"], "")
	a.image.Tags = strings.Split(tags, ", ")
	a.image.NewUUID()
	a.image.SetTime()
	data, _, err := r.FormFile("data")
	defer data.Close()
	if err != nil {
		_ = json.NewEncoder(w).Encode("Unable to parse file data")
		return
	}
	a.image.Validate(a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Missing name or tag")
		return
	}
	err = a.db.StoreImage(&a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while storing")
		return
	}
	err = a.file.AddFile(a.image.Uuid, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		a.db.DeleteImage(a.image.Uuid)
		_ = json.NewEncoder(w).Encode("some other error")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//Implements methods on the APi port
func (a Application) DeleteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.db.DeleteImage(vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete image")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//Implements methods on the APi port
func (a Application) UpdateImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = a.db.UpdateImage(vars["uuid"], &a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return
	}
	w.WriteHeader(http.StatusOK)
}
