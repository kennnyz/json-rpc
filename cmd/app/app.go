package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kennnyz/lamoda/internal/app"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app.Run("")
}
