package stoppage

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/log"
)

// Shutdown ajuda a gerenciar múltiplos Stopper ao encerrarmos uma aplicação.
type Shutdown struct {
	Logger   log.Logger
	stoppers []Stopper
}

func New(logger log.Logger) Shutdown {
	return Shutdown{
		Logger:   logger,
		stoppers: nil,
	}
}

// Add adicionará um Stopper dado ao Shutdown.
func (s *Shutdown) Add(stopper Stopper) {
	s.stoppers = append(s.stoppers, stopper)
}

// GracefulSignal aceita um contexto para retornar um novo que observa sinais
// emitidos pelo OS para encerrar a aplicação.
func (s Shutdown) GracefulSignal(ctx context.Context) context.Context {
	ctx, done := context.WithCancel(ctx)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer done()

		signalCalled := <-quit
		s.Logger.Infof(ctx, "Starting shutdown by signal: ", signalCalled.String())
		signal.Stop(quit)
		close(quit)

		s.Graceful(ctx)
	}()

	return ctx
}

// Graceful executará o método Stop de todos os seus Stopper e irá esperá-los
// finalizar por no máximo 30 segundos.
func (s Shutdown) Graceful(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	for _, stopper := range s.stoppers {
		err := stopper.Stop(ctx)
		if err == nil {
			return
		}

		stopErr := stopper.StopError()
		if stopErr != nil {
			s.Logger.Errorf(ctx, stopErr.Error())
			return
		}

		s.Logger.Error(ctx, err.Error())
	}
}
