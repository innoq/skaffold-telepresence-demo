package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler(r gin.IRouter) *Handler {
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
	req, err := http.NewRequest("GET", "https://kubernetes.default.svc", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Response status code: %s", res.Status)
	}

	c.String(http.StatusOK, "Hello World")
}
