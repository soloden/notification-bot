package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, username string, password string, host string, port string, database string) (*mongo.Database, error) {
	mongoDBURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	credentials := options.Credential{
		Username:    username,
		Password:    password,
		PasswordSet: true,
	}

	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(reqCtx, options.Client().ApplyURI(mongoDBURL).SetAuth(credentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	return client.Database(database), nil
}
