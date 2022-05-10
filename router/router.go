package router

import (
	"example/crud/controllers"
	"net/http"
)

const HttpAddr = ":3333"

func Start() http.Handler {
	//Controllers
	userController := controllers.UserController{}
	userController.Init()

	apiMux := http.NewServeMux()
	apiMux.Handle("/users", http.HandlerFunc(userController.CreateAccount))
	apiMux.Handle("/users/list", http.HandlerFunc(userController.ListAccounts))

	return apiMux
}
