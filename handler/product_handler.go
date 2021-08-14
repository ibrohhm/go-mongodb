package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/go-mongo/entity"
	"github.com/go-mongo/repository"
	"github.com/go-mongo/utility/response"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	Repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		Repo: repo,
	}
}

func (p *ProductHandler) Healthz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("ok")
}

func (p *ProductHandler) Insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.WriteError(w, err)
		return
	}

	var product entity.Product
	err = json.Unmarshal(b, &product)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	fmt.Printf("%+v", product)

	newProduct, err := p.Repo.Insert(product)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteSuccess(w, newProduct, "")
}

func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	products, err := p.Repo.GetAll()
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteSuccess(w, products, "")
}

func (p *ProductHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		response.WriteError(w, err)
		return
	}

	product, err := p.Repo.Get(_id)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteSuccess(w, product, "")
}

func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		response.WriteError(w, err)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.WriteError(w, err)
		return
	}

	var product entity.Product
	err = json.Unmarshal(b, &product)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	fmt.Printf("%+v", product)

	product, err = p.Repo.Update(_id, product)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteSuccess(w, product, "")
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_id, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		response.WriteError(w, err)
		return
	}

	err = p.Repo.Delete(_id)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteSuccess(w, nil, "success")
}
