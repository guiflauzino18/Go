package route

import (
	"go-project/route/routes"

	"github.com/gorilla/mux"
)

func CreateRouters() *mux.Router {
	mux := mux.NewRouter()
	return routes.ConfigureRouters(mux)
}

/*
Exemplo simples de criação de Rotas para interpretação
*    r := mux.NewRouter()

    // Rotas da API
    r.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
    r.HandleFunc("/api/v1/users/{id}", GetUser).Methods("GET")
    r.HandleFunc("/api/v1/users", CreateUser).Methods("POST")

    // Rota do Swagger
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    http.ListenAndServe(":8080", r)
*/
