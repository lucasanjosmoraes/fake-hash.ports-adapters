package nethttpadapter

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	projecthttp "github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/http"
	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/log"
	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/toolkit"
)

// ServerConfig guarda todas as variáveis de configuração do servidor.
type ServerConfig struct {
	Port string
}

// ServerAdapter implementa http.Server, usando o pacote net/http.
type ServerAdapter struct {
	Logger log.Logger
	server *http.Server
	port   int
}

// NewServer instancia ServerAdapter e nos ajuda a validar se ele implementa http.Server.
func NewServer(ctx context.Context, l log.Logger, c ServerConfig) projecthttp.Server {
	port, err := strconv.Atoi(c.Port)
	if err != nil {
		l.Infof(ctx, "invalid port: %s; using default", c.Port)
	}

	if port == 0 {
		port = 3000
	}

	return ServerAdapter{
		Logger: l,
		server: &http.Server{
			Addr:           fmt.Sprintf(":%d", port),
			ReadTimeout:    time.Second * 5,
			WriteTimeout:   time.Second * 10,
			IdleTimeout:    time.Second * 10,
			MaxHeaderBytes: 1 << 20,
		},
		port: port,
	}
}

// Listen carregará todos os handlers em um router, da biblioteca gorilla, e irá
// executar o ListenAndServe do http.Server para servir os endpoints na porta
// definida.
func (a ServerAdapter) Listen(ctx context.Context, router *projecthttp.Router) error {
	r := mux.NewRouter()

	for _, h := range router.Handlers {
		r.HandleFunc(h.Path, a.handler(ctx, h)).Methods(h.Method)
	}
	a.server.Handler = r

	a.Logger.Infof(ctx, "listening on port %d 🚀🚀🚀", a.port)
	return a.server.ListenAndServe()
}

// Stop irá desabilitar a flag `KeepAlives` e irá encerrar o servidor.
func (a ServerAdapter) Stop(ctx context.Context) error {
	a.server.SetKeepAlivesEnabled(false)

	err := a.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("could not shutdown the server: %s", err)
	}

	return nil
}

// StopError retorna um erro que pode ser diferente de nil caso a rotina
// de shutdown tenha sido disparada devido a um erro.
func (a ServerAdapter) StopError() error {
	return http.ErrServerClosed
}

func (a ServerAdapter) handler(ctx context.Context, h projecthttp.Handler) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		r := Request{NetHTTPRequest: req}
		res := Response{ResponseWriter: rw}
		ctx = toolkit.LoadCorrelationID(ctx, r)

		h.Handler(ctx, r, res)
	}
}
