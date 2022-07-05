package adapter

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

//Implements the db port interface
type DbAdapter struct {
	client  *mongo.Client
	timeout time.Duration
}

//Constructor
func NewDbAdapter() (*DbAdapter, error) {
	fmt.Println("Initiating DbAdabter")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(("DB_URI"))))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &DbAdapter{client: client, timeout: 5 * time.Second}, nil
	//Maybe I just need to only return the db as a variable
}

func (a *DbAdapter) FindImage(uuid string, querytype string) (*core.Image, error) {
	fmt.Println("mongodb called")
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	collection := a.client.Database("Image-Database").Collection("images")
	image := &core.Image{}
	err := collection.FindOne(ctx, bson.M{querytype: uuid}).Decode(&image)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return image, nil
}

func (a *DbAdapter) FindImages(uuid []string, querytype string) ([]*core.Image, error) {
	fmt.Println("mongodb called")
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	collection := a.client.Database("Image-Database").Collection("images")
	images := []*core.Image{}
	for _, v := range uuid {
		image := core.Image{}
		err := collection.FindOne(ctx, bson.M{querytype: v}).Decode(&image)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		images = append(images, &image)
	}
	return images, nil
}

func (a *DbAdapter) Store(image *core.Image) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	collection := a.client.Database("Image-Database").Collection("images")
	_, err := collection.InsertOne(ctx, image)
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) Update(uuid string, image *core.Image) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	collection := a.client.Database("Image-Database").Collection("images")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$set": image})
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) Delete(uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	collection := a.client.Database("Image-Database").Collection("images")
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}
