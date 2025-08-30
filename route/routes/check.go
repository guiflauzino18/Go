package routes

import (
	"go-project/controllers"
	"net/http"
)

var checkRoute = Route{
	URI:      "/check",
	Method:   http.MethodGet,
	Funcao:   controllers.Check,
	WithAuth: false,
}
