package repository

import (
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"go.mongodb.org/mongo-driver/bson"
)

//Mongodb implementation
//Implements methods on the port: type DbAccAuthPort interface
//this file is responcible for database functions that are related to user authentication. 

// adds a users credentials to the database.
func (a *DbAdapter) Signup(ctx context.Context, creds core.Credentials) error {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	_, err := collection.InsertOne(ctx, creds)
	if err != nil {
		return err
	}
	return nil
}

// Fetches user information for password comparison in other functions.
func (a *DbAdapter) Signin(ctx context.Context, username string) (core.Credentials, error) {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	var cred core.Credentials
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&cred)
	if err != nil {
		return cred, err
	}
	return cred, nil
}

// Permanently deletes a user account
func (a *DbAdapter) DeleteAccount(ctx context.Context, username string) error {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	result := collection.FindOneAndDelete(ctx, username)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
