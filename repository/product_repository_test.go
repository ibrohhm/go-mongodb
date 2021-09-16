package repository_test

import (
	"errors"
	"testing"

	"github.com/go-mongo/entity"
	"github.com/go-mongo/mock"
	"github.com/go-mongo/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositorySuite struct {
	suite.Suite
	mocker *gomock.Controller
}

func TestRepositorySuite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := &RepositorySuite{}
	s.mocker = ctrl
	suite.Run(t, s)
}

func (suite *RepositorySuite) SetupTest() {
}

// Insert
func (suite *RepositorySuite) TestInsertSuccess() {
	dbMock := mock.NewMockProductDBer(suite.mocker)
	dbMock.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(nil, nil)

	r := repository.NewProductRepository(dbMock)
	_, err := r.Insert(entity.Product{})

	assert.Equal(suite.T(), err, nil)
}

func (suite *RepositorySuite) TestInsertError() {
	errReturn := errors.New("error")
	dbMock := mock.NewMockProductDBer(suite.mocker)
	dbMock.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(nil, errReturn)

	r := repository.NewProductRepository(dbMock)
	_, err := r.Insert(entity.Product{})

	assert.Equal(suite.T(), err, errReturn)
}

// Update
func (suite *RepositorySuite) TestUpdateSuccess() {
	id := primitive.NewObjectID()
	dbMock := mock.NewMockProductDBer(suite.mocker)
	dbMock.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.All()).Return(nil, nil)

	r := repository.NewProductRepository(dbMock)
	_, err := r.Update(id, entity.Product{})

	assert.Equal(suite.T(), err, nil)
}

func (suite *RepositorySuite) TestUpdateError() {
	id := primitive.NewObjectID()
	errReturn := errors.New("error")
	dbMock := mock.NewMockProductDBer(suite.mocker)
	dbMock.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.All()).Return(nil, errReturn)

	r := repository.NewProductRepository(dbMock)
	_, err := r.Update(id, entity.Product{})

	assert.Equal(suite.T(), err, errReturn)
}

// Delete
func (suite *RepositorySuite) TestDeleteSuccess() {
	id := primitive.NewObjectID()
	dbMock := mock.NewMockProductDBer(suite.mocker)
	dbMock.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil, nil)

	r := repository.NewProductRepository(dbMock)
	err := r.Delete(id)

	assert.Equal(suite.T(), err, nil)
}

func (suite *RepositorySuite) TestDeleteError() {
	id := primitive.NewObjectID()
	errReturn := errors.New("error")
	dbMock := mock.NewMockProductDBer(suite.mocker)
	dbMock.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil, errReturn)

	r := repository.NewProductRepository(dbMock)
	err := r.Delete(id)

	assert.Equal(suite.T(), err, errReturn)
}
