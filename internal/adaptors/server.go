package adapter

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"github.com/gorilla/mux"
)

type ServerAdapter struct {
	api    ports.ApiPort
	router http.Handler
}

//helper function
func analyzeInterface(mp ports.ApiPort) {
	fmt.Printf("Interface type: %T\n", mp)
	fmt.Printf("Interface value: %v\n", mp)
	fmt.Printf("Interface is nil: %t\n", mp == nil)
}

//Constructor
func NewServerAdapter(api ports.ApiPort) *ServerAdapter {
	return &ServerAdapter{api: api}
}

//Wrapper for router object
func (s *ServerAdapter) Router() http.Handler {
	return s.router
}

//Should prolly add subroutes
//need to add middleware
//Routering and paths
func Router(s *ServerAdapter) {

	r := mux.NewRouter()
	r.HandleFunc("/images/", s.findImages).Queries("uuid", "{uuid}").Methods(http.MethodGet)
	r.HandleFunc("/image/", s.findImage).Queries("uuid", "{uuid}").Methods(http.MethodGet)
	r.HandleFunc("/image/", s.addImage).Methods(http.MethodPost) //form data
	r.HandleFunc("/image/{uuid}", s.updateImage).Queries("uuid", "{uuid}").Methods(http.MethodPut)
	r.HandleFunc("/image/{uuid}", s.deleteImage).Queries("uuid", "{uuid}").Methods(http.MethodDelete)
	r.HandleFunc("/healthcheck", s.healthcheck)

	//New query routes
	//r.HandleFunc("/tag{tag}", s.findtag.Methods(http.MethodGet))
	//r.HandleFunc("/tags{tags}", s.findtags.Methods(http.MethodGet))

	//Query based on tags provides 1 random image from random order
	//Query based on tags provides slice of images in random order

	srv := &http.Server{ //&http.Server
		Handler:      r,
		Addr:         "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	s.router = r
}

//Entry point for http calls
func (s *ServerAdapter) findImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("findimage")
	s.api.FindImage(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) findImages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("findimages")
	s.api.FindImages(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) addImage(w http.ResponseWriter, r *http.Request) {
	s.api.AddImage(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) updateImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateimage")
	s.api.UpdateImage(w, r)
}

//Entry point for http calls
func (s *ServerAdapter) deleteImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteimage")
	s.api.DeleteImage(w, r)
}

func (s *ServerAdapter) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
