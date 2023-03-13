package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	// decoder stream request json to golang type response
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	// add header writer for tell content as json
	writer.Header().Add("Content-Type", "application/json")

	// encoder data from type golang to stream json
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicError(err)
}
