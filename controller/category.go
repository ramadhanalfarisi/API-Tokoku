package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi-basic/helper"
	"restapi-basic/model"
	"github.com/google/uuid"
)

type Category interface{}

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var category model.TkCategory
	err_cat := json.NewDecoder(r.Body).Decode(&category)

	if err_cat != nil {
		helper.Failed(err,"Failed to Decode JSON Body")
		log.Fatal(err_cat)
	}
	category.Products = nil
	category.CategoryId = uuid.New()
	category.LocationId = uuid.New()

	err_insert := category.InsertCategory(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Insert Category")
		log.Fatal(err_insert)
	}

	response := helper.Success(category,nil,"Insert category successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func SelectAllCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var category model.TkCategory

	result, err_insert := category.SelectAllCategory(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Select All Category")
		log.Fatal(err_insert)
	}

	response := helper.Success(result,nil,"Select all category successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func SelectAllMenu(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var category model.TkCategory

	result, err_insert := category.SelectAllMenu(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Select All Category")
		log.Fatal(err_insert)
	}

	response := helper.Success(result,nil,"Select all category successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var category model.TkCategory
	err_cat := json.NewDecoder(r.Body).Decode(&category)

	if err_cat != nil {
		helper.Failed(err,"Failed to Decode JSON Body")
		log.Fatal(err_cat)
	}
	category.Products = nil
	category.LocationId = uuid.Nil
	category.Isactive = "1"

	category, err_update := category.UpdateCategory(db)
	if err_update != nil {
		helper.Failed(err,"Failed to Update Category")
		log.Fatal(err_update)
	}

	response := helper.Success(category,nil,"Update category successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}
