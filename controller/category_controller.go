package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	GetAllController(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetOneController(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	CreateController(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	UpdateController(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	DeleteController(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
}
