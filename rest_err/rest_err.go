package rest_err

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status_code"`
	Err   string `json:"error"`
}

func (r *RestError) Error() string{
	return r.Message
}

func NewRestError(message string, status int, err string) *RestError {
	return &RestError{
		Message: message,
		Status:  status,
		Err:   err,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Err:  "bad_request",
	}
}

func NewInternalServerError() *RestError{
	return &RestError{
		Message: "Erro interno do Servidor",
		Status:  http.StatusInternalServerError,
		Err:  "internal_server_error",
	}
}

func NewNotFoundError(message string) *RestError{
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Err:  "not_fund",
	}
}