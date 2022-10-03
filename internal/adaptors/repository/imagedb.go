package repository

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/go-redis/redis/v8"
)

//Mongodb implementation
//Implements methods on the port: type DbImagePort interface
//This file is responcible for all CRUD operations for storing image object json data.

// This is the main struct that contains all database related methods
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
	   	a.NewIndex("Image-Database", "images", "hash", false)
	   	a.NewIndex("Image-Database", "images", "elo", false)

	   	a.NewIndex("Credential-Database", "Credentials", "key", false)
	   	a.NewIndex("Credential-Database", "Credentials", "username", false) */
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

// Performs either a random or specific uuid query against the database for an image object. Returns a single object.
func (a *DbAdapter) FindImage(ctx context.Context, querytype string, query string) (*core.Image, error) {
	collection := a.client.Database("Image-Database").Collection("Images")
	// can accept uuid or random
	image := &core.Image{}
	images := []core.Image{}
	switch querytype {
	case "uuid":
		err := collection.FindOne(ctx, bson.D{{Key: querytype, Value: query}}).Decode(&image)
		if err != nil {
			return nil, err
		}
		return image, nil
	case "random":
		cursor, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": 1}}})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &images); err != nil {
			return nil, err
		}
		return &images[0], nil
	}
	return nil, nil
}

// Performs either a random query or a query by tags against the database. If multiple tags are provided then only images that adhere to all tags are returned. Returns an array of multiple objects based on quantity provided.
func (a *DbAdapter) FindImages(ctx context.Context, querytype string, query []string, quantity int) ([]core.Image, error) {
	collection := a.client.Database("Image-Database").Collection("Images")
	//Can accept tags, hash or random
	images := []core.Image{}

	switch querytype {
	case "tags":
		projection := bson.D{
			{Key: "_id", Value: 0},
			{Key: "hash", Value: 0},
		}
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
		cursor, err := collection.Find(ctx, otps, options.Find().SetProjection(projection))
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &images); err != nil {
			return nil, err
		}
		images = randomizeArray(images, quantity)
	case "random":
		projection := bson.D{
			{Key: "_id", Value: 0},
			{Key: "hash", Value: 0},
		}
		cursor, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": quantity}}, bson.M{"$project": projection}})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &images); err != nil {
			return nil, err
		}
	case "hash":
		//Only returns uuid and hash value to the slice
		projection := bson.D{
			{Key: "hash", Value: 1},
			{Key: "uuid", Value: 1},
			{Key: "_id", Value: 0},
		}
		cursor, err := collection.Find(ctx, bson.D{{}}, options.Find().SetProjection(projection))
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &images); err != nil {
			return nil, err
		}
	}
	return images, nil
}

// Stores an image object in the database -- does not contain the image itself but the filepath its hosted CDN position.
func (a *DbAdapter) StoreImage(ctx context.Context, image *core.Image) error {
	collection := a.client.Database("Image-Database").Collection("Images")
	_, err := collection.InsertOne(ctx, image)
	if err != nil {
		return err
	}
	return nil
}

// Updates an image object in the database
func (a *DbAdapter) UpdateImage(ctx context.Context, image *core.Image) error {
	collection := a.client.Database("Image-Database").Collection("Images")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": image.Uuid}, bson.M{"$set": image})
	if err != nil {
		return err
	}
	return nil
}

// Deletes an image object in the database
func (a *DbAdapter) DeleteImage(ctx context.Context, uuid string) error {
	collection := a.client.Database("Image-Database").Collection("Images")
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}

// randomizes the images based on the quantity
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

func (a *DbAdapter) GetLeaderBoardImages(ctx context.Context) ([]core.Image, error) {
	collection := a.client.Database("Image-Database").Collection("Images")

	images := []core.Image{}
	projection := bson.D{
		{Key: "title", Value: 1},
		{Key: "uuid", Value: 1},
		{Key: "tags", Value: 1},
		{Key: "filepath", Value: 1},
		{Key: "elo", Value: 1},
		{Key: "blurhash", Value: 1},
		{Key: "_id", Value: 0},
	}
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "elo", Value: -1}}).SetLimit(100).SetProjection(projection)

	cursor, err := collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &images); err != nil {
		return nil, err
	}
	return images, nil
}

// Choose a random image from the database and then aggregate for another image within a certain threshhold
func (a *DbAdapter) FindMatch(ctx context.Context) ([]core.Image, error) {
	collection := a.client.Database("Image-Database").Collection("Images")
	images := []core.Image{}

	projection := bson.D{
		{Key: "title", Value: 1},
		{Key: "uuid", Value: 1},
		{Key: "tags", Value: 1},
		{Key: "filepath", Value: 1},
		{Key: "elo", Value: 1},
		{Key: "blurhash", Value: 1},
		{Key: "_id", Value: 0},
	}
	cursor, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": 1}}, bson.M{"$project": projection}})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &images); err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "$and",
		Value: bson.A{
			bson.D{{Key: "elo", Value: bson.D{{Key: "$gte", Value: images[0].Elo - 300}}}},
			bson.D{{Key: "elo", Value: bson.D{{Key: "$lte", Value: images[0].Elo + 300}}}},
		},
	},
	}
	temp := []core.Image{}
	opts := options.Find().SetProjection(projection)
	cursor, err = collection.Find(ctx, filter, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = cursor.All(ctx, &temp); err != nil {
		return nil, err
	}

	images = append(images, randomizeArray(temp, 1)[0])

	if images[1].Uuid == images[0].Uuid {
		a.FindMatch(ctx)
	}
	return images, nil
}
