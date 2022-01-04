package httpserver

import (
	"net/http"
	"strings"

	"github.com/Sanjungliu/assesment-user-service/internal/app"
	"github.com/Sanjungliu/assesment-user-service/internal/auth"
	user_handler "github.com/Sanjungliu/assesment-user-service/internal/httpserver/handler/user"
	"github.com/Sanjungliu/assesment-user-service/internal/user"
	"github.com/Sanjungliu/assesment-user-service/pkg/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(app *app.App) *http.Server {
	s := &http.Server{
		Handler: buildRoute(app),
	}
	return s
}

func buildRoute(app *app.App) http.Handler {
	auth := authMiddleware(app.Auth, app.User)
	userHandler := user_handler.NewUserHandler(app.User, app.Auth)

	router := gin.Default()
	router.Use(
		cors.Default(),
	)

	api := router.Group("/api/v1")

	api.POST("/users/sessions", userHandler.Login)
	api.POST("/users/signup", userHandler.RegisterUser)

	api.Use(auth).GET("/users", userHandler.FetchUser)
	api.Use(auth).GET("/users/:id", userHandler.FetchUser)

	router.Run()
	return router
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := claim["user_id"].(string)
		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
