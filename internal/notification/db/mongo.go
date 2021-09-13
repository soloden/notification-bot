package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/soloden/notificator-bot/internal/apperror"
	"github.com/soloden/notificator-bot/internal/notification"
	"github.com/soloden/notificator-bot/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *db) GetOne(ctx context.Context, id string) (ntf notification.Notification, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ntf, fmt.Errorf("failed to convert hex to objectId, err: %w", err)
	}

	filter := bson.M{"_id": objectId}

	result := s.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		s.logger.Error(result.Err())
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return ntf, apperror.NewAppError("404", "notification not found")
		}

		return ntf, fmt.Errorf("failed to query execute, err: %w", err)
	}

	if result.Decode(&ntf); err != nil {
		return ntf, fmt.Errorf("failed to decode document, error: %w", err)
	}

	return ntf, nil
}

func (s *db) GetNotificationByUserId(ctx context.Context, userId string) (ntfs []notification.Notification, err error) {
	filter := bson.M{"owner_id": bson.M{"$eq": userId}}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ntfs, apperror.NewAppError("404", "notifications not found")
		}

		return ntfs, fmt.Errorf("failed to query execute, err: %w", err)
	}

	if cursor.All(ctx, &ntfs); err != nil {
		return ntfs, fmt.Errorf("failed to decode document, error: %w", err)
	}

	return ntfs, nil
}

func (s *db) Update(ctx context.Context, ntf notification.Notification) error {
	filter := bson.M{"_id": ntf.Id}
	userByte, err := bson.Marshal(ntf)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	err = bson.Unmarshal(userByte, &updateObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")

	update := bson.M{
		"$set": updateObj,
	}

	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if result.MatchedCount == 0 {
		return apperror.NewAppError("404", "notification not found")
	}

	s.logger.Tracef("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (s *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convet objectid to hex. error: %w", err)
	}
	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	result, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query")
	}
	if result.DeletedCount == 0 {
		return apperror.NewAppError("404", "notification not found")
	}

	s.logger.Tracef("Delete %v documents.\n", result.DeletedCount)

	return nil
}
