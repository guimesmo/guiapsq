package handlers

import (
	"net/http"

	"github.com/guimesmo/guiapsq/internal/psq/services"
	"github.com/guimesmo/guiapsq/repository/mongo"
	"github.com/labstack/echo/v4"
)

// Handler handles the requests for PSQ area
type Handler struct {
	DB *mongo.Connection
}

// PSQ Creation endpoint
func (h *Handler) PsqCreate(c echo.Context) error {
	psq, err := services.PsqCreate(echo.Context, h.DB)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, psq)
}
