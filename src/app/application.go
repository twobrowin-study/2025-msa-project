package app

import (
	"context"
	"net/http"

	"otus.ru/tbw/msa-25/src/deps"

	"os"
	"os/signal"
	"syscall"
)

func Run(deps *deps.Deps) {
	runChan := make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		deps.Config.Server.Timeout.Server,
	)
	defer cancel()

	router := newRouter(deps)
	handler := deps.Log.Middleware(router)

	server := &http.Server{
		Addr:         deps.Config.Server.Host + ":" + deps.Config.Server.Port,
		Handler:      handler,
		ReadTimeout:  deps.Config.Server.Timeout.Read,
		WriteTimeout: deps.Config.Server.Timeout.Write,
		IdleTimeout:  deps.Config.Server.Timeout.Idle,
	}

	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	deps.Log.Infof("Starting server at %s:%s", deps.Config.Server.Host, deps.Config.Server.Port)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				deps.Log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	interrupt := <-runChan
	deps.Log.Infof("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		deps.Log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
