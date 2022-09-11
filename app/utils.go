package main

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, v any) {
	res, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "InternalServerError", 500)
		return
	}
	w.Write(res)
}
