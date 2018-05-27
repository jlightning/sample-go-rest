package handlers

import (
	"net/http"
	"encoding/json"
)

func wrapFunc(f func(r *http.Request) (interface{}, error)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, r *http.Request) {
		result, err := f(r)
		if err != nil {
			if err1, ok := err.(HttpError); ok {
				http.Error(writer, err1.Error(), err1.httpCode)
			} else {
				http.Error(writer, err.Error(), 500)
			}
			return
		}

		if result != nil {
			jsonResponse, err := json.Marshal(result)
			if err != nil {
				http.Error(writer, err.Error(), 500)
				return
			}

			writer.Write(jsonResponse)
		}
	}
}
