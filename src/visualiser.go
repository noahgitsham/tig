package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func visualise() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Tig",
			"content": "visualise your repository",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
