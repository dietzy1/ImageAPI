package repository

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dietzy1/imageAPI/internal/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongomock struct {
	mongoClient ports.DbImagePort
}

func newMongoMock() *mongomock {
	mongo, err := NewMongoTestAdapter()
	if err != nil {
		log.Fatal(err)
	}
	return &mongomock{mongoClient: mongo}
}

func NewMongoTestAdapter() (*DbAdapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	uri := ""
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(uri)))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	a := &DbAdapter{client: client}
	return a, nil
}

func AddTestData() {
	/* 	mock := newMongoMock()
	   	image := &core.Image{
	   		Uuid:    "test",
	   		Tags:    []string{"test"},
	   		Created: time.Now(),
	   		Updated: time.Now(),
	   	}
	   	err := mock.mongoClient.StoreImage(image)
	   	if err != nil {
	   		log.Fatal(err)
	   	} */
}

func DeleteTestData() {
	mock := newMongoMock()
	ctx := context.Background()
	err := mock.mongoClient.DeleteImage(ctx, "test")
	if err != nil {
		log.Fatal(err)

	}
}

func TestFindImage(t *testing.T) {
	mock := newMongoMock()
	ctx := context.Background()
	image, err := mock.mongoClient.FindImage(ctx, "", "test")
	if err != nil {
		t.Error(err)
	}
	if image.Uuid != "test" {
		t.Error("Wrong image")
	}

}

func TestFindImages(t *testing.T) {

}

func TestStoreImages(t *testing.T) {

}

func TestUpdateImages(t *testing.T) {

}

func TestDeleteImages(t *testing.T) {

}
