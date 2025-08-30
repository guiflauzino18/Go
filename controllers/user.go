package controllers

import (
	"encoding/json"
	"go-project/database"
	"go-project/model"
	"go-project/repository"
	"go-project/response"
	"io"
	"net/http"
)

// Login godoc
// @Summary Create an user
// @Description Register a new user in Database
// @Tags User
// @Success 200 {object} model.User
// @Failure 403 Action Not permited
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	// Recupera dados do Body da REquest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Unmarshal do dados do body para o struct user
	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Prepara os dados do usuario para ser salvo no banco
	if err := user.Prepare("register"); err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Conecta ao banco
	db, err := database.Connect()
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepo(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)

}
