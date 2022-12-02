package helper

import (
	"encoding/json"
	"net/http"
)

func JsonToMarshal(data interface{}) string {
	result, err := json.Marshal(data)
	PanicIfError(err)
	return string(result)
}

func JsonToUnmarshal(data string, result interface{}) {
	err := json.Unmarshal([]byte(data), result)
	PanicIfError(err)
}

func ReadFromRequestBody(request *http.Request, data interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, responseData interface{}) {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(responseData)
	PanicIfError(err)
}
