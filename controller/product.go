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
	Product.ProductName = r.FormValue("productName")
	Product.ProductDesc = r.FormValue("productDesc")
	Product.ProductPrice = r.FormValue("productPrice")

	res_uuid_cat, err := uuid.Parse(r.FormValue("categoryId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.CategoryId = res_uuid_cat

	res_uuid_loc, err := uuid.Parse(r.FormValue("locationId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.LocationId = res_uuid_loc

	arr_modifiers := r.Form["modifiers"]
	var modifiers []model.TkProductModifier
	for _, arr := range arr_modifiers {
		res_uuid_mod, err := uuid.Parse(arr)
		if err != nil {
			log.Fatal(err)
		}
		var modifier model.TkProductModifier
		modifier.ProductId = Product.ProductId
		modifier.ModifierId = res_uuid_mod
		modifiers = append(modifiers, modifier)
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

	filename := os.Getenv("BASE_URL") + "/files/" + Product.ProductId.String()
	err_upload := helper.Uploader(uploadedFile, handler, Product.ProductId.String())

	if err_upload != nil {
		log.Fatal(err_upload)
	}

	Product.ProductImage = filename

	err_insert := Product.InsertProduct(db)
	if err_insert != nil {
		helper.Failed(err, "Failed to Insert Product")
		log.Fatal(err_insert)
	}
	
	var productModifier model.TkProductModifier
	err_insert_mod := productModifier.InsertProductModifier(db, modifiers)
	if err_insert_mod != nil {
		helper.Failed(err, "Failed to Insert Modifier")
		log.Fatal(err_insert_mod)
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
	res_uuid_prod, err := uuid.Parse(r.FormValue("productId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.ProductId = res_uuid_prod
	Product.ProductName = r.FormValue("productName")
	Product.ProductDesc = r.FormValue("productDesc")
	Product.ProductPrice = r.FormValue("productPrice")

	res_uuid_cat, err := uuid.Parse(r.FormValue("categoryId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.CategoryId = res_uuid_cat

	res_uuid_loc, err := uuid.Parse(r.FormValue("locationId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.LocationId = res_uuid_loc

	arr_modifiers := r.Form["modifiers"]
	var modifiers []model.TkProductModifier
	for _, arr := range arr_modifiers {
		res_uuid_mod, err := uuid.Parse(arr)
		if err != nil {
			log.Fatal(err)
		}
		var modifier model.TkProductModifier
		modifier.ProductId = Product.ProductId
		modifier.ModifierId = res_uuid_mod
		modifiers = append(modifiers, modifier)
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

	filename := os.Getenv("BASE_URL") + "/files/" + Product.ProductId.String()
	err_upload := helper.Uploader(uploadedFile, handler, Product.ProductId.String())

	if err_upload != nil {
		log.Fatal(err_upload)
	}

	Product.ProductImage = filename
	Product.LocationId = uuid.Nil

	cat, err_update := Product.UpdateProduct(db)
	if err_update != nil {
		helper.Failed(err, "Failed to Update Product")
		log.Fatal(err_update)
	}

	var productModifier model.TkProductModifier
	productModifier.ProductId = res_uuid_prod
	err_delete_mod := productModifier.DeleteProductModifier(db)
	if err_delete_mod != nil {
		helper.Failed(err, "Failed to Insert Modifier")
		log.Fatal(err_delete_mod)
	}

	err_insert_mod := productModifier.InsertProductModifier(db, modifiers)
	if err_insert_mod != nil {
		helper.Failed(err, "Failed to Insert Modifier")
		log.Fatal(err_insert_mod)
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
