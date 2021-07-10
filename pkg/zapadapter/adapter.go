package zapadapter

import (
	"context"
	"encoding/json"
	"fmt"

	"fake-hash.ports-adapters/pkg/log"
	"fake-hash.ports-adapters/pkg/toolkit"
	"go.uber.org/zap"
)

// Config armazena todas as variáveis de configuração do log.
type Config struct {
	LogLevel string
	Version  string
}

// Adapter é uma implementação de log.Logger baseada na dependência uber-go/zap.
type Adapter struct {
	logger         *zap.Logger
	suggaredLogger *zap.SugaredLogger
}

// New instancia um novo Adapter, mas também nos ajuda a validar se ele implementa
// log.Logger corretamente.
func New(c Config) (log.Logger, error) {
	rawConfig := `{
	"level": "%s",
	"enconding": "json",
	"outputPaths": ["stdout"],
	"errorOutputPaths": ["stderr"],
	"initialFields": {"version": "%s"},
	"encoderConfig": {
		"messageKey": "message",
		"levelKey": "level",
		"levelEncoder": "lowercase"
	}
}`
	rawStringConfig := fmt.Sprintf(rawConfig, c.LogLevel, c.Version)
	rawJSONConfig := []byte(rawStringConfig)

	var config zap.Config
	err := json.Unmarshal(rawJSONConfig, &config)
	if err != nil {
		return nil, err
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	sugar := logger.Sugar()
	return Adapter{
		logger:         logger,
		suggaredLogger: sugar,
	}, nil
}

func (a Adapter) Debug(ctx context.Context, args ...interface{}) {
	a.withCorrelationID(ctx).Debug(args...)
}

func (a Adapter) Debugf(ctx context.Context, format string, args ...interface{}) {
	a.withCorrelationID(ctx).Debugf(format, args...)
}

func (a Adapter) Error(ctx context.Context, args ...interface{}) {
	a.withCorrelationID(ctx).Error(args...)
}

func (a Adapter) Errorf(ctx context.Context, format string, args ...interface{}) {
	a.withCorrelationID(ctx).Errorf(format, args...)
}

func (a Adapter) Info(ctx context.Context, args ...interface{}) {
	a.withCorrelationID(ctx).Info(args...)
}

func (a Adapter) Infof(ctx context.Context, format string, args ...interface{}) {
	a.withCorrelationID(ctx).Infof(format, args...)
}

func (a Adapter) Panic(ctx context.Context, args ...interface{}) {
	a.withCorrelationID(ctx).Panic(args...)
}

func (a Adapter) Panicf(ctx context.Context, format string, args ...interface{}) {
	a.withCorrelationID(ctx).Panicf(format, args...)
}

func (a Adapter) Stop(ctx context.Context) error {
	return a.logger.Sync()
}

func (a Adapter) StopError() error {
	return context.Canceled
}

func (a Adapter) withCorrelationID(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		ctx = context.Background()
	}

	correlationIDValue := ctx.Value(toolkit.CorrelationKey)
	correlationID := fmt.Sprintf("%v", correlationIDValue)
	return a.suggaredLogger.With(zap.String("correlationID", correlationID))
}
