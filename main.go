package main

import (
	"fmt"
	"go-project/config"
	"go-project/route"
	"log"
	"net/http"

	_ "go-project/docs" // import dos docs gerados

	httpSwagger "github.com/swaggo/http-swagger"
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

	//Cria as rotas da API
	r := route.CreateRouters()

	// Rota do Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Printf("API on port %d", config.Port)
	//Inicia o servidor web
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
