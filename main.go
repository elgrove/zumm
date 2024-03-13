package main

import (
	"net/http"
	"zumm/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	initialiseRoutes(r)
	return r
}

func initialiseRoutes(r *gin.Engine) {
	r.GET("/", helloWorldHandler)
	r.GET("/user/create", userCreateHandler)
}

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func userCreateHandler(c *gin.Context) {
	user := models.CreateRandomUser()
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
		"name":     user.Name,
		"gender":   user.Gender,
		"age":      user.Age,
	})
}

func main() {
	r := setupRouter()
	r.Run()
}
