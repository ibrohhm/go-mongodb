package handler

import (
	"encoding/json"
	"net/http"

	"io/ioutil"

	"github.com/go-mongo/entity"
	"github.com/go-mongo/repository"
	"github.com/go-mongo/utility/response"
	"github.com/julienschmidt/httprouter"
	"github.com/kataras/i18n"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	Repo repository.ProductRepositorier
}

func NewProductHandler(repo repository.ProductRepositorier) *ProductHandler {
	return &ProductHandler{
		Repo: repo,
	}
}

func (p *ProductHandler) Insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return response.WriteError(w, err)
	}

	var product entity.Product
	err = json.Unmarshal(b, &product)
	if err != nil {
		return response.WriteError(w, err)
	}

	newProduct, err := p.Repo.Insert(product)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, newProduct, i18n.Tr("en", "en.message.inserted", "product"))
}

func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	products, err := p.Repo.GetAll()
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, products, i18n.Tr("en", "en.message.succeed", "product"))
}

func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	_id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		return response.WriteError(w, err)
	}

	product, err := p.Repo.Get(_id)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, product, i18n.Tr("en", "en.message.succeed", "product"))
}

func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	_id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		return response.WriteError(w, err)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return response.WriteError(w, err)
	}

	var product entity.Product
	err = json.Unmarshal(b, &product)
	if err != nil {
		return response.WriteError(w, err)
	}

	newProduct, err := p.Repo.Update(_id, product)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, newProduct, i18n.Tr("en", "en.message.updated", "product", _id.Hex()))
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	_id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		return response.WriteError(w, err)
	}

	err = p.Repo.Delete(_id)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, nil, i18n.Tr("en", "en.message.deleted", "product", _id.Hex()))
}
