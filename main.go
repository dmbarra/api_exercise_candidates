package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	docs "github.com/dmbarra/api_exercise_candidates/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var users = []user{}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping exercise
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {user} users
// @Router /exercise/users [get]
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/exercise")
		{
			eg.GET("/users", getUsers)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}
