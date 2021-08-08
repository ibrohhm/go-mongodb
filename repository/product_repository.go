package repository

import "gopkg.in/mgo.v2"

type ProductRepository struct {
	db *mgo.Session
}

func NewProductRepository(db *mgo.Session) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
