package servicios

import (
	"encoding/json"
	"microConsultaPuntos/dominio/repositorio"
	"net/http"
	"strconv"
)

func SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "query is required", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(query)
	usuarioId, err := repositorio.BuscarUsuarioId(id)
	if err != nil {
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usuarioId)
	if err != nil {
		return
	}
}
