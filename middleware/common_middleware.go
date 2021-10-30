package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi-basic/helper"
	"restapi-basic/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type M map[string]interface{}

var APPLICATION_NAME = "TokoKocak"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("kocakersjayajayajaya")

type UserClaims struct {
	jwt.StandardClaims
	UserId    uuid.UUID `json:"userId"`
	UserEmail string    `json:"userEmail"`
	UserRole  string    `json:"userRole"`
}

func GenerateJWT(userData model.TkUser) string {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UserId:    userData.UserId,
		UserEmail: userData.UserEmail,
		UserRole:  userData.UserRole,
	}
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, _ := token.SignedString(JWT_SIGNATURE_KEY)
	return signedToken
}

func MiddlewareJWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			resToken := r.Header.Get("Authorization")
			if !strings.Contains(resToken, "Bearer") {
				response := helper.Failed("Token invalid", "Token invalid")
				json, _ := json.Marshal(response)
				rw.Write(json)
			} else {
				tokenString := strings.Replace(resToken, "Bearer ", "", -1)
				token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
					if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Signing method invalid")
					} else if method != JWT_SIGNING_METHOD {
						return nil, fmt.Errorf("Signing method invalid")
					}
					err_claim := t.Claims.Valid()
					if err_claim != nil {
						return nil, err_claim
					}
					return JWT_SIGNATURE_KEY, nil
				})

				if err != nil {
					http.Error(rw, err.Error(), http.StatusBadRequest)
					return
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok || !token.Valid {
					http.Error(rw, err.Error(), http.StatusBadRequest)
					return
				}
				ctx := context.WithValue(context.Background(), "userInfo", claims)
				r = r.WithContext(ctx)
				handler.ServeHTTP(rw, r)
			}
		} else {
			response := helper.Failed("Unauthorized", "Unauthorized")
			json, _ := json.Marshal(response)
			rw.Write(json)
		}
	})
}
