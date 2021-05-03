package controllers

import (
	"net/http"

	"saunaApi/api/responses"
)

func (server *Server) Top(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Sauna API")
}
