package api

import (
	"net/http"
	"time"

	"github.com/duohedron/orders/internal/orders"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, svc *orders.Service) {
	e.GET("/healthz", func(c echo.Context) error { return c.String(200, "ok") })

	e.POST("/orders", func(c echo.Context) error {
		req := struct {
			Item string `json:"item"`
		}{}
		if err := c.Bind(&req); err != nil {
			return err
		}
		o := &orders.Order{
			ID:        uuid.New(),
			Item:      req.Item,
			CreatedAt: time.Now(),
		}
		if err := svc.Create(c.Request().Context(), o); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, o)
	})

	e.GET("/orders/:id", func(c echo.Context) error {
		id, _ := uuid.Parse(c.Param("id"))
		o, err := svc.GetByID(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusOK, o)
	})
}
