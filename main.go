package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			configureGinEngine,
			createServer),
		fx.Invoke(
			registerHooks,
			newHandler),
	)

	if err := app.Err(); err == nil {
		app.Run()
	} else {
		panic(err)
	}
}

func createServer(h http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
}

func configureGinEngine() (http.Handler, gin.IRouter) {
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

type Handler struct {
	router gin.IRouter
}

func newHandler(r gin.IRouter) *Handler {
	h := &Handler{router: r}

	h.registerRoutes(h.router.Group("/"))
	return h
}

func (h *Handler) registerRoutes(router gin.IRoutes) {
	router.GET("hello", h.sayHello)
}

func (h *Handler) sayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
