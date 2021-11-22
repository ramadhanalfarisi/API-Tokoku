package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"restapi-basic/helper"
	"restapi-basic/model"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Product interface{}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {

		log.Fatal(err)
	}
	var Product model.TkProduct

	Product.ProductId = uuid.New()
	Product.ProductName = r.FormValue("productName")
	Product.ProductDesc = r.FormValue("productDesc")
	productPrice, _ := strconv.ParseFloat(r.FormValue("productPrice"), 64)
	Product.ProductPrice = productPrice

	res_uuid_cat, err := uuid.Parse(r.FormValue("categoryId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.CategoryId = res_uuid_cat

	res_uuid_loc, err := uuid.Parse(r.FormValue("locationId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.UserId = res_uuid_loc

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
	if uploadedFile != nil {
		filename := os.Getenv("BASE_URL") + "/files/" + Product.ProductId.String()
		err_upload := helper.Uploader(uploadedFile, handler, Product.ProductId.String())

		if err_upload != nil {
			log.Fatal(err_upload)
		}

		Product.ProductImage = filename
	}

	validateProducts := helper.Validate(Product)
	if validateProducts != nil {
		response := helper.FailedValidate("Invalid data", validateProducts)
		json, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		err_insert := Product.InsertProduct(db)
		if err_insert != nil {

			log.Fatal(err_insert)
		}

		var productModifier model.TkProductModifier
		err_insert_mod := productModifier.InsertProductModifier(db, modifiers)
		if err_insert_mod != nil {

			log.Fatal(err_insert_mod)
		}

		response := helper.Success(Product, nil, "Insert Product successfully")
		json, err_json := json.Marshal(response)
		if err_json != nil {

			log.Fatal(err_json)
		}
		w.Write(json)
	}

	defer helper.CloseConnection(db)
}

func SelectAllProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {

		log.Fatal(err)
	}
	var Product model.TkProduct

	result, err_insert := Product.SelectAllProduct(db)
	if err_insert != nil {

		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {

		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func SelectOneProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Product_id := params["id"]
	if err != nil {

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

		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {

		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {

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
	productPrice, _ := strconv.ParseFloat(r.FormValue("productPrice"), 64)
	Product.ProductPrice = productPrice

	res_uuid_cat, err := uuid.Parse(r.FormValue("categoryId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.CategoryId = res_uuid_cat

	res_uuid_loc, err := uuid.Parse(r.FormValue("locationId"))
	if err != nil {
		log.Fatal(err)
	}
	Product.UserId = res_uuid_loc

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
	if uploadedFile != nil {
		// delete exist file
		getProduct, _ := Product.SelectOneProduct(db)
		photoProduct := getProduct.ProductImage
		e := os.Remove(photoProduct)
		if e != nil {
			log.Fatal(e)
		}

		// upload new file
		filename := os.Getenv("BASE_URL") + "/files/" + Product.ProductId.String()
		err_upload := helper.Uploader(uploadedFile, handler, Product.ProductId.String())

		if err_upload != nil {
			log.Fatal(err_upload)
		}
		Product.ProductImage = filename
	}

	Product.UserId = uuid.Nil

	validateProducts := helper.Validate(Product)
	if validateProducts != nil {
		response := helper.FailedValidate("Invalid Data", validateProducts)
		json, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		cat, err_update := Product.UpdateProduct(db)
		if err_update != nil {

			log.Fatal(err_update)
		}

		var productModifier model.TkProductModifier
		productModifier.ProductId = res_uuid_prod
		err_delete_mod := productModifier.DeleteProductModifier(db)
		if err_delete_mod != nil {

			log.Fatal(err_delete_mod)
		}

		err_insert_mod := productModifier.InsertProductModifier(db, modifiers)
		if err_insert_mod != nil {

			log.Fatal(err_insert_mod)
		}

		response := helper.Success(cat, nil, "Update Product successfully")
		json, err_json := json.Marshal(response)
		if err_json != nil {

			log.Fatal(err_json)
		}
		w.Write(json)
	}

	defer helper.CloseConnection(db)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Product_id := params["id"]
	if err != nil {

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

		log.Fatal(err_insert)
	}

	response := helper.Success(nil, nil, "Delete Product successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {

		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}
