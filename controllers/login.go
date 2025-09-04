package controllers

import (
	"go-project/database"
	"go-project/model"
	"go-project/repository"
	"go-project/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Faz login na aplicação
// @Tags Login
// @Success 200 {object} model.User
// @Failure 403 Erro no login
// @Router /login [post]
func Login(c *gin.Context) {

	//Recupera dados
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Erro ao ler dados"})
		return
	}

	//Conecta no Banco
	db, erro := database.Connect()
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar no banco de dados"})
		return
	}
	defer db.Close()

	//Busca usuario pelo email
	repo := repository.NewUserRepo(db)
	userDB, err := repo.FindByMail(user.Mail)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Erro ao ler dados"})
		return
	}

	// Compara senha passada com hash da senha do banco
	if err := security.CompareHashPassword(userDB.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha inválida"})
		return
	}

	// Se autenticação for válida gera token e o retorna
	token, err := security.TokenGenerate(userDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler dados " + err.Error()})
		return
	}

	userIdString := strconv.FormatUint(userDB.ID, 10)

	c.JSON(http.StatusOK, gin.H{"id": userIdString, "token": token})

}
