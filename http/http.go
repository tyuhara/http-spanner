package http

import (
	"github.com/tyuhara/http-spanner/config"
	"go.uber.org/zap"
)

type HttpServer struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func New(conf *config.Config, logger *zap.SugaredLogger) (*HttpServer, error) {
	return &HttpServer{
		config: conf,
		logger: logger,
	}, nil
}
