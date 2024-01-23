package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func UrlRewrite(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		regex := c.Param("regex")
		if regex != "" {
			log.Println(regex)
			url := strings.Trim(c.Request.URL.Path, regex)

			delimiter := "@"
			index := strings.Index(regex, delimiter)
			if index != -1 {
				url = fmt.Sprintf("%s%s", url, regex[index+1:])
				regex = regex[:index]

				c.Request.URL.Path = url
				r.HandleContext(c)
			}
		}
	}
}
