package handlers

import (
	"net/http"

	"github.com/guimesmo/guiapsq/internal/psq/models"
	"github.com/guimesmo/guiapsq/internal/psq/services"
	"github.com/labstack/echo/v4"
)

// Handler handles the requests for PSQ area
type Handler struct {
	service services.PsqService
}

func NewHandler(service services.PsqService) Handler {
	return Handler{service}
}

// PSQ Creation endpoint
func (h *Handler) PsqCreate(c echo.Context) error {
	psq := new(models.Psq)

	c.Bind(psq)
	h.service.Insert(c.Request().Context(), *psq)

	return c.JSON(http.StatusOK, psq)
}
