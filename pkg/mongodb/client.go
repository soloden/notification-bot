package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, username string, password string, host string, port string, database string) (*mongo.Database, error) {
	var mongoDBURL string
	var anonymous bool

	if username != "" || password != "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	} else {
		anonymous = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	if !anonymous {
		clientOptions.SetAuth(options.Credential{
			Username:    username,
			Password:    password,
			PasswordSet: true,
		})
	}

	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(reqCtx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	return client.Database(database), nil
}
