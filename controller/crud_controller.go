package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/helper"
	"github.com/danangkonang/crud-rest/model"
	"github.com/danangkonang/crud-rest/service"
)

func NewAnimalHandler(db *config.DB) *Animal {
	return &Animal{
		Service: service.NewServiceAnimal(db),
	}
}

type Animal struct {
	Service service.AnimalService
}

func (p *Animal) AnimalCreate(w http.ResponseWriter, r *http.Request) {
	var animal model.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()
	animal.ID = helper.UnixRandomString(8)
	animal.CreateAt = time.Now()
	animal.UpdateAt = time.Now()
	if err := p.Service.SaveAnimal(&animal); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", animal)
}

func (p *Animal) AnimalShow(w http.ResponseWriter, r *http.Request) {
	res, err := p.Service.FindAnimal()
	if err != nil {
		helper.MakeRespon(w, 500, err.Error(), nil)
	}
	if len(res) == 0 {
		helper.MakeRespon(w, 200, "success", make([]string, 0))
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (p *Animal) AnimalDetail(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	// var animal model.Animal
	id := r.FormValue("id")
	if id == "" {
		helper.MakeRespon(w, 400, "params id required", nil)
		return
	}
	// p.Service.DetailAnimal(id)
	// animal.ID = id
	res, err := p.Service.DetailAnimal(id)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (p *Animal) AnimalEdit(w http.ResponseWriter, r *http.Request) {
	var animal model.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, "form required", nil)
		return
	}
	defer r.Body.Close()
	animal.UpdateAt = time.Now()
	if err := p.Service.UpdateAnimal(&animal); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", animal)
}

func (p *Animal) AnimalDelete(w http.ResponseWriter, r *http.Request) {
	var animal model.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, "form required", nil)
		return
	}
	defer r.Body.Close()
	if err := p.Service.DeleteAnimal(&animal); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "animal di hapus", nil)
}
