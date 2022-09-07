package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/dietzy1/imageAPI/internal/ports"

	"github.com/gorilla/mux"
)

//This file contains the main Application struct
//This file is responcible for delegating the main CRUD API operations to the database layer and returning http responses.
//Implements methods on the type ApiPort interface.
//All of these methods are called from the handlers layer.

type Application struct {
	dbImage   ports.DbImagePort
	dbAccAuth ports.DbAccAuthPort
	dbKeyAuth ports.DbKeyAuthPort
	session   ports.SessionPort

	cdn   ports.CdnPort
	image core.Image
	creds core.Credentials
}

// Constructor
func NewApplication(dbImage ports.DbImagePort, dbAccAuth ports.DbAccAuthPort, dbKeyAuth ports.DbKeyAuthPort, session ports.SessionPort, cdn ports.CdnPort) *Application {
	return &Application{dbImage: dbImage, dbAccAuth: dbAccAuth, dbKeyAuth: dbKeyAuth, cdn: cdn, session: session}
}

// Returns a single image based on query parameters from the database.
func (a Application) FindImage(ctx context.Context, w http.ResponseWriter, r *http.Request, query string, querytype string) {
	image, err := a.dbImage.FindImage(ctx, querytype, query)
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

// Returns multiple images based on query parameters from the database.
func (a Application) FindImages(ctx context.Context, w http.ResponseWriter, r *http.Request, query []string, querytype string, quantity int) {
	images, err := a.dbImage.FindImages(ctx, querytype, query, quantity)
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

// Adds a single image CDN and the image meta data to database. Simple verification and image convertion is done aswell.
func (a Application) AddImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse the multipartform. Here is the error value:", core.Errconv(err)})
		return
	}

	file, _, err := r.FormFile("file")
	if file == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("1")
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse the file. Here is the error value:", core.Errconv(err)})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("2")
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse file data. Here is the error value:", core.Errconv(err)})
		return
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	err = core.ConvertToJPEG(buf, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("3")
		_ = json.NewEncoder(w).Encode([]any{"Unable to convert file to jpg. Here is the error value:", core.Errconv(err)})
		return
	}

	image := core.Image{
		Name:       r.Form.Get("name"),
		Uuid:       a.image.NewUUID(),
		Tags:       core.Split(r.Form.Get("tags")), //there is a bug here whitespace is not removed
		Created_At: a.image.SetTime(),
		Filesize:   a.image.FileSize(*buf),
		Hash:       a.image.HashSet(*buf),
	}

	err = image.Validate(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("4")
		_ = json.NewEncoder(w).Encode([]any{"Unable to validate. Here is the error value:", core.Errconv(err)})
		return
	}

	//Logic for checking if the image already exists in the database.
	if r.Form.Get("skip") == "" {
		images, err := a.dbImage.FindImages(ctx, "hash", nil, 0) //UUID and hash are the only two fields that are returned
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("5")
			_ = json.NewEncoder(w).Encode([]any{"Unable to retrieve hashes of images from db. Here is the error value:", core.Errconv(err)})
			return
		}
		centralhash, err := a.image.CentralHash(*buf)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("6")
			_ = json.NewEncoder(w).Encode([]any{"Unable to retrieve central hash of image. Here is the error value:", core.Errconv(err)})
			return
		}

		//Implement the batch processing here
		//Temporary comparison solution

		for _, v := range images {
			if centralhash == v.Hash {
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("7")
				_ = json.NewEncoder(w).Encode([]any{"Image potentially already exists in the fdatabase, set the query parameter skip to true to skip this logic check. The uuid of the image is:", v.Uuid})
				return
			}
		}
		//end of skip control group
	}

	url, err := a.cdn.UploadFile(ctx, image, *buf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("8")
		_ = json.NewEncoder(w).Encode([]any{"Unable to upload file to cdn. Here is the error value:", core.Errconv(err)})
		return
	}
	image.Filepath = url

	err = a.dbImage.StoreImage(ctx, &image)
	if err != nil {
		a.cdn.DeleteFile(ctx, image.Uuid)
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("9")
		_ = json.NewEncoder(w).Encode("Unable to add image while storing")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)
}

// Deletes an image from the CDN and database.
func (a Application) DeleteImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.cdn.DeleteFile(ctx, vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to delete the image in the cdn. Here is the error value:", core.Errconv(err)})
		return
	}

	err = a.dbImage.DeleteImage(ctx, vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to delete the image in the mongodb. Here is the error value:", core.Errconv(err)})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Updates image metadata in the CDN and database. Simple verification is done aswell.
func (a Application) UpdateImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse the multipartform. Here is the error value:", core.Errconv(err)})
		return
	}

	image := core.Image{
		Name:       r.Form.Get("name"),
		Uuid:       r.Form.Get("uuid"),
		Tags:       core.Split(r.Form.Get("tags")),
		Created_At: a.image.SetTime(),
	}
	err = image.Validate(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to validate. Here is the error value:", core.Errconv(err)})
		return
	}

	err = a.dbImage.UpdateImage(ctx, &image)
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
