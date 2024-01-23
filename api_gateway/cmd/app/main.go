package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	//router.GET("/catalog/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/catalog/:regex", func(c *gin.Context) {
		r, err := regexp.Compile(`.+@.+`)
		if err != nil {
			panic(err)
			return
		}
		url := c.Param("regex")
		if r.MatchString(url) == true {
			c.String(http.StatusOK, "match")
		} else {
			c.String(http.StatusOK, "not match")
		}
	})

	router.Run(":8080")
}
