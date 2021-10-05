package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.GetAllController)
	router.POST("/api/categories", categoryController.CreateController)
	router.GET("/api/categories/:categoryId", categoryController.GetOneController)
	router.PUT("/api/categories/:categoryId", categoryController.UpdateController)
	router.DELETE("/api/categories/:categoryId", categoryController.DeleteController)

	router.PanicHandler = exception.ErrorHandler

	return router
}
