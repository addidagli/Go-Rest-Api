package routes

import (
	"github.com/gorilla/mux"
	"main/controllers"
)

func SetUserRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/register", controllers.Register).Methods("POST")
	subRoute.HandleFunc("/login", controllers.Login).Methods("POST")
	subRoute.HandleFunc("/getUser/{id}", controllers.GetUser).Methods("GET")
	subRoute.HandleFunc("/getAllUser", controllers.GetAllUser).Methods("GET")
	subRoute.HandleFunc("/logout/{id}", controllers.Logout).Methods("POST")
	subRoute.HandleFunc("/addWallet", controllers.AddWallet).Methods("POST")
	subRoute.HandleFunc("/wallets/{id}/balance", controllers.GetBalance).Methods("GET")
	subRoute.HandleFunc("/wallets/{id}/credit", controllers.AddCredit).Methods("POST")
	subRoute.HandleFunc("/wallets/{id}/debit", controllers.Debit).Methods("POST")

}
