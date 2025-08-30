package routes

import (
	"go-project/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Funcao:   controllers.Login,
	WithAuth: false,
}
