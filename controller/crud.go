package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danangkonang/crud-rest/helper"
	"github.com/danangkonang/crud-rest/service"
)

func AnimalCreate(w http.ResponseWriter, r *http.Request) {
	var animal service.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()
	animal.ID = helper.UnixRandomString(8)
	animal.CreateAt = time.Now()
	animal.UpdateAt = time.Now()
	if err := service.SaveAnimal(&animal); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", animal)
}

func AnimalShow(w http.ResponseWriter, r *http.Request) {
	res, err := service.FindAnimal()
	if err != nil {
		helper.MakeRespon(w, 500, err.Error(), nil)
	}
	if len(res) == 0 {
		helper.MakeRespon(w, 200, "success", make([]string, 0))
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func AnimalDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var animal service.Animal
	id := r.FormValue("id")
	if id == "" {
		helper.MakeRespon(w, 400, "params id required", nil)
		return
	}
	animal.ID = id
	err := service.DetailAnimal(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, "id not found", nil)
		return
	}
	helper.MakeRespon(w, 200, "success", animal)
}

func AnimalEdit(w http.ResponseWriter, r *http.Request) {
	var animal service.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, "form required", nil)
		return
	}
	defer r.Body.Close()
	animal.UpdateAt = time.Now()
	if err := service.UpdateAnimal(&animal); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", animal)
}

func AnimalDelete(w http.ResponseWriter, r *http.Request) {
	var animal service.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		helper.MakeRespon(w, 400, "form required", nil)
		return
	}
	defer r.Body.Close()
	if err := service.DeleteAnimal(&animal); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "animal di hapus", nil)
}
