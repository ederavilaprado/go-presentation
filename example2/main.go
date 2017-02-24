package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	router.GET("/sleep/:time", func(c *gin.Context) {
		t, _ := strconv.ParseInt(c.Param("time"), 10, 64)
		time.Sleep(time.Duration(t) * time.Second)
		c.String(http.StatusOK, "I've been sleeping for %d seconds", t)
	})

	router.Run("localhost:8080")
}
