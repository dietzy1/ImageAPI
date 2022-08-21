package repository

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/go-redis/redis/v8"
)

// Implements the db port interface and dbApiKey interface
type DbAdapter struct {
	client      *mongo.Client
	redisClient *redis.Client
}

// Constructor
func NewMongoAdapter() (*DbAdapter, error) {
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
	a := &DbAdapter{client: client}
	//Hard coded index
	/* 	a.NewIndex("Image-Database", "images", "tags", false) //Collection name, field, unique
	   	a.NewIndex("Image-Database", "images", "uuid", false)
	   	a.NewIndex("Credential-Database", "credentials", "key", false)
	   	a.NewIndex("Credential-Database", "credentials", "username", false) */
	return a, nil
}

// Mongodb index - b tree
func (a *DbAdapter) NewIndex(database string, collectionName string, field string, unique bool) {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := a.client.Database(database).Collection(collectionName)

	index, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Created new index:", index)

}

func (a *DbAdapter) FindImage(ctx context.Context, querytype string, query string) (*core.Image, error) {
	collection := a.client.Database("Image-Database").Collection("images")
	// can accept uuid or random
	image := &core.Image{}
	switch querytype {
	case "uuid":
		cursor, err := collection.Find(ctx, bson.D{{Key: querytype, Value: query}})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &image); err != nil {
			return nil, err
		}

	case "random":
		cursor, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": 1}}})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &image); err != nil {
			return nil, err
		}
	}
	return image, nil
}

func (a *DbAdapter) FindImages(ctx context.Context, querytype string, query []string, quantity int) ([]core.Image, error) {
	collection := a.client.Database("Image-Database").Collection("images")
	//Can accept tags or random
	images := []core.Image{}

	switch querytype {
	case "tags":
		var otps bson.D
		if len(query) == 1 {
			otps = bson.D{{Key: querytype, Value: query[0]}}
		}
		if len(query) == 2 {
			otps = bson.D{{Key: querytype, Value: query[0]}, {Key: querytype, Value: query[1]}}
		}
		if len(query) >= 3 {
			otps = bson.D{{Key: querytype, Value: query[0]}, {Key: querytype, Value: query[1]}, {Key: querytype, Value: query[2]}}
		}
		cursor, err := collection.Find(ctx, otps)
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &images); err != nil {
			return nil, err
		}
		images = randomizeArray(images, quantity)
	case "random":
		cursor, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": quantity}}})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &images); err != nil {
			return nil, err
		}
	}

	return images, nil
}

func (a *DbAdapter) StoreImage(ctx context.Context, image *core.Image) error {
	collection := a.client.Database("Image-Database").Collection("images")
	_, err := collection.InsertOne(ctx, image)
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) UpdateImage(ctx context.Context, uuid string, image *core.Image) error {
	collection := a.client.Database("Image-Database").Collection("images")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$set": image})
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) DeleteImage(ctx context.Context, uuid string) error {
	collection := a.client.Database("Image-Database").Collection("images")
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}

/* func randomize(images []core.Image) *core.Image {
	if len(images) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(images))
	image := images[randomIndex]
	return &image
} */

// Must randomize the images based on the quantity
func randomizeArray(images []core.Image, quantity int) []core.Image {
	if len(images) == 0 {
		return nil
	}
	if len(images) < quantity {
		quantity = len(images)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndexes := rand.Perm(len(images))
	randomImages := []core.Image{}
	for i := 0; i < quantity; i++ {
		randomImages = append(randomImages, images[randomIndexes[i]])
	}
	return randomImages
}
