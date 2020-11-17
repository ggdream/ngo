package middle

import (
	"github.com/ggdream/ngo/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(conf.Origin, ","))
	c.Header("Access-Control-Allow-Methods", "GET,POST,POTIONS,HEAD")
	c.Header("Access-Control-Allow-Headers", "Content-Length,Range")
	c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
	c.Header("Access-Control-Max-Age", "172800")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
