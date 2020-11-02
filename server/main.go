package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tyuhara/http-spanner/config"
	"github.com/tyuhara/http-spanner/http"
	"github.com/tyuhara/http-spanner/log"
)

func main() {
	ctx := context.Background()
	conf, err := config.New(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to read environment variables: %s\n", err)
		return
	}

	logger, err := log.New(conf)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup logger: %s\n", err)
		return
	}
	defer func() {
		_ = logger.Sync()
	}()
	sugar := logger.Sugar()

	server, _ := http.New(conf, sugar)

	http.Handler(server)
}
