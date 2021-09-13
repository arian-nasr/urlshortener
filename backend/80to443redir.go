package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/*path", func(c *gin.Context) {
		c.Redirect(302, "https://onebounce.me/" + c.Param("variable"))
	})
	r.Run("10.128.0.3:80")
}
