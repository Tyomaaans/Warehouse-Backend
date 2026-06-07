package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/httpclient"
	"Backend-Warehouse/interface/usecase"
)

type EmployeeHandler struct {
	empUC              *usecase.EmployeeUseCase
	pythonClient       *httpclient.PythonClient
	refreshTokenExpiry time.Duration
}

func NewEmployeeHandler(empUC *usecase.EmployeeUseCase, pythonClient *httpclient.PythonClient, refreshTokenExpiry time.Duration) *EmployeeHandler {
	return &EmployeeHandler{
		empUC:              empUC,
		pythonClient:       pythonClient,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (h *EmployeeHandler) RegisterEmployee(c *gin.Context) {
	var req dto.RegisterEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.empUC.RegisterEmployee(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Register Account Success!",
	})
}

func (h *EmployeeHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
        c.JSON(401, gin.H{"error": "refresh token not found"})
        return
    }
	
	res, err := h.empUC.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
    c.SetCookie(
        "refresh_token",
        res.RefreshToken,
        int(h.refreshTokenExpiry/time.Second),
        "/api/v1/auth",
        "",
        false,
        true,
    )

	c.JSON(200, gin.H{
		"message": "Refresh Token Success!",
		"response": gin.H{
			"access_token": res.AccessToken,
		},
	})
}

func (h *EmployeeHandler) LoginEmployee(c *gin.Context) {
	var req dto.LoginEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	res, token, err := h.empUC.LoginEmployee(c.Request.Context(), req) 
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		token.RefreshToken,
		int(h.refreshTokenExpiry/time.Second),
		"/api/v1/auth",
		"",
		false,
		true,
	)

	c.JSON(200, gin.H{
		"message": "Login Success!",
		"response": res,
	})
}

func (h *EmployeeHandler) LogoutEmployee(c *gin.Context) {
	accessToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
        c.JSON(400, gin.H{"error": "refresh token cookie not found"})
        return
    }

	if err := h.empUC.LogoutEmployee(c.Request.Context(), accessToken, refreshToken); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
    c.SetCookie(
		"refresh_token",
		"", 
		-1, 
		"/api/v1/auth", 
		"", 
		false, 
		true,
	)

    c.JSON(200, gin.H{
		"message": "Logout Success!",
	})
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	employeeID := c.GetString("employeeID")

	res, err := h.empUC.GetEmployee(c.Request.Context(), employeeID)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Employee Success!",
		"response": res,
	})
}

func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	res, err := h.empUC.GetEmployees(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Employees Success!",
		"response": res,
	})
}

func (h *EmployeeHandler) UpdateEmployeeForUser(c *gin.Context) {
	var req dto.UpdateEmployeeForUserRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.empUC.UpdateEmployeeForUser(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Employee Success!",
	})
}

func (h *EmployeeHandler) UpdateEmployeeForAdmin(c *gin.Context) {
	var req dto.UpdateEmployeeForAdminRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.empUC.UpdateEmployeeForAdmin(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Employee Success!",
	})
}

func (h *EmployeeHandler) UpdateAvatarForUser(c *gin.Context) {
    // 1. Ambil file dari request
    file, header, err := c.Request.FormFile("avatar")
    if err != nil {
		c.JSON(400, gin.H{"error": "avatar file not found!"})
        return
    }
    defer file.Close()

    // 2. Forward file ke Python service
	result, err := h.pythonClient.UploadAvatar(c.Request.Context(), file, header)
    if err != nil {
        c.JSON(500, gin.H{"error": "failed to upload avatar"})
        return
    }

	var req dto.UpdateEmployeeForUserRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	req.Avatar = &result.URL

    // 4. Update avatar di DB via usecase yang sudah ada
    if err := h.empUC.UpdateEmployeeForUser(c.Request.Context(), req); err != nil {
        c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
        return
    }

    c.JSON(200, gin.H{"message": "Changed Avatar Success!", "avatar_url": result.URL})
}

func (h *EmployeeHandler) UpdateAvatarForAdmin(c *gin.Context) {
    // 1. Ambil file dari request
    file, header, err := c.Request.FormFile("avatar")
    if err != nil {
        c.JSON(400, gin.H{"error": "avatar file not found!"})
        return
    }
    defer file.Close()

    // 2. Forward file ke Python service
    result, err := h.pythonClient.UploadAvatar(c.Request.Context(), file, header)
    if err != nil {
        c.JSON(500, gin.H{"error": "failed to upload avatar"})
        return
    }

	var req dto.UpdateEmployeeForAdminRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	req.Avatar = &result.URL

    // 4. Update avatar di DB via usecase yang sudah ada
    if err := h.empUC.UpdateEmployeeForAdmin(c.Request.Context(), req); err != nil {
        c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
        return
    }

    c.JSON(200, gin.H{"message": "Changed Avatar Success!", "avatar_url": result.URL})
}