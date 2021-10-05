package exception

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(wr http.ResponseWriter, req *http.Request, err interface{}) {

	if notFoundError(wr, req, err) {
		return
	}
	if validationError(wr, req, err) {
		return
	}
	internalServerError(wr, req, err)
}
func validationError(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		wr.Header().Add("content-type", "application/json")
		wr.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(wr, webResponse)

		return true
	} else {
		return false
	}
}
func notFoundError(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		wr.Header().Add("content-type", "application/json")
		wr.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(wr, webResponse)

		return true
	} else {
		return false
	}
}
func internalServerError(wr http.ResponseWriter, req *http.Request, err interface{}) {
	wr.Header().Add("content-type", "application/json")
	wr.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(wr, webResponse)
}
