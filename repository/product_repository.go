package repository

import (
	"context"
	"fmt"

	"github.com/go-mongo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	client     *mongo.Client
	collection string
}

func NewProductRepository(client *mongo.Client) *ProductRepository {
	return &ProductRepository{
		client:     client,
		collection: "product",
	}
}

func (p *ProductRepository) Insert(product entity.Product) (entity.Product, error) {
	collection := p.client.Database("go_mongo_learn").Collection("product")
	product.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return entity.Product{}, err
	}

	fmt.Printf("insert id: %+v", res.InsertedID)
	return product, nil
}

func (p *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	collection := p.client.Database("go_mongo_learn").Collection("product")
	findOptions := options.Find()
	// findOptions.SetLimit(5)
	cur, err := collection.Find(context.Background(), bson.D{}, findOptions)

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
	collection := p.client.Database("go_mongo_learn").Collection("product")
	var product entity.Product

	err := collection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) Update(_id primitive.ObjectID, product entity.Product) (entity.Product, error) {
	collection := p.client.Database("go_mongo_learn").Collection("product")
	filter := bson.M{"_id": _id}
	update := bson.D{{"$set", bson.D{{"name", product.Name}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return entity.Product{}, err
	}

	product.ID = _id
	return product, nil
}

func (p *ProductRepository) Delete(_id primitive.ObjectID) error {
	collection := p.client.Database("go_mongo_learn").Collection("product")

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", _id}})
	if err != nil {
		return err
	}

	return nil
}
