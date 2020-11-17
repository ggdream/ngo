package app

import (
	"fmt"
	"github.com/ggdream/ngo/conf"
	"github.com/ggdream/ngo/middle"
	"github.com/gin-gonic/gin"
	"net/http"
)


func New() error {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(middle.Cors)
	engine.Use(gin.Recovery())

	if conf.PrintLog {
		engine.Use(middle.Logger())
	}

	engine.StaticFS(conf.PATH, http.Dir(conf.STATIC))

	addr := fmt.Sprintf("%s:%d", conf.HOST, conf.PORT)
	return engine.Run(addr)
}
