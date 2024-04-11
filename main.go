package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "Hello World",
	})
}

func GetUsernameAndPassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}

func PostAccount(c *gin.Context) {
	var user Account
	c.BindJSON(&user)
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func PutPathParameter(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}

func DeleteQuery(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{
		"message": username + " deleted!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/helloworld", GetHelloWorld)
	r.GET("/", GetUsernameAndPassword)
	r.POST("/", PostAccount)
	r.PUT("/account/:username/:password", PutPathParameter)
	r.DELETE("/del", DeleteQuery)

	r.Run()
}
