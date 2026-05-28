package handler

import (
	"errors"
	"net/http"
)

func ExtractId(r *http.Request) (int, error) {

	idContext, ok := r.Context().Value("UserID").(int)

	if !ok {
		return -1, errors.New("error en el id de la sesion")
	}

	return idContext, nil
}