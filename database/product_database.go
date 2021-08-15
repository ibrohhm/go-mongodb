package database

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductDB struct {
	dbName     string
	colName    string
	collection *mongo.Collection
}

type ProductDBer interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

func NewProductDB(client *mongo.Client, dbName string, colName string) *ProductDB {
	return &ProductDB{
		dbName:     dbName,
		colName:    colName,
		collection: client.Database(dbName).Collection(colName),
	}
}

func (p *ProductDB) sendLogger(command string, filter interface{}) {
	log.WithFields(log.Fields{
		"command":    command,
		"filter":     filter,
		"collection": p.colName,
		"database":   p.dbName,
	}).Info()
}

func (p *ProductDB) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	p.sendLogger("InsertOne", nil)
	return p.collection.InsertOne(ctx, document, opts...)
}

func (p *ProductDB) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	p.sendLogger("Find", filter)
	return p.collection.Find(ctx, filter, opts...)
}

func (p *ProductDB) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	p.sendLogger("FindOne", filter)
	return p.collection.FindOne(ctx, filter, opts...)
}

func (p *ProductDB) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	p.sendLogger("UpdateOne", filter)
	return p.collection.UpdateOne(ctx, filter, update, opts...)
}

func (p *ProductDB) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	p.sendLogger("DeleteOne", filter)
	return p.collection.DeleteOne(ctx, filter, opts...)
}
