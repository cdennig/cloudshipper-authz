package server

import (
	"context"
	"fmt"
	"time"

	"github.com/cdennig/cloudshipper-authz/internal/config"
	"github.com/kataras/iris/v12"
)

// Start to start the server
func Start(cfg *config.Config, app *iris.Application) error {

	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts.
		app.Shutdown(ctx)
		close(idleConnsClosed)
	})

	// [...]
	app.Listen(fmt.Sprintf("0.0.0.0:%s", cfg.Port), iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnsClosed

	return nil
}
