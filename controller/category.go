package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi-basic/helper"
	"restapi-basic/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Category interface{}

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	if err != nil {

		log.Fatal(err)
	}
	var category model.TkCategory
	err_cat := json.NewDecoder(r.Body).Decode(&category)

	if err_cat != nil {

		log.Fatal(err_cat)
	}
	category.Products = nil
	category.CategoryId = uuid.New()
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	category.UserId = userId

	validateCategories := helper.Validate(category)
	if validateCategories != nil {
		response := helper.FailedValidate("Invalid Data", validateCategories)
		json, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		err_insert := category.InsertCategory(db)
		if err_insert != nil {
			log.Fatal(err_insert)
		}

		response := helper.Success(category, nil, "Insert category successfully")
		json, err_json := json.Marshal(response)
		if err_json != nil {
			log.Fatal(err_json)
		}
		w.Write(json)
	}
	defer helper.CloseConnection(db)
}

func SelectAllCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	if err != nil {
		log.Fatal(err)
	}
	var category model.TkCategory
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	category.UserId = userId
	result, err_insert := category.SelectAllCategory(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all category successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func SelectOneCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	params := mux.Vars(r)
	category_id := params["id"]
	if err != nil {
		log.Fatal(err)
	}
	var category model.TkCategory
	res_uuid, err := uuid.Parse(category_id)
	if err != nil {
		log.Fatal(err)
	}
	category.CategoryId = res_uuid
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	category.UserId = userId
	result, err_select := category.SelectOneCategory(db)
	if err_select != nil {
		log.Fatal(err_select)
	}

	response := helper.Success(result, nil, "Select all category successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func SelectAllMenu(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	if err != nil {
		log.Fatal(err)
	}
	var category model.TkCategory
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	category.UserId = userId
	result, err_insert := category.SelectAllMenu(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all category successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	if err != nil {
		log.Fatal(err)
	}
	var category model.TkCategory
	err_cat := json.NewDecoder(r.Body).Decode(&category)

	if err_cat != nil {
		log.Fatal(err_cat)
	}
	category.Products = nil
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	category.UserId = userId

	validateCategories := helper.Validate(category)
	if validateCategories != nil {
		response := helper.FailedValidate("Invalid Data", validateCategories)
		json, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		cat, err_update := category.UpdateCategory(db)
		if err_update != nil {
			log.Fatal(err_update)
		}

		response := helper.Success(cat, nil, "Update category successfully")
		json, err_json := json.Marshal(response)
		if err_json != nil {
			log.Fatal(err_json)
		}
		w.Write(json)
	}
	defer helper.CloseConnection(db)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	params := mux.Vars(r)
	category_id := params["id"]
	if err != nil {
		log.Fatal(err)
	}
	var category model.TkCategory
	res_uuid, err := uuid.Parse(category_id)
	if err != nil {
		log.Fatal(err)
	}
	category.CategoryId = res_uuid
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	category.UserId = userId
	err_insert := category.DeleteCategory(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(nil, nil, "Delete category successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}
