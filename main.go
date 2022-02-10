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
			configureRouter,
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

func createServer(engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
}

func configureRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	return router
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
	e *gin.Engine
}

func newHandler(e *gin.Engine) *Handler {
	h := &Handler{e: e}

	h.registerRoutes(h.e.Group("/"))
	return h
}

func (h *Handler) registerRoutes(router *gin.RouterGroup) {
	router.GET("hello", h.main)
}

func (h *Handler) main(c *gin.Context) {
	fmt.Println("In /hello *")
	c.String(http.StatusOK, "Hello World")
}
