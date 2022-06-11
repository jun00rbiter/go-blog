package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jun00rbiter/go-blog/app/routers"
)

func main() {
	engine := gin.Default()
	routers.RegisterAPIRouter(engine)
	engine.Run()
}
