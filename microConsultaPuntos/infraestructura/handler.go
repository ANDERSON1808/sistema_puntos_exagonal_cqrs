package infraestructura

import (
	"github.com/gorilla/mux"
	"microConsultaPuntos/dominio/servicios"
	"net/http"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/api/v1/puntos/consulta", servicios.SearchUserHandler).Methods(http.MethodGet)
	return
}
