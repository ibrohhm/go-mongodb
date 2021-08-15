package repository

import (
	"context"

	"github.com/go-mongo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client) *ProductRepository {
	return &ProductRepository{
		collection: client.Database("go_mongo_learn").Collection("product"),
	}
}

func (p *ProductRepository) Insert(product entity.Product) (entity.Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	findOptions := options.Find()
	// findOptions.SetLimit(5)
	cur, err := p.collection.Find(context.Background(), bson.D{}, findOptions)

	if err != nil {
		return []entity.Product{}, err
	}

	for cur.Next(context.TODO()) {
		var product entity.Product
		err = cur.Decode(&product)
		if err != nil {
			return []entity.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *ProductRepository) Get(_id primitive.ObjectID) (entity.Product, error) {
	var product entity.Product

	err := p.collection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) Update(_id primitive.ObjectID, product entity.Product) (entity.Product, error) {
	product.ID = _id
	update := bson.D{{"$set", product}}
	_, err := p.collection.UpdateOne(context.TODO(), bson.M{"_id": _id}, update)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) Delete(_id primitive.ObjectID) error {
	_, err := p.collection.DeleteOne(context.TODO(), bson.D{{"_id", _id}})
	if err != nil {
		return err
	}

	return nil
}
