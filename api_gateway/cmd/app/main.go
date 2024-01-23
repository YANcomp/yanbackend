package main

import (
	"github.com/YANcomp/yanbackend/api_gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(middleware.UrlRewrite(router))

	router.GET("/catalog/types", func(c *gin.Context) {
		log.Println("test")
		c.String(http.StatusOK, "types")
	})

	router.GET("/catalog/categories", func(c *gin.Context) {
		c.String(http.StatusOK, "categories")
	})

	router.GET("/catalog/:regex", func(c *gin.Context) {
		c.String(http.StatusOK, "regex")
	})

	router.GET("/products", func(c *gin.Context) {
		c.String(http.StatusOK, "products")

	})

	router.Run(":8080")
}
