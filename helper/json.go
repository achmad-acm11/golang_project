package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, result interface{}) {
	err := json.NewDecoder(req.Body).Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(wr http.ResponseWriter, response interface{}) {
	wr.Header().Add("content-type", "application/json")
	err := json.NewEncoder(wr).Encode(response)
	PanicIfError(err)
}
