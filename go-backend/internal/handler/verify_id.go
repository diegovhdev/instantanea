package handler

import (
	"errors"
	"net/http"
	"strconv"
)

func VerifyId(r *http.Request) (int, int, error) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		return -1, http.StatusBadRequest, errors.New("el argumento del path tiene que ser un numero")
	}

	idContext, ok := r.Context().Value("UserID").(int)

	if !ok {
		return -1, http.StatusInternalServerError, errors.New("error en el id de la sesion")
	}

	if id != idContext {
		return -1, http.StatusForbidden, errors.New("el id pedido y el que se tiene no coinciden")
	}
	return id, 0, nil
}