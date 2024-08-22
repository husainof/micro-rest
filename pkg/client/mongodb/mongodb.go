package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	var mongodbURI string
	var isAuth bool

	if username == "" && password == "" {
		mongodbURI = fmt.Sprintf("mongodb:\\%s:%s", host, port)
	} else {
		mongodbURI = fmt.Sprintf("mongodb:\\%s:%s@%s:%s", host, port, username, password)
	}

	clientOptions := options.Client().ApplyURI(mongodbURI)

	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			Username:   username,
			Password:   password,
			AuthSource: authDB,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongoDB with error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB with error: %v", err)
	}

	return client.Database(database), nil
}
