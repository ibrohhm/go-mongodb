package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Price  float64            `bson:"price" json:"price"`
	Weight int                `bson:"weight" json:"weight"`
	Stock  int                `bson:"stock" json:"stock"`
}
