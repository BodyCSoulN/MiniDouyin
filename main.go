package main

import (
	"fmt"

	"github.com/MiniDouyin/config"
	_ "github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Conf.Server.Mode)
	r := gin.Default()
	r.Static("/static", "./public")
	r.Run(fmt.Sprintf("%s:%s", config.Conf.Server.Host, config.Conf.Server.Port))
}
