package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"main/connections"
	"main/models"
	"net/http"
)

func AddWallet(writer http.ResponseWriter, request *http.Request) {
	wallet := models.Wallet{}
	user := models.User{}

	db := connections.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&wallet)

	result := db.First(&wallet, "id = ?", wallet.ID)
	result2 := db.First(&user, "id = ?", wallet.UserId)
	if result.RowsAffected > 0 {
		Message := []byte(`{"Error": "Wallet is already exist"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	} else {
		if wallet.Balance < 0 || wallet.Debit < 0 || wallet.Credit < 0 {
			Message := []byte(`{"Error": "Balance,debit or credit can not be less than 0"}`)
			connections.ErrorMsg(writer, http.StatusBadRequest, Message)
		}
		if result2.RowsAffected > 0 {
			wallet.User.ID = user.ID
			wallet.User.Email = user.Email
			wallet.User.Password = user.Password
			wallet.User.Name = user.Name
			wallet.User.Token = user.Token

			if error != nil {
				log.Fatal(error)
				connections.SendError(writer, http.StatusBadRequest)
				return
			}

			error = db.Save(&wallet).Error

			if error != nil {
				log.Fatal(error)
				connections.SendError(writer, http.StatusInternalServerError)
				return
			}

			json, _ := json.Marshal(wallet)

			connections.SendReponse(writer, http.StatusCreated, json)
		} else {
			Message := []byte(`{"Error": "Invalid user id "}`)
			connections.ErrorMsg(writer, http.StatusBadRequest, Message)
		}
	}

}

func GetBalance(writer http.ResponseWriter, request *http.Request) {
	wallet := models.Wallet{}
	user := models.User{}

	id := mux.Vars(request)["id"]

	println("id: ", id)

	db := connections.GetConnection()
	defer db.Close()

	db.First(&wallet, "id = ?", id)
	result2 := db.First(&user, "id = ?", wallet.UserId)

	if result2.RowsAffected > 0 {
		wallet.User.ID = user.ID
		wallet.User.Email = user.Email
		wallet.User.Password = user.Password
		wallet.User.Name = user.Name
		wallet.User.Token = user.Token
	}

	if wallet.ID > 0 {
		if user.Token != "" {
			json, _ := json.Marshal(wallet)
			connections.SendReponse(writer, http.StatusOK, json)
		} else {
			Message := []byte(`{"Error": "Please login first "}`)
			connections.ErrorMsg(writer, http.StatusUnauthorized, Message)
		}
	} else {
		Message := []byte(`{"Error": "Invalid wallet id "}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}

func AddCredit(writer http.ResponseWriter, request *http.Request) {
	wallet := models.Wallet{}
	user := models.User{}

	id := mux.Vars(request)["id"]

	println("id: ", id)

	db := connections.GetConnection()
	defer db.Close()

	result := db.First(&wallet, "id = ?", id)
	result2 := db.First(&user, "id = ?", wallet.UserId)

	if result2.RowsAffected > 0 {
		wallet.User.ID = user.ID
		wallet.User.Email = user.Email
		wallet.User.Password = user.Password
		wallet.User.Name = user.Name
		wallet.User.Token = user.Token
	}

	error := json.NewDecoder(request.Body).Decode(&wallet)
	if error != nil {
		log.Fatal(error)
		connections.SendError(writer, http.StatusBadRequest)
		return
	}

	if result.RowsAffected > 0 {
		if user.Token != "" {
			if wallet.Credit < 0 {
				Message := []byte(`{"Error": "Credit can not be less than 0'"}`)
				connections.ErrorMsg(writer, http.StatusBadRequest, Message)
			} else {
				wallet.Balance += wallet.Credit
				error = db.Model(&wallet).Where("id = ?", id).Updates(map[string]interface{}{"credit": wallet.Credit, "balance": wallet.Balance}).Error

				if error != nil {
					log.Fatal(error)
					connections.SendError(writer, http.StatusInternalServerError)
					return
				}

				json, _ := json.Marshal(wallet)

				connections.SendReponse(writer, http.StatusOK, json)
			}
		} else {
			Message := []byte(`{"Error": "Please login first "}`)
			connections.ErrorMsg(writer, http.StatusUnauthorized, Message)
		}
	} else {
		Message := []byte(`{"Error": "Invalid wallet id"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}

func Debit(writer http.ResponseWriter, request *http.Request) {
	wallet := models.Wallet{}
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := connections.GetConnection()
	defer db.Close()

	result := db.First(&wallet, "id = ?", id)
	result2 := db.First(&user, "id = ?", wallet.UserId)

	if result2.RowsAffected > 0 {
		wallet.User.ID = user.ID
		wallet.User.Email = user.Email
		wallet.User.Password = user.Password
		wallet.User.Name = user.Name
		wallet.User.Token = user.Token
	}

	error := json.NewDecoder(request.Body).Decode(&wallet)
	if error != nil {
		log.Fatal(error)
		connections.SendError(writer, http.StatusBadRequest)
		return
	}

	if result.RowsAffected > 0 {
		if user.Token != "" {
			if wallet.Debit < 0 {
				Message := []byte(`{"Error": "Debit can not be less than 0'"}`)
				connections.ErrorMsg(writer, http.StatusBadRequest, Message)
			} else {
				wallet.Balance -= wallet.Debit

				if wallet.Balance < 0 {
					Message := []byte(`{"Error": "Balance is not enough to play'"}`)

					db.Model(&wallet).Where("id = ?", id).Update("debit", wallet.Debit)

					connections.ErrorMsg(writer, http.StatusBadRequest, Message)

				} else {

					error = db.Model(&wallet).Where("id = ?", id).Updates(map[string]interface{}{"debit": wallet.Debit, "balance": wallet.Balance}).Error

					if error != nil {
						log.Fatal(error)
						connections.SendError(writer, http.StatusInternalServerError)
						return
					}

					json, _ := json.Marshal(wallet)

					connections.SendReponse(writer, http.StatusOK, json)
				}

			}
		} else {
			Message := []byte(`{"Error": "Please login first "}`)
			connections.ErrorMsg(writer, http.StatusUnauthorized, Message)
		}
	} else {
		Message := []byte(`{"Error": "Invalid wallet id"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}
