package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/dietzy1/imageAPI/internal/ports"

	"github.com/gorilla/mux"
)

//Implements the Api port methods
type Application struct {
	db    ports.DbPort
	image core.Image
	file  ports.FilePort
}

//Constructor
func NewApplication(db ports.DbPort, file ports.FilePort) *Application {
	return &Application{db: db, file: file}
}

//Implements methods on the APi port
func (a Application) FindImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//mongoDB function to return the json
	image, err := a.db.FindImage(vars["uuid"], "uuid")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//fmt.Println(image)
	w.Header().Set("Content-Type", "application/json")
	if image == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(image)

	//fileDB
	file, err := a.file.FindFile(image.Uuid)
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	//need to find out what format is good for sending the file over
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(file)
}

//Implements methods on the APi port
func (a Application) FindImages(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	uuid := []string{}
	for _, v := range q["uuid"] {
		uuid = append(uuid, v)
	}
	images, err := a.db.FindImages(uuid, "uuid")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if images == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(images)

	//Need to do filedatabase call

}

func (a Application) AddImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inner API application recieved signal")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while parsing")
		return
	}
	a.image.Name = r.Form.Get("name")
	a.image.Tags = r.Form["tags"]
	//a.image.Data = r.Form.Get("data")
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
	err = a.db.Store(&a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to add image while storing")
		return
	}
	//Potentially add in a call to the encoding of the image so application logic remains seperate
	err = a.file.AddFile(a.image.Uuid, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//Potentially need to add in missing a delete function for the database if an error occurs
		a.db.Delete(a.image.Uuid)
		_ = json.NewEncoder(w).Encode("some other error")
		return
	}
	w.WriteHeader(http.StatusCreated)

}

//Implements methods on the APi port
func (a Application) DeleteImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inner API application recieved signal")
	vars := mux.Vars(r)
	err := a.db.Delete(vars["uuid"])
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
	fmt.Println("inner API application recieved signal")
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = a.db.Update(vars["uuid"], &a.image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to update image")
		return
	}
	w.WriteHeader(http.StatusOK)
}
