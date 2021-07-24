package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi-basic/helper"
	"restapi-basic/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Modifier interface{}

func InsertModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var Modifier model.TkModifierParent
	err_cat := json.NewDecoder(r.Body).Decode(&Modifier)

	if err_cat != nil {
		helper.Failed(err,"Failed to Decode JSON Body")
		log.Fatal(err_cat)
	}

	err_insert := Modifier.InsertModifierParent(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Insert Modifier")
		log.Fatal(err_insert)
	}

	response := helper.Success(Modifier,nil,"Insert Modifier successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func SelectAllModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var Modifier model.TkModifier

	result, err_insert := Modifier.SelectAllModifier(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Select All Modifier")
		log.Fatal(err_insert)
	}

	response := helper.Success(result,nil,"Select all Modifier successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func SelectOneModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Modifier_id := params["id"]
	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var Modifier model.TkModifier
	res_uuid, err := uuid.Parse(Modifier_id)
	if err != nil{
		log.Fatal(err)
	}
	Modifier.ModifierId = res_uuid
	result, err_insert := Modifier.SelectOneModifier(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Select All Modifier")
		log.Fatal(err_insert)
	}

	response := helper.Success(result,nil,"Select all Modifier successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func UpdateModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var Modifier model.TkModifier
	err_cat := json.NewDecoder(r.Body).Decode(&Modifier)

	if err_cat != nil {
		helper.Failed(err,"Failed to Decode JSON Body")
		log.Fatal(err_cat)
	}
	Modifier.Products = nil
	Modifier.LocationId = uuid.Nil
	Modifier.Isactive = "1"

	cat, err_update := Modifier.UpdateModifier(db)
	if err_update != nil {
		helper.Failed(err,"Failed to Update Modifier")
		log.Fatal(err_update)
	}

	response := helper.Success(cat,nil,"Update Modifier successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}

func DeleteModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Modifier_id := params["id"]
	if err != nil {
		helper.Failed(err,"Failed to Connect Database")
		log.Fatal(err)
	}
	var Modifier model.TkModifier
	res_uuid, err := uuid.Parse(Modifier_id)
	if err != nil{
		log.Fatal(err)
	}
	Modifier.ModifierId = res_uuid
	err_insert := Modifier.DeleteModifier(db)
	if err_insert != nil {
		helper.Failed(err,"Failed to Delete Modifier")
		log.Fatal(err_insert)
	}

	response := helper.Success(nil,nil,"Delete Modifier successfully")
	json,err_json := json.Marshal(response)
	if err_json != nil {
		helper.Failed(err,"Failed to create response")
		log.Fatal(err_json)
	}
	w.Write(json)
}


