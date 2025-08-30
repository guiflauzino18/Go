package routes

import (
	"go-project/controllers"
	"net/http"
)

var userRoute = []Route{
	{
		URI:      "/register",
		Method:   http.MethodPost,
		Funcao:   controllers.Register,
		WithAuth: false,
	},
}
