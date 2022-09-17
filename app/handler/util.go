package handler

import (
	"encoding/json"
	appError "github.com/tomoyasuzuki/sample-api/app/error"
	"net/http"
)

func Response(w http.ResponseWriter, v any, err *appError.AppError) {
	if err != nil {
		http.Error(w, "", err.GetHttpCode())
	} else {
		if v != nil {
			res, err := json.Marshal(v)
			if _, err = w.Write(res); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
