package controllers

import (
	"encoding/json"
	"go-project/database"
	"go-project/model"
	"go-project/repository"
	"go-project/response"
	"go-project/security"
	"io"
	"net/http"
	"strconv"
)

// Login godoc
// @Summary Login
// @Description Faz login na aplicação
// @Tags Login
// @Success 200 {object} model.User
// @Failure 403 Erro no login
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {

	//Recupera dados
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.ErrorJSON(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//Unmarshal dos dados recebidos para o model user
	var user model.User
	if erro := json.Unmarshal(body, &user); erro != nil {
		response.ErrorJSON(w, http.StatusBadRequest, erro)
		return
	}

	//Conecta no Banco
	db, erro := database.Connect()
	if erro != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	//Busca usuario pelo email
	repo := repository.NewUserRepo(db)
	userDB, err := repo.FindByMail(user.Mail)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Compara senha passada com hash da senha do banco
	if err := security.CompareHashPassword(userDB.Password, user.Password); err != nil {
		response.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	// Se autenticação for válida gera token e o retorna
	token, err := security.TokenGenerate(userDB.ID)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	userIdString := strconv.FormatUint(userDB.ID, 10)

	response.JSON(w, http.StatusOK, model.LoginResponse{ID: userIdString, Token: token})

}
