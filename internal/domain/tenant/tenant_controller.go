package tenant

import (
	"mikrotikapp/internal/domain/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TenantController struct {
	service *services.TenantService
}

func NewTenantController(service *services.TenantService) *TenantController {
	return &TenantController{service: service}
}

type CreateTenantRequest struct {
	Name string `json:"name"`
}

func (h *TenantController) Create(c echo.Context) error {
	var req CreateTenantRequest
	if err := c.Bind(&req); err != nil {
		return  c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}
	role := c.Get("role").(string)

	tenant, err := h.service.Create(req.Name, role)

	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		return c.JSON(status, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, tenant)
}

func (h *TenantController) List(c echo.Context) error {
	role := c.Get("role").(string)

	tenants, err := h.service.List(role)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, tenants)
}

func (h *TenantController) Delete(c echo.Context) error {
	roleVal := c.Get("role")
	if roleVal == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	role, ok := roleVal.(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid role"})
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	if err := h.service.Delete(id, role); err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "tenant deleted"})
}
