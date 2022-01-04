package user

import (
	"net/http"

	"github.com/Sanjungliu/assesment-user-service/internal/auth"
	"github.com/Sanjungliu/assesment-user-service/internal/user"
	"github.com/Sanjungliu/assesment-user-service/pkg/helper"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to register user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, refreshToken, err := h.authService.GenerateToken(newUser.UserID, newUser.Role)
	if err != nil {
		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := user.FormatUser(newUser, token, refreshToken)

	response := helper.APIResponse("Account succeed to registered", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := map[string]interface{}{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, refreshToken, err := h.authService.GenerateToken(loggedInUser.UserID, loggedInUser.Role)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatted := user.FormatUser(loggedInUser, token, refreshToken)
	response := helper.APIResponse("Succeed to Login", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatted := user.FormatUser(currentUser, "", "")

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatted)

	c.JSON(http.StatusOK, response)
}
