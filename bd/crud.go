package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Create(entity interface{}, collectionName string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection(collectionName)

	result, err := col.InsertOne(ctx, entity)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

func FindOne(condition interface{}, collectionName string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection(collectionName)

	return col.FindOne(ctx, condition)
}

func UpdateOne(filter interface{}, fields interface{}, collectionName string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection(collectionName)

	return col.UpdateOne(ctx, filter, fields)
}

func FindAllByCondition(filter interface{}, options *options.FindOptions, collectionName string) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection(collectionName)

	return col.Find(ctx, filter, options)
}

func DeleteOne(filter interface{}, collectionName string) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection(collectionName)

	_, err := col.DeleteOne(ctx, filter)

	return err
}