// handler/auth_handler.go
package handler

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewAuthHandler(uc usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{uc}
}

// Login handler
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.UserUsecase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Bikin manual versi custom response-nya
	data := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"department": user.Department.Department,
	}

	response := entity.LoginResponse{
		Message: "login successful",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

// Register handler
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Name         string `json:"name"`
		DepartmentID int    `json:"department_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user := &entity.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Department: &entity.Department{
			ID: req.DepartmentID,
		},
	}

	if err := h.UserUsecase.Register(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
