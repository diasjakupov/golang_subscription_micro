package delieveries

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	echo *echo.Echo
}

func NewRouter(handler *Handler) *Router {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	e.POST("/api/v1/subscriptions", handler.CreateSubscriptionHandler)
	e.GET("/api/v1/subscriptions", handler.CheckSubscriptionHandler)
	e.POST("/api/v1/subscriptions/cancel", handler.CancelSubscriptionHandler)
	e.POST("/api/v1/subscriptions/renew", handler.RenewSubscriptionHandler)
	e.GET("/api/v1/plans", handler.GetPlansHandler)
	return &Router{echo: e}
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
