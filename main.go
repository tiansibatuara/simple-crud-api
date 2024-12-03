package main

import (
	"fmt"
	"net/http"
	"simple-crud-api/config"
	"simple-crud-api/controller"
	"simple-crud-api/helper"
	"simple-crud-api/repository"
	"simple-crud-api/router"
	"simple-crud-api/service"
)

func main() {
	fmt.Printf("Start main")

	// database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewSongRepository(db)

	// service
	bookService := service.NewSongRepositoryImpl(bookRepository)

	// controller
	bookController := controller.NewSongController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
