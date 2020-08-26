package middle

import "github.com/gin-gonic/gin"

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,POTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Length,Range")
	c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
	c.Header("Access-Control-Max-Age", "172800")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
	}
	c.Next()
}
