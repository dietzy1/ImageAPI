package repository

import (
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"go.mongodb.org/mongo-driver/bson"
)

//Needs to find the field in the collection and update it
func (a *DbAdapter) StoreKey(ctx context.Context, newKey string, username string) error {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "key", Value: newKey}}}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) AuthenticateKey(ctx context.Context, key string) bool {
	return true
}

func (a *DbAdapter) DeleteKey(ctx context.Context, username string) error {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "key", Value: ""}}}} //Set key to empty string
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) Signup(ctx context.Context, creds core.Credentials) error {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	_, err := collection.InsertOne(ctx, creds)
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) Signin(ctx context.Context, username string) (core.Credentials, error) {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	var cred core.Credentials
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&cred)
	if err != nil {
		return cred, err
	}
	return cred, nil
}
