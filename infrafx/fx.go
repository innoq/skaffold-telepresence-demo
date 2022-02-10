package infrafx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	
	"go.uber.org/fx"

	"hello-go/handler"
)

var Module = fx.Options(
	fx.Provide(
		createGinEngine,
		createServer),
	fx.Invoke(registerHooks),
	handler.Module,
)

func createServer(h http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
}

func createGinEngine() (http.Handler, gin.IRouter) {
	router := gin.New()
	router.Use(gin.Recovery())
	return router, router
}

func registerHooks(lifecycle fx.Lifecycle, server *http.Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					// service connections
					if err := server.ListenAndServe(); err != http.ErrServerClosed {
						fmt.Println("Server terminated unexpected")
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Tearing Down!")
				server.Shutdown(ctx)
				return nil
			},
		},
	)
}
