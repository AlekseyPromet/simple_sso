package storage

import (
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	mongo.Client
	collection *mongo.Collection
}

var _ models.IStorage = (*MongoStorage)(nil)

func (s *MongoStorage) New() *MongoStorage {
	return s
}

func (s *MongoStorage) Run(ctx context.Context, cfg models.StorageConfig) error {
	const op = "storage connection"

	client, err := mongo.Connect(ctx)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	db := client.Database(cfg.Database)

	s.collection = db.Collection(cfg.Collection)

	return nil
}

func (s *MongoStorage) Shutdown(ctx context.Context) error {
	return s.Client.Disconnect(ctx)
}

func (s *MongoStorage) Register(ctx context.Context, user models.User) (null uuid.UUID, err error) {

	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return null, err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return uuid.Parse(id.String())
	}

	return null, fmt.Errorf("insert into storage error")
}

func (s *MongoStorage) Logining(ctx context.Context, lp models.LoginPass) (err error) {

	filter := bson.M{
		"email":    lp.Email,
		"password": lp.Password,
	}

	result := s.collection.FindOne(ctx, filter)

	return result.Err()
}
