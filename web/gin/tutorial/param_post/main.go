package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		//c.String(http.StatusOK,string(body)) //传入json

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")
		c.String(http.StatusOK, "%s %s %s", firstName, lastName, string(body))
	})
	r.Run(":8080")
}
