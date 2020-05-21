package handlers

import (
	"github.com/labstack/echo/v4"
)

type ClinetHandler struct {
}

func NewClinetHandler() ClinetHandler {
	return ClinetHandler{}
}

func (h ClinetHandler) PsqCreate(c echo.Context) error {
	return nil
}
