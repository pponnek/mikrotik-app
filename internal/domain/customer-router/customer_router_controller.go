package customerrouterController

import (
	"mikrotikapp/internal/domain/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CustomerRouterController struct {
	service *services.CustomerRouterService
}

func NewCustomerRouterController(s *services.CustomerRouterService) *CustomerRouterController {
	return &CustomerRouterController{service: s}
}

func (h *CustomerRouterController) Create(c echo.Context) error {
	var req struct {
		CoreRouterID uuid.UUID `json:"core_router_id"`
		Name         string    `json:"name"`
		StaticIP     string    `json:"static_ip"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	tenantID := c.Get("tenant_id").(uuid.UUID)

	err := h.service.Create(
		tenantID,
		req.CoreRouterID,
		req.Name,
		req.StaticIP,
	)

	if err != nil {
		switch err {
		case services.ErrCustomerRouterExists:
			return c.JSON(http.StatusConflict, echo.Map{"error": err.Error()})
		default:
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "customer router created"})
}

func (h *CustomerRouterController) List(c echo.Context) error {
	tenantID := c.Get("tenant_id").(uuid.UUID)

	data, err := h.service.List(tenantID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *CustomerRouterController) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	tenantID := c.Get("tenant_id").(uuid.UUID)

	err = h.service.Delete(tenantID, id)
	if err != nil {
		if err == services.ErrCustomerRouterNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "customer router deleted"})
}

func (h *CustomerRouterController) Block(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	tenantID := c.Get("tenant_id").(uuid.UUID)

	err = h.service.Block(tenantID, id)
	if err != nil {
		if err == services.ErrCustomerRouterNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "customer router blocked",
	})
}

func (h *CustomerRouterController) Unblock(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	tenantID := c.Get("tenant_id").(uuid.UUID)

	err = h.service.Unblock(tenantID, id)
	if err != nil {
		if err == services.ErrCustomerRouterNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "customer router unblocked",
	})
}
