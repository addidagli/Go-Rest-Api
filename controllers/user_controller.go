package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"main/connections"
	"main/models"
	"net/http"
	"time"
)

const SecretKey = "secret"

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	db := connections.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&user)

	result := db.First(&user, "email = ?", user.Email)
	if result.RowsAffected > 0 {
		Message := []byte(`{"Error": "User already exist"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	} else {
		if error != nil {
			log.Fatal(error)
			connections.SendError(writer, http.StatusBadRequest)
			return
		}

		error = db.Create(&user).Error

		if error != nil {
			log.Fatal(error)
			connections.SendError(writer, http.StatusInternalServerError)
			return
		}

		json, _ := json.Marshal(user)

		connections.SendReponse(writer, http.StatusCreated, json)
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	db := connections.GetConnection()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result := db.First(&user, "(email,password) = (?,?)", user.Email, user.Password)
	if user.Token != "" {
		Message := []byte(`{"Error": "Already logged in"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	} else {
		if result.RowsAffected > 0 {
			expirationTime := time.Now().Add(time.Minute * 5)

			claims := &Claims{
				Username: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			user.Token = tokenString

			db.Model(&user).Where("email = ?", user.Email).Update("token", user.Token)

			http.SetCookie(writer,
				&http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
				})
			json, _ := json.Marshal(user)

			connections.SendReponse(writer, http.StatusOK, json)
		} else {
			Message := []byte(`{"Error": "Email or password is incorrect"}`)
			connections.ErrorMsg(writer, http.StatusBadRequest, Message)
		}
	}

}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := connections.GetConnection()
	defer db.Close()

	db.First(&user, "id = ?", id)

	if user.ID > 0 {
		json, _ := json.Marshal(user)
		connections.SendReponse(writer, http.StatusOK, json)
	} else {
		Message := []byte(`{"Error": "Invalid user id "}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}

func GetAllUser(writer http.ResponseWriter, request *http.Request) {
	var users []models.User
	db := connections.GetConnection()
	defer db.Close()

	result := db.Find(&users)
	if result.RowsAffected > 0 {
		json, _ := json.Marshal(users)
		connections.SendReponse(writer, http.StatusOK, json)
	} else {
		Message := []byte(`{"Error": "User not found"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}

func Logout(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := connections.GetConnection()
	defer db.Close()

	result := db.First(&user, "id = ?", id)
	if user.Token != "" {
		if result.RowsAffected > 0 {
			db.Model(&user).Where("id = ?", id).Update("token", "")

			http.SetCookie(writer,
				&http.Cookie{
					Name:    "jwt",
					Value:   "",
					Expires: time.Now().Add(-time.Hour),
				})
			json, _ := json.Marshal(user)

			connections.SendReponse(writer, http.StatusOK, json)
		} else {
			Message := []byte(`{"Error": "User not found"}`)
			connections.ErrorMsg(writer, http.StatusBadRequest, Message)
		}
	} else {
		Message := []byte(`{"Error": "Already Logged out"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}
