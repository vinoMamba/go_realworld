package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba.com/go_realworld/handler"
)

func SetupHttpServer() {
	r := gin.Default()
	handler.UserHandler(r)
	r.Run(":3000") // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}
