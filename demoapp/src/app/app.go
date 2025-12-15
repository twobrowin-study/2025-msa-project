package app

import (
	"context"
	"net/http"

	"twb-otus-25/demoapp/src/config"

	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)


func Run(log *logrus.Logger, config *config.Config) {
	var runChan = make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(
        context.Background(),
        config.Server.Timeout,
    )
    defer cancel()

	server := &http.Server{
        Addr:    "0.0.0.0:" + config.Server.Port,
        Handler: newRouter(log),
    }

	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	log.Infof("Starting server at port %s", config.Server.Port)
	go func() {
        if err := server.ListenAndServe(); err != nil {
            if err != http.ErrServerClosed {
                log.Fatalf("Server failed to start due to err: %v", err)
            }
        }
    }()

    interrupt := <-runChan
    log.Printf("Server is shutting down due to %+v\n", interrupt)
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
    }
}
