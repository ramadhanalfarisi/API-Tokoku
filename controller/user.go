package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"restapi-basic/helper"
	"restapi-basic/middleware"
	"restapi-basic/model"

	"github.com/google/uuid"
)

type ResponseJWT struct {
	Token   string `json:"token"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Register(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	if err != nil {
		log.Fatal(err)
	}

	var user model.TkUser
	res := json.NewDecoder(r.Body).Decode(&user)
	user.UserId = uuid.New()
	user.UserPassword = getMD5Hash(user.UserPassword)
	if res != nil {
		log.Fatal(res)
	}

	err_insert := user.InsertUser(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(nil, nil, "Success for register account")
	json, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
	defer helper.CloseConnection(db)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	if err != nil {
		log.Fatal(err)
	}

	var user model.TkUser
	res := json.NewDecoder(r.Body).Decode(&user)
	user.UserPassword = getMD5Hash(user.UserPassword)
	if res != nil {
		log.Fatal(res)
	}

	result, err := user.LoginUser(db)
	if err != nil {
		log.Fatal(err)
	}

	jwt := middleware.GenerateJWT(result)
	response := ResponseJWT{
		Token:   jwt,
		Status:  "success",
		Message: "Login Successfully",
	}
	json, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}
