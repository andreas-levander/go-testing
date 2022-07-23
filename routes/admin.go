package routes

import (
	"go-test/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/v1/admin")

	{
		admin.GET("/users", users)
	}
}

func users(c *gin.Context) {
	usrs := database.GetUsers()
	c.JSON(http.StatusOK, usrs)
}