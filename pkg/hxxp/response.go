package hxxp

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, obj interface{})  {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, _ := json.Marshal(obj)
	w.Write(jsonBytes)
}
