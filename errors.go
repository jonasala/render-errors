package rendererr

import (
	"net/http"

	"github.com/go-chi/render"
)

//ErrResponse é a representação de um erro no sistema
type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        string `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

//Render processa a renderização da resposta
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

//ErrInvalidRequest monta um erro demonstrando uma requisição inválida
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     http.StatusText(http.StatusBadRequest),
		ErrorText:      err.Error(),
	}
}

//ErrRender monta um erro demonstrando problemas em montar a resposta
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     http.StatusText(http.StatusUnprocessableEntity),
		ErrorText:      err.Error(),
	}
}

//ErrUnauthorized monta um erro demonstrando a falta de uma JWT Token
var ErrUnauthorized = &ErrResponse{
	HTTPStatusCode: http.StatusUnauthorized,
	StatusText:     http.StatusText(http.StatusUnauthorized),
}

//ErrForbidden monta um erro demonstrando falta de permissão para executar a requisição
var ErrForbidden = &ErrResponse{
	HTTPStatusCode: http.StatusForbidden,
	StatusText:     http.StatusText(http.StatusForbidden),
}

//ErrNotFound é o bom e velho 404
var ErrNotFound = &ErrResponse{HTTPStatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}
