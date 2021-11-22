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

func InsertLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	if err != nil {

		log.Fatal(err)
	}

	var location model.TkLocation
	req := json.NewDecoder(r.Body).Decode(&location)
	if req != nil {

		log.Fatal(req)
	}

	errorValidation := helper.Validate(location)

	if errorValidation != nil {
		response := helper.FailedValidate("Invalid Data", errorValidation)
		json, err_json := json.Marshal(response)

		if err_json != nil {
			log.Fatal(err_json)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		location.LocationId = uuid.New()
		userId, _ := uuid.Parse(userInfo["userId"].(string))
		location.UserId = userId

		res := location.InsertLocation(db)
		if res != nil {

			log.Fatal(res)
		}

		response := helper.Success(location, nil, "Sucesse to insert location")
		json, err_json := json.Marshal(response)

		if err_json != nil {
			log.Fatal(err_json)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}

	defer helper.CloseConnection(db)
}

func GetAllLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	if err != nil {

		log.Fatal(err)
	}

	var location model.TkLocation
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	location.UserId = userId
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
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	params := mux.Vars(r)
	location_id := params["id"]

	if err != nil {
		log.Fatal(err)
	}

	var location model.TkLocation
	loc_uuid, err := uuid.Parse(location_id)
	location.LocationId = loc_uuid
	userId, _ := uuid.Parse(userInfo["userId"].(string))
	location.UserId = userId
	result, err := location.SelectOneLocation(db)

	if err != nil {
		log.Fatal(err)
	}

	response := helper.Success(result, err, "Select location successfully")

	json, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)

	defer helper.CloseConnection(db)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	params := mux.Vars(r)
	location_id := params["id"]
	if err != nil {
		log.Fatal(err)
	}

	var location model.TkLocation
	err_json := json.NewDecoder(r.Body).Decode(&location)

	errorValidation := helper.Validate(location)

	if errorValidation != nil {
		response := helper.FailedValidate("Invalid Data", errorValidation)
		json, err_json := json.Marshal(response)

		if err_json != nil {
			log.Fatal(err_json)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		loc_uuid, err := uuid.Parse(location_id)
		location.UserId = loc_uuid
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
		w.WriteHeader(http.StatusOK)
		w.Write(json_res)
	}

	defer helper.CloseConnection(db)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	location_id := params["id"]

	if err != nil {

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
