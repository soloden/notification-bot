package db

import (
	"context"
	"fmt"

	"github.com/soloden/notificator-bot/internal/notification"
	"github.com/soloden/notificator-bot/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(storage *mongo.Database, collection string, logger *logging.Logger) notification.Storage {
	return &db{
		collection: storage.Collection(collection),
		logger:     logger,
	}
}

func (s *db) Create(ctx context.Context, ntf notification.Notification) (string, error) {
	result, err := s.collection.InsertOne(ctx, ntf)
	if err != nil {
		return "", fmt.Errorf("failed to create notification due err - %s", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed ti convert objectid to hex")
}

func (s *db) GetOne(ctx context.Context, id string) (notification.Notification, error) {
	panic("not implemented") // TODO: Implement
}

func (s *db) GetMany(ctx context.Context, ids []string) ([]notification.Notification, error) {
	panic("not implemented") // TODO: Implement
}

func (s *db) Update(ctx context.Context, ntf notification.Notification) error {
	panic("not implemented") // TODO: Implement
}

func (s *db) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
