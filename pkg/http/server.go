package http

import (
	"context"

	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/stoppage"
)

// Server define o que é necessário para criar um servidor HTTP.
type Server interface {
	stoppage.Stopper
	Listen(ctx context.Context, router *Router) error
}

// Router contém os handlers dos endpoints que serão disponibilizados pelo servidor.
type Router struct {
	Handlers []Handler
}

func (r *Router) Post(url string, handler HandleFunc) {
	r.Handlers = createHandler(r.Handlers, url, "POST", handler)
}

// Handler define como um handler de um endpoint deve ser definido.
type Handler struct {
	Path    string
	Method  string
	Handler HandleFunc
}

// HandleFunc define como a função do handler de um endpoint deve ser criado.
type HandleFunc = func(ctx context.Context, req Request, res Response)

// Request define todos os métodos necessários para gerenciar uma requisição HTTP.
type Request interface {
	BodyBytes() []byte
	Url() string
	Header(name string) string
}

// Response define todos os métodos necessários para gerenciar uma resposta HTTP.
type Response interface {
	Write(statusCode int, body []byte) error
	WriteCreated(body []byte) error
	WriteBadRequest(body []byte) error
	WriteInternalError(body []byte) error
	JsonResponse()
}

func createHandler(handlers []Handler, url string, method string, handler func(context.Context, Request, Response)) []Handler {
	return append(handlers, Handler{
		Path:    url,
		Method:  method,
		Handler: handler,
	})
}
