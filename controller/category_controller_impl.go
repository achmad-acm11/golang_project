package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{service: service}
}

func (c *CategoryControllerImpl) GetAllController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	categoryResponses := c.service.GetAllCategory(req.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(wr, webResponse)
}
func (c *CategoryControllerImpl) GetOneController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := c.service.GetOneCategory(req.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(wr, webResponse)

}
func (c *CategoryControllerImpl) CreateController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(req, &categoryCreateRequest)

	categoryResponse := c.service.CreateCategory(req.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(wr, webResponse)
}
func (c *CategoryControllerImpl) UpdateController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(req, &categoryUpdateRequest)

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id
	categoryResponse := c.service.UpdateCategory(req.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(wr, webResponse)
}
func (c *CategoryControllerImpl) DeleteController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	c.service.DeleteCategory(req.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(wr, webResponse)
}
