package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		fristName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default_last_name")
		c.String(http.StatusOK, "%s %s", fristName, lastName)
	})
	r.Run(":8080")
}
