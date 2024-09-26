package routes

import (
	"net/http"

	"github.com/labstrack/echo/v4"
)



func NewApiRoutes(e *echo.Echo) {
      group := e.Group("/api")

    group.GET("/health-check", func(ctx *echo.Context) error {
    return ctx.JSON(http.StatusOK, struct{
      OK bool `json:"ok"`
    }{
        OK: true,
      })
  })
}
