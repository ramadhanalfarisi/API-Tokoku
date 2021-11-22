package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi-basic/helper"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ParameterMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		location_id := params["id"]
		if r.Method == "PUT" || r.Method == "DELETE" {
			if location_id == "" {
				response := helper.Failed("Params invalid", "Params {id} is required")

				json, err := json.Marshal(response)

				if err != nil {
					log.Fatal(err)
				}
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write(json)
				return
			}
		}
		if location_id != "" {
			_, err := uuid.Parse(location_id)
			if err != nil {
				response := helper.Failed("Params invalid", "Params {id} invalid")

				json, err := json.Marshal(response)

				if err != nil {
					log.Fatal(err)
				}
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write(json)
				return
			}
		}
		rw.Header().Add("Content-Type", "application/json")
		handler.ServeHTTP(rw, r)
	})
}
