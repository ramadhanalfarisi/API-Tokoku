package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"restapi-basic/helper"
	"restapi-basic/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Product interface{}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}
	var Product model.TkProduct

	Product.ProductId = uuid.New()
	Product.ProductName = r.FormValue("product_name")
	Product.ProductDesc = r.FormValue("product_desc")
	Product.ProductPrice = r.FormValue("product_price")

	res_uuid_cat, err := uuid.Parse(r.FormValue("category_id"))
	if err != nil {
		log.Fatal(err)
	}
	Product.CategoryId = res_uuid_cat

	res_uuid_loc, err := uuid.Parse(r.FormValue("location_id"))
	if err != nil {
		log.Fatal(err)
	}
	Product.LocationId = res_uuid_loc

	arr_modifiers := r.Form["modifiers"]
	for i, arr := range arr_modifiers {
		res_uuid_mod, err := uuid.Parse(arr)
		if err != nil {
			log.Fatal(err)
		}
		Product.Modifiers[i].ModifierParentId = res_uuid_mod
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal(err_env)
	}
	
	filename := os.Getenv("BASE_URL") + "/files/"+uuid.
	helper.Uploader(uploadedFile,handler,)

	err_insert := Product.InsertProduct(db)
	if err_insert != nil {
		helper.Failed(err, "Failed to Insert Product")
		log.Fatal(err_insert)
	}

	response := helper.Success(Product, nil, "Insert Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err, "Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func SelectAllProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}
	var Product model.TkProduct

	result, err_insert := Product.SelectAllProduct(db)
	if err_insert != nil {
		helper.Failed(err, "Failed to Select All Product")
		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err, "Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func SelectOneProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Product_id := params["id"]
	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}
	var Product model.TkProduct
	res_uuid, err := uuid.Parse(Product_id)
	if err != nil {
		log.Fatal(err)
	}
	Product.ProductId = res_uuid
	result, err_insert := Product.SelectOneProduct(db)
	if err_insert != nil {
		helper.Failed(err, "Failed to Select All Product")
		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err, "Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}
	var Product model.TkProduct
	err_cat := json.NewDecoder(r.Body).Decode(&Product)

	if err_cat != nil {
		helper.Failed(err, "Failed to Decode JSON Body")
		log.Fatal(err_cat)
	}
	Product.LocationId = uuid.Nil

	cat, err_update := Product.UpdateProduct(db)
	if err_update != nil {
		helper.Failed(err, "Failed to Update Product")
		log.Fatal(err_update)
	}

	response := helper.Success(cat, nil, "Update Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err, "Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Product_id := params["id"]
	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}
	var Product model.TkProduct
	res_uuid, err := uuid.Parse(Product_id)
	if err != nil {
		log.Fatal(err)
	}
	Product.ProductId = res_uuid
	err_insert := Product.DeleteProduct(db)
	if err_insert != nil {
		helper.Failed(err, "Failed to Delete Product")
		log.Fatal(err_insert)
	}

	response := helper.Success(nil, nil, "Delete Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err, "Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}
