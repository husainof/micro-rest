package mongo

import (
	"context"
	"micro/internal/user"
	"micro/pkg/logging"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	logger     *logging.Logger
	collection *mongo.Collection
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *db) Create(ctx context.Context, u *user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, u)
	if err != nil {
		return "", errors.WithMessage(err, "failed insert user to the mongodb")
	}
	d.logger.Debug("convert users id")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(u)
	return "", errors.WithMessage(err, "failed convert object id to  hex")
}

func (d *db) FindOne(ctx context.Context, id string) (*user.User, error) {

	return nil, nil
}

func (d *db) Update(ctx context.Context, u *user.User) error {
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	return nil
}
