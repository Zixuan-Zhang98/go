package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")               // for a directory, use default FileSystem: gin.Dir()
	r.StaticFS("/static", http.Dir("static"))     // for a custom http.FileSystem
	r.StaticFile("/favicon.ico", "./favicon.ico") // for a single file
	r.Run()
}
