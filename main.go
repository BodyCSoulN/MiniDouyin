package main

import (
	"fmt"
	"github.com/MiniDouyin/config"
	"github.com/MiniDouyin/router"
	_ "github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Conf.Server.Mode)
	r := router.Router()
	r.Run(fmt.Sprintf("%s:%s", config.Conf.Server.Host, config.Conf.Server.Port))
}
