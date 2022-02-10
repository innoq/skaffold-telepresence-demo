package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newHandler(r gin.IRouter) *Handler {
	h := &Handler{router: r}

	h.registerRoutes(h.router.Group("/"))
	return h
}

type Handler struct {
	router gin.IRouter
}

func (h *Handler) registerRoutes(router gin.IRoutes) {
	router.GET("hello", h.sayHello)
}

func (h *Handler) sayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
