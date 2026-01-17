package auth

import (
	"net/http"
	"fmt"

	"mikrotikapp/internal/domain/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// POST /register
func (h *AuthController) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}

	var tenantUUID *uuid.UUID
	role := req.Role
	if role == "" {
		role = "TENANT_ADMIN"
	}

	// TENANT_ADMIN wajib punya tenant
	if role == "TENANT_ADMIN" {
		if req.TenantID == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "tenant_id is required for TENANT_ADMIN",
			})
		}

		u, err := uuid.Parse(req.TenantID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "invalid tenant_id",
			})
		}
		tenantUUID = &u
	}

	user, err := h.authService.Register(
		req.Username,
		req.Password,
		tenantUUID,
		role,
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id":         user.ID,
		"username":   user.Username,
		"tenant_id":  user.TenantID,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	})
}

// POST /login
func (h *AuthController) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}

	userID, tenantID, token, err := h.authService.Login(
		req.Username,
		req.Password,
	)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":   userID,
		"tenant_id": tenantID,
		"token":     token,
	})
}

// GET /profile (JWT protected)
func (h *AuthController) Profile(c echo.Context) error {
	userIDRaw := c.Get("user_id")
	if userIDRaw == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}


	userID, ok := userIDRaw.(uuid.UUID)
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid user id",
		})
	}
	
	fmt.Println("PROFILE user_id:", userID)

	user, err := h.authService.Profile(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "user not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id":         user.ID,
		"username":   user.Username,
		"tenant_id":  user.TenantID,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	})
}
