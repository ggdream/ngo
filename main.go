package main

import (
	"fmt"
	"github.com/ggdream/ngo/middle"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	defaultAddr  = "127.0.0.1:80"
	defaultEntry = "./"
)

type Serve struct {
	Addr  string
	Entry string
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

func service(serve Serve) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	l := gin.New()

	go func() {
		for {
			if len(middle.Id) == 0 && len(middle.Ch) != 0 {
				<-middle.Ch
			}
		}
	}()

	r.Use(gin.Recovery())
	r.Use(middle.Cors)
	r.Use(middle.Record())

	l.Use(gin.Recovery())
	l.Use(middle.Cors)

	r.Static("/", serve.Entry)
	l.GET("/log", middle.Logger(middle.Ch, middle.Id))

	go func() {
		if err := l.Run(":8888"); err != nil {
			panic(err.Error())
		}
	}()
	if err := r.Run(serve.Addr); err != nil {
		panic(err.Error())
	}
}
