package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	animal.Image = helper.RamdomString() + ".jpg"
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
		return
	}
	if len(res) == 0 {
		helper.MakeRespon(w, 200, "success", make([]string, 0))
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (p *Animal) AnimalDetail(w http.ResponseWriter, r *http.Request) {
	var animal model.Animal
	id := r.FormValue("animal_id")
	if id == "" {
		helper.MakeRespon(w, 400, "params id required", nil)
		return
	}
	animal_id, err := strconv.Atoi(id)
	if err != nil {
		helper.MakeRespon(w, 400, "invalid id", nil)
		return
	}
	animal.ID = animal_id
	res, err := p.Service.DetailAnimal(animal.ID)
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
	helper.MakeRespon(w, 200, "success", nil)
}
