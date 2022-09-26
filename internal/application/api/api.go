package api

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	"log"
	"net/http"
	"sync"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/vitali-fedulov/images4"

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
	dbElo     ports.DbEloSystemPort
	session   ports.SessionPort

	cdn   ports.CdnPort
	image core.Image
	creds core.Credentials
}

// Constructor
func NewApplication(dbImage ports.DbImagePort, dbAccAuth ports.DbAccAuthPort, dbKeyAuth ports.DbKeyAuthPort, dbElo ports.DbEloSystemPort, session ports.SessionPort, cdn ports.CdnPort) *Application {
	return &Application{dbImage: dbImage, dbAccAuth: dbAccAuth, dbKeyAuth: dbKeyAuth, cdn: cdn, session: session}
}

// Returns a single image based on query parameters from the database.
func (a Application) FindImage(ctx context.Context, w http.ResponseWriter, r *http.Request, query string, querytype string) {
	image, err := a.dbImage.FindImage(ctx, querytype, query)
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

// Returns multiple images based on query parameters from the database.
func (a Application) FindImages(ctx context.Context, w http.ResponseWriter, r *http.Request, query []string, querytype string, quantity int) {
	images, err := a.dbImage.FindImages(ctx, querytype, query, quantity)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if images == nil {
		w.WriteHeader(http.StatusNotFound)
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

	img, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		log.Println("Error decoding image")
	}

	image := core.Image{
		Title:      r.Form.Get("title"),
		Uuid:       a.image.NewUUID(),
		Tags:       core.Split(r.Form.Get("tags")), //there is a bug here whitespace is not removed
		Created_At: a.image.SetTime(),
		Filesize:   a.image.FileSize(*buf),
		Hash:       a.image.HashSet(img),
		Width:      a.image.FindWidth(img),
		Height:     a.image.FindHeight(img),
		Elo:        1500,
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		image.BlurHash = image.BlurHashing(img)
	}()

	if r.Form.Get("skip") == "" {
		hashImages, err := a.dbImage.FindImages(ctx, "hash", nil, 0) //UUID and hash are the only two fields that are returned
		if err != nil {
			_ = json.NewEncoder(w).Encode([]any{"Unable to retrieve hashes of images from db. Here is the error value:", core.Errconv(err)})
		}

		c := make(chan bool, 2)
		wg.Add(len(hashImages))
		for _, v := range hashImages {
			go func(v core.Image) {
				defer wg.Done()
				if images4.Similar(v.Hash, image.Hash) {
					c <- true
				}
			}(v)
		}
		recieve := <-c
		if recieve {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode([]any{"Image already exists in the database. Here is the error value:", core.Errconv(err)})
			return
		}
		wg.Wait()
		close(c)
	}

	err = image.Validate(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to validate. Here is the error value:", core.Errconv(err)})
		return
	}

	url, err := a.cdn.UploadFile(ctx, image, *buf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to upload file to cdn. Here is the error value:", core.Errconv(err)})
		return
	}
	image.Filepath = url
	//Database should not fail so
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)

	err = a.dbImage.StoreImage(ctx, &image)
	if err != nil {
		a.cdn.DeleteFile(ctx, image.Uuid)
		w.WriteHeader(http.StatusBadRequest)

		_ = json.NewEncoder(w).Encode("Unable to add image while storing")
		return
	}
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
		Title:      r.Form.Get("title"),
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
