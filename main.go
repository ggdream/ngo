package main


import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

const (
	defaultAddr  = "127.0.0.1:80"
	defaultEntry = "./"
)

type Serve struct {
	Addr	string
	Entry	string
}

func main() {
	args := parse()
	service(args)
}

func parse() Serve {
	args := os.Args

	if l := len(args); l < 2 {
		return Serve{defaultAddr, defaultEntry}
	} else if l == 2 {
		return Serve{args[1], defaultEntry}
	} else if l == 3 {
		return Serve{args[1], args[2]}
	}

	fmt.Println("error: you passed extra args.")
	os.Exit(-1)
	return Serve{}
}

func service(serve Serve){
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,POTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Length,Range")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
		c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS"{
			c.AbortWithStatus(204)
		}
		c.Next()
	})

	r.Static("/", serve.Entry)

	r.Run(serve.Addr)
}
