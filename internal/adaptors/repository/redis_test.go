package repository

import (
	"context"
	"log"
	"testing"

	"github.com/dietzy1/imageAPI/internal/ports"
)

//Mock test object
type mock struct {
	redisClient ports.SessionPort
	key         string
	value       string
}

//Mock test constructor
func newMock() *mock {
	redis, err := NewRedisAdapter()
	if err != nil {
		log.Fatal(err)
	}
	return &mock{redisClient: redis, key: "testkey", value: "testvalue"}
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestRedis(t *testing.T) {
	mock := newMock()
	ctx := context.Background()
	mock.redisClient.Get(ctx, mock.key)

	//Add test key into redis
	if err := mock.redisClient.Set(ctx, mock.key, mock.value); err != nil {
		t.Error(err)
	}
	//Retrieve test key --Correct one
	val, err := mock.redisClient.Get(ctx, mock.key)
	if err != nil {
		t.Error(err)
	}
	if val != mock.value {
		t.Errorf("Expected %s, got %s", mock.value, val)
	}

	//Retrieve test key --incorrect One
	val, err = mock.redisClient.Get(ctx, "wrongkey")

	if err == nil {
		t.Errorf("Expected error, got %s", val)
	}
	if val != "" {
		t.Errorf("Expected empty string, got %s", val)
	}

	//Update test key --Correct one
	if err := mock.redisClient.Update(ctx, mock.key); err != nil {
		t.Error(err)
	}

	//Update test key --incorrect One
	if err := mock.redisClient.Update(ctx, "wrongkey"); err != nil {
		t.Errorf("Expected error, got %s", err)
	}

	//delete the key
	if err := mock.redisClient.Delete(ctx, mock.key); err != nil {
		t.Error(err)
	}
	//Check again if the key exists
	val, err = mock.redisClient.Get(ctx, mock.key)
	if err == nil {
		t.Errorf("Expected error, got %s", val)
	}
	if val != "" {
		t.Errorf("Expected empty string, got %s", val)
	}

}
