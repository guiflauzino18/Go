package main

import (
	"fmt"
	"go-project/config"
	"go-project/controllers"
	"go-project/database"
	"go-project/repository"
	"go-project/routes"
	"go-project/service"
	"log"

	_ "go-project/docs" // import dos docs gerados

	"github.com/gin-gonic/gin"
)

// @title Devbook
// @version 1.0.0
// @contact.name Guilherme Flauzino
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func main() {
	//Carrega as variáveis de ambientes
	config.LoadEnv()

	// Router Gin
	r := gin.Default()

	// Routes Configure
	setupDependecies(r)

	// Inicia servidor web
	r.Run(fmt.Sprintf(":%d", config.Port))

}

func setupDependecies(r *gin.Engine) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	routes.LoginRouters(r)

	routes.SwaggerRouters(r)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(*userRepo)
	userController := controllers.NewUserController(userService)
	routes.UserRouters(r, *userController)
}
