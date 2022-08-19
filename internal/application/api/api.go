package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/dietzy1/imageAPI/internal/ports"

	"github.com/gorilla/mux"
)

// Implements the Api port methods
type Application struct {
	db      ports.DbPort
	dbauth  ports.DbAuthenticationPort
	session ports.SessionPort
	file    ports.FilePort
	image   core.Image
	creds   core.Credentials
}

// Constructor
func NewApplication(db ports.DbPort, dbauth ports.DbAuthenticationPort, file ports.FilePort, session ports.SessionPort) *Application {
	return &Application{db: db, dbauth: dbauth, file: file, session: session}
}

// Implements methods on the APi port
func (a Application) FindImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	querytype := "uuid"
	query := vars["uuid"]
	if vars["uuid"] == "" {
		querytype = "tags"
		query = vars["tag"]
	}

	image, err := a.db.FindImage(ctx, querytype, query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if image == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(image)
}

// Implements methods on the APi port
func (a Application) FindImages(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	querytype := "tags"
	quantity, err := strconv.Atoi(strings.Join(q["quantity"], ""))
	if err != nil || quantity <= 0 { //<= 0 is a hack to allow for a default value
		quantity = 1
	}
	query := []string{}
	tags := strings.Join(q["tags"], "")
	query = strings.Split(tags, ", ")

	images, err := a.db.FindImages(ctx, querytype, query, quantity)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
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

// Need to implement role
func (a Application) AddImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}

	image := core.Image{
		Name:    r.Form.Get("Name"),
		Uuid:    a.image.NewUUID(),
		Tags:    core.Split(r.Form.Get("Tags")),
		Created: a.image.SetTime(),
	}
	data, _, err := r.FormFile("data")
	if err != nil {
		_ = json.NewEncoder(w).Encode("Unable to parse file data")
		return
	}
	defer data.Close()

	buf := new(bytes.Buffer)
	err = core.ConvertToJPEG(buf, data)
	if err != nil {
		_ = json.NewEncoder(w).Encode("Unable to convert file to jpg")
		fmt.Println(err)
		return
	}
	image.Validate(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Missing name or tag")
		return
	}

	url, err := a.file.UploadFile(ctx, image, buf)
	if err != nil {
		_ = json.NewEncoder(w).Encode("Unable to upload file")
		fmt.Println(err)
		return
	}
	image.Filepath = url

	err = a.db.StoreImage(ctx, &image)
	if err != nil {
		//Call to delete file
		a.file.DeleteFile(ctx, image)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while storing")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// need to implement role
// Implements methods on the APi port
func (a Application) DeleteImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.db.DeleteImage(ctx, vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to delete image")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Need to implement role
// Implements methods on the APi port
func (a Application) UpdateImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = a.db.UpdateImage(ctx, vars["uuid"], &a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return
	}
	err = a.file.UpdateFile(ctx, a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return

	}
	w.WriteHeader(http.StatusOK)
}
