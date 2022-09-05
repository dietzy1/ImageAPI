package repository

import (
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"go.mongodb.org/mongo-driver/bson"
)

//Mongodb implementation
//Implements methods on the port: type DbKeyAuthPort interface
//This file is responcible for the base layer of database operations related to API-keys.

// Updates a API key value in the credentials database.
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

// Checks the credentials database if the submitted key is valid.
func (a *DbAdapter) AuthenticateKey(ctx context.Context, key string) (string, bool) {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	//Check if the key is in the database if not return false
	cred := core.Credentials{}
	err := collection.FindOne(ctx, bson.M{"key": key}).Decode(&cred)
	if err != nil {
		return "", false
	}
	return cred.Username, true
}

// Sets the key field to empty string in the credentials database.
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

// Retrrieves a key from the credentials database based on username input.
func (a *DbAdapter) GetKey(ctx context.Context, username string) (string, error) {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	cred := core.Credentials{}
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&cred)
	if err != nil {
		return "", err
	}
	return cred.Key, nil
}
