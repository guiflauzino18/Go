package routes

import (
	"go-project/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Funcao   func(http.ResponseWriter, *http.Request)
	WithAuth bool
}

func ConfigureRouters(mux *mux.Router) *mux.Router {
	var r []Route

	r = append(r, routeLogin)
	r = append(r, checkRoute)
	r = append(r, userRoute...)

	for _, route := range r {

		//If route with auth call the middlewares
		if route.WithAuth {
			mux.HandleFunc(route.URI, middlewares.AuthValidate(route.Funcao)).Methods(route.Method)

		} else {
			mux.HandleFunc(route.URI, route.Funcao).Methods(route.Method)
		}
	}

	return mux
}
