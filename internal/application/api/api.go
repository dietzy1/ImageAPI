package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/dietzy1/imageAPI/internal/ports"

	"github.com/gorilla/mux"
)

// Implements the Api port methods
type Application struct {
	db      ports.DbPort
	dbauth  ports.DbAuthenticationPort
	session ports.SessionPort
	cdn     ports.CdnPort
	image   core.Image
	creds   core.Credentials
}

// Constructor
func NewApplication(db ports.DbPort, dbauth ports.DbAuthenticationPort, cdn ports.CdnPort, session ports.SessionPort) *Application {
	return &Application{db: db, dbauth: dbauth, cdn: cdn, session: session}
}

// Implements methods on the APi port
func (a Application) FindImage(ctx context.Context, w http.ResponseWriter, r *http.Request, query string, querytype string) {
	image, err := a.db.FindImage(ctx, querytype, query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode([]any{"Unable to find the image. Here is the error value:", core.Errconv(err)})

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
func (a Application) FindImages(ctx context.Context, w http.ResponseWriter, r *http.Request, query []string, querytype string, quantity int) {
	images, err := a.db.FindImages(ctx, querytype, query, quantity)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode([]any{"Unable to find the image. Here is the error value:", core.Errconv(err)})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if images == nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode([]any{"Unable to find the image. Here is the error value:", core.Errconv(err)})
		return
	}
	json.NewEncoder(w).Encode(images)
}

// Need to implement role
func (a Application) AddImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse the multipartform. Here is the error value:", core.Errconv(err)})
		return
	}

	image := core.Image{
		Name:    r.Form.Get("name"),
		Uuid:    a.image.NewUUID(),
		Tags:    core.Split(r.Form.Get("tags")),
		Created: a.image.SetTime(),
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse file data. Here is the error value:", core.Errconv(err)})
		return
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	err = core.ConvertToJPEG(buf, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to convert file to jpg. Here is the error value:", core.Errconv(err)})
		return
	}
	err = image.Validate(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to validate. Here is the error value:", core.Errconv(err)})
		return
	}

	url, err := a.cdn.UploadFile(ctx, image, buf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to upload file to cdn. Here is the error value:", core.Errconv(err)})
		return
	}
	image.Filepath = url

	err = a.db.StoreImage(ctx, &image)
	if err != nil {
		a.cdn.DeleteFile(ctx, image.Uuid)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while storing")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)
}

// Implements methods on the APi port
func (a Application) DeleteImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.cdn.DeleteFile(ctx, vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to delete the image in the cdn. Here is the error value:", core.Errconv(err)})
		return
	}

	err = a.db.DeleteImage(ctx, vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to delete the image in the mongodb. Here is the error value:", core.Errconv(err)})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Implements methods on the APi port
func (a Application) UpdateImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse the multipartform. Here is the error value:", core.Errconv(err)})
		return
	}

	image := core.Image{
		Name:    r.Form.Get("name"),
		Uuid:    r.Form.Get("uuid"),
		Tags:    core.Split(r.Form.Get("tags")),
		Created: a.image.SetTime(),
	}
	err = image.Validate(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to validate. Here is the error value:", core.Errconv(err)})
		return
	}

	err = a.db.UpdateImage(ctx, &image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to update the image in mongodb. Here is the error value:", core.Errconv(err)})
		return
	}
	err = a.cdn.UpdateFile(ctx, image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse the image in cdn. Here is the error value:", core.Errconv(err)})
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(image)
}
