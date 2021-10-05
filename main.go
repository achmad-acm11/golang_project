package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Getenv("PORT")
	db := app.NewDB()
	validate := validator.New()

	// Repository
	categoryRepository := repository.NewCategoryRepository()
	// Service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//Controller
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		// Addr:    "localhost:3000",
		Addr:    ":" + port,
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
