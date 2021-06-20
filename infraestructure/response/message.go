package response

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/jpastorm/dialogflowbot/model"
)

// errorTypes
const (
	// Failure is a special code for send the custom error message, and API message from the logic and not from responses map
	Failure model.StatusCode = "failure"

	UnexpectedError             model.StatusCode = "unexpected_error"
	UniqueError                 model.StatusCode = "unique_error"
	ForeignKeyError             model.StatusCode = "foreign_key_error"
	NotNullError                model.StatusCode = "not_null_error"
	InvalidParameter            model.StatusCode = "invalid_parameter"
	InvalidPagination           model.StatusCode = "invalid_pagination"
	BindFailed                  model.StatusCode = "bind_failed"
	ValidationFailed            model.StatusCode = "validation_failed"
	RecordCreated               model.StatusCode = "record_created"
	RecordUpdated               model.StatusCode = "record_updated"
	RecordDeleted               model.StatusCode = "record_deleted"
	Ok                          model.StatusCode = "ok"
	RecordNotFound              model.StatusCode = "record_not_found"
	InvalidPaginationParameters model.StatusCode = "invalid_pagination_parameter"
)

type data struct {
	Status  int
	Message string
}

var responses = map[model.StatusCode]data{
	UnexpectedError: {
		Status:  http.StatusInternalServerError,
		Message: "¡Upps! algo inesperado ocurrió",
	},
	UniqueError: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! nos enviaste un registro duplicado, revisa la documentación del paquete",
	},
	ForeignKeyError: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! el id de un modelo relacionado en el payload no existe o no fue enviado",
	},
	NotNullError: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! nos enviaste datos nulos, revisa la documentación del paquete",
	},
	InvalidParameter: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! el valor del parámetro que has enviado no es valido. Por favor revisa el formato y tipo",
	},
	InvalidPagination: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! los parámetros de paginación no son validos",
	},
	BindFailed: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! el payload que enviaste no es valido, verifica la documentación del paquete",
	},
	ValidationFailed: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! la información que nos enviasté esta incompleta, verifica la documentación del paquete",
	},
	RecordCreated: {
		Status:  http.StatusCreated,
		Message: "Registro creadó",
	},
	RecordUpdated: {
		Status:  http.StatusOK,
		Message: "Registro actualizadó",
	},
	RecordDeleted: {
		Status:  http.StatusOK,
		Message: "Registro eliminadó",
	},
	Ok: {
		Status:  http.StatusOK,
		Message: "",
	},
	RecordNotFound: {
		Status:  http.StatusOK,
		Message: "No se encontraron registros",
	},
	InvalidPaginationParameters: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! los parámetros de paginacián no son válidos",
	},
}

// MessageResponse contains the response message
type MessageResponse struct {
	Data     interface{} `json:"data,omitempty"`
	Errors   []*Response `json:"errors,omitempty"`
	Messages []*Response `json:"messages,omitempty"`
}

type Response struct {
	Code    model.StatusCode `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
}

func Successfull(code model.StatusCode, data interface{}) (status int, resp MessageResponse) {
	status = http.StatusOK
	res := &Response{code, "Ok"}

	if e, ok := responses[code]; ok {
		status = e.Status
		res.Message = e.Message
	}
	resp.Messages = append(resp.Messages, res)
	resp.Data = data

	return
}

// Failed returns an Error
func Failed(who string, co model.StatusCode, err error) *model.Error {
	fun, _, line, _ := runtime.Caller(1)

	errData := model.NewError()

	errData.SetCode(co)
	errData.SetError(err)
	errData.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	errData.SetWho(who)

	return errData
}
