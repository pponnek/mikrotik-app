package corerouterController

import (
	"mikrotikapp/internal/domain/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CoreRouterController struct {
	service *services.CoreRouterService
}

func NewCoreRouterController(s *services.CoreRouterService) *CoreRouterController {
	return &CoreRouterController{service: s}
}

func (h *CoreRouterController) Create(c echo.Context) error {
	var req struct {
		Name     string `json:"name"`
		Host     string `json:"host"`
		ApiPort  int    `json:"api_port"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	tenantID, ok := c.Get("tenant_id").(uuid.UUID)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "tenant not found"})
	}

	err := h.service.Create(
		tenantID,
		req.Name,
		req.Host,
		req.Username,
		req.Password,
		req.ApiPort,
	)

	if err != nil {
		switch err {
		case services.ErrMissingFields:
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		case services.ErrRouterAlreadyExists:
			return c.JSON(http.StatusConflict, echo.Map{"error": err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "core router created"})
}

func (h *CoreRouterController) List(c echo.Context) error {
	tenantID := c.Get("tenant_id").(uuid.UUID)
	routers, err := h.service.List(tenantID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, routers)
}

func (h *CoreRouterController) Delete(c echo.Context) error {
	tenantID := c.Get("tenant_id").(uuid.UUID)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	err = h.service.Delete(tenantID, id)
	if err != nil {
		if err == services.ErrRouterNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "router deleted"})
}