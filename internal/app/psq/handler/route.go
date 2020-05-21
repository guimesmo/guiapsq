package handlers

import "github.com/labstack/echo/v4"

func Route(echo *echo.Echo, handler ClinetHandler) {
	v1 := echo.Group("/api/v1")
	{
		v1.POST("/psq/create", handler.PsqCreate)
	}
}
