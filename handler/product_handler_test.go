package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-mongo/config"
	"github.com/go-mongo/entity"
	"github.com/go-mongo/handler"
	"github.com/go-mongo/mock"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ = config.Initialize("../")

type HandlerSuite struct {
	suite.Suite
	mocker  *gomock.Controller
	product entity.Product
}

func TestHandlerSuite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := &HandlerSuite{}
	s.mocker = ctrl
	suite.Run(t, s)
}

func (suite *HandlerSuite) SetupTest() {
	suite.product = entity.Product{
		ID:     primitive.NewObjectID(),
		Name:   "product_name",
		Price:  1000,
		Weight: 1,
		Stock:  10,
	}
}

// insert
func (suite *HandlerSuite) TestInsertSuccess() {
	reqBody, _ := json.Marshal(suite.product)
	req, _ := http.NewRequest("POST", "/products", strings.NewReader(string(reqBody)))
	repo := mock.NewMockProductRepositorier(suite.mocker)
	repo.EXPECT().Insert(gomock.Any()).Return(suite.product, nil)

	h := handler.NewProductHandler(repo)
	err := h.Insert(httptest.NewRecorder(), req, nil)
	assert.Equal(suite.T(), err, nil)
}

func (suite *HandlerSuite) TestInsertError() {
	reqBody, _ := json.Marshal(suite.product)
	req, _ := http.NewRequest("POST", "/products", strings.NewReader(string(reqBody)))
	repo := mock.NewMockProductRepositorier(suite.mocker)
	errExpected := errors.New("error")
	repo.EXPECT().Insert(gomock.Any()).Return(suite.product, errExpected)

	h := handler.NewProductHandler(repo)
	err := h.Insert(httptest.NewRecorder(), req, nil)
	assert.Equal(suite.T(), err, errExpected)
}

// get all
func (suite *HandlerSuite) TestGetAllSuccess() {
	req, _ := http.NewRequest("GET", "/products", nil)
	repo := mock.NewMockProductRepositorier(suite.mocker)
	repo.EXPECT().GetAll().Return(nil, nil)

	h := handler.NewProductHandler(repo)
	err := h.GetAll(httptest.NewRecorder(), req, nil)
	assert.Equal(suite.T(), err, nil)
}

func (suite *HandlerSuite) TestGetAllError() {
	req, _ := http.NewRequest("GET", "/products", nil)
	repo := mock.NewMockProductRepositorier(suite.mocker)
	errExpected := errors.New("error")
	repo.EXPECT().GetAll().Return(nil, errExpected)

	h := handler.NewProductHandler(repo)
	err := h.GetAll(httptest.NewRecorder(), req, nil)
	assert.Equal(suite.T(), err, errExpected)
}

// get
func (suite *HandlerSuite) TestGetSuccess() {
	id := suite.product.ID.Hex()
	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	req, _ := http.NewRequest("GET", "/products/"+id, nil)
	repo := mock.NewMockProductRepositorier(suite.mocker)
	repo.EXPECT().GetAll().Return(nil, nil)

	h := handler.NewProductHandler(repo)
	err := h.GetAll(httptest.NewRecorder(), req, params)
	assert.Equal(suite.T(), err, nil)
}

func (suite *HandlerSuite) TestGetError() {
	id := suite.product.ID.Hex()
	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	req, _ := http.NewRequest("GET", "/products/"+id, nil)
	repo := mock.NewMockProductRepositorier(suite.mocker)
	errExpected := errors.New("error")
	repo.EXPECT().GetAll().Return(nil, errExpected)

	h := handler.NewProductHandler(repo)
	err := h.GetAll(httptest.NewRecorder(), req, params)
	assert.Equal(suite.T(), err, errExpected)
}

// update
func (suite *HandlerSuite) TestUpdateSuccess() {
	id := suite.product.ID.Hex()
	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	reqBody, _ := json.Marshal(suite.product)
	req, _ := http.NewRequest("PUT", "/products/"+id, strings.NewReader(string(reqBody)))
	repo := mock.NewMockProductRepositorier(suite.mocker)
	repo.EXPECT().Update(gomock.Any(), suite.product).Return(suite.product, nil)

	h := handler.NewProductHandler(repo)
	err := h.Update(httptest.NewRecorder(), req, params)
	assert.Equal(suite.T(), err, nil)
}

func (suite *HandlerSuite) TestUpdateError() {
	id := suite.product.ID.Hex()
	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	reqBody, _ := json.Marshal(suite.product)
	req, _ := http.NewRequest("PUT", "/products/"+id, strings.NewReader(string(reqBody)))
	repo := mock.NewMockProductRepositorier(suite.mocker)
	errExpected := errors.New("error")
	repo.EXPECT().Update(gomock.Any(), suite.product).Return(suite.product, errExpected)

	h := handler.NewProductHandler(repo)
	err := h.Update(httptest.NewRecorder(), req, params)
	assert.Equal(suite.T(), err, errExpected)
}

// delete
func (suite *HandlerSuite) TestDeleteSuccess() {
	id := suite.product.ID.Hex()
	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	req, _ := http.NewRequest("DELETE", "/products/"+id, nil)
	repo := mock.NewMockProductRepositorier(suite.mocker)
	repo.EXPECT().Delete(gomock.Any()).Return(nil)

	h := handler.NewProductHandler(repo)
	err := h.Delete(httptest.NewRecorder(), req, params)
	assert.Equal(suite.T(), err, nil)
}

func (suite *HandlerSuite) TestDeleteError() {
	id := suite.product.ID.Hex()
	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	req, _ := http.NewRequest("DELETE", "/products/"+id, nil)
	repo := mock.NewMockProductRepositorier(suite.mocker)
	errExpected := errors.New("error")
	repo.EXPECT().Delete(gomock.Any()).Return(errExpected)

	h := handler.NewProductHandler(repo)
	err := h.Delete(httptest.NewRecorder(), req, params)
	assert.Equal(suite.T(), err, errExpected)
}
