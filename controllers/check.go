package controllers

import (
	"go-project/response"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, "OK")
}
