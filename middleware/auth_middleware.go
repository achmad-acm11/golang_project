package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if "RAHASIA" == req.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(wr, req)
	} else {
		wr.Header().Set("content-type", "application/json")
		wr.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(wr, webResponse)
	}
}
