package repository

import (
	"context"

	"github.com/go-mongo/database"
	"github.com/go-mongo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	db database.ProductDBer
}

// mockgen --source=repository/product_repository.go --destination=mock/product_repository.go -package=mock
type ProductRepositorier interface {
	Insert(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Get(_id primitive.ObjectID) (entity.Product, error)
	Update(_id primitive.ObjectID, product entity.Product) (entity.Product, error)
	Delete(_id primitive.ObjectID) error
}

func NewProductRepository(db database.ProductDBer) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Insert(product entity.Product) (entity.Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.db.InsertOne(context.TODO(), product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	findOptions := options.Find()
	// findOptions.SetLimit(5)
	cur, err := p.db.Find(context.TODO(), bson.D{}, findOptions)

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

	err := p.db.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) Update(_id primitive.ObjectID, product entity.Product) (entity.Product, error) {
	product.ID = _id
	update := bson.D{{"$set", product}}
	_, err := p.db.UpdateOne(context.TODO(), bson.M{"_id": _id}, update)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) Delete(_id primitive.ObjectID) error {
	_, err := p.db.DeleteOne(context.TODO(), bson.D{{"_id", _id}})
	if err != nil {
		return err
	}

	return nil
}
