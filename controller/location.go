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

func InsertLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}

	var location model.TkLocation
	req := json.NewDecoder(r.Body).Decode(&location)
	if req != nil {
		helper.Failed(req, "Failed to decode JSON")
		log.Fatal(req)
	}
	location.LocationId = uuid.New()

	res := location.InsertLocation(db)
	if res != nil {
		helper.Failed(res, "Failed to insert location")
		log.Fatal(res)
	}

	response := helper.Success(location, nil, "Sucesse to insert location")
	json, err_json := json.Marshal(response)

	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func GetAllLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}

	var location model.TkLocation
	result, err := location.SelectAllLocation(db)

	if err != nil {
		log.Fatal(err)
	}

	response := helper.Success(result, err, "Select location successfully")
	json, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
	defer helper.CloseConnection(db)
}

func GetOneLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	location_id := params["id"]

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}

	var location model.TkLocation
	loc_uuid, err := uuid.Parse(location_id)

	if err != nil {
		log.Fatal(err)
	}

	location.LocationId = loc_uuid
	result, err := location.SelectOneLocation(db)

	if err != nil {
		log.Fatal(err)
	}

	response := helper.Success(result, err, "Select location successfully")
	json, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
	defer helper.CloseConnection(db)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		log.Fatal(err)
	}

	var location model.TkLocation
	err_json := json.NewDecoder(r.Body).Decode(&location)

	if err_json != nil {
		log.Fatal(err)
	}

	res, err := location.UpdateLocation(db)

	if err != nil {
		log.Fatal(err)
	}
	response := helper.Success(res, err, "Update location successfully")
	json_res, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json_res)
	defer helper.CloseConnection(db)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	location_id := params["id"]

	if err != nil {
		helper.Failed(err, "Failed to Connect Database")
		log.Fatal(err)
	}

	var location model.TkLocation
	loc_uuid, err := uuid.Parse(location_id)

	if err != nil {
		log.Fatal(err)
	}

	location.LocationId = loc_uuid
	res_err := location.DeleteLocation(db)

	if res_err != nil {
		log.Fatal(res_err)
	}

	response := helper.Success(nil, err, "Delete location successfully")
	json, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
	defer helper.CloseConnection(db)
}
