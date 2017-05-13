package router

import (
	"github.com/edot92/gotherimg/controllers"
	_ "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.GET("/gambar/:pointx/:pointy", controllers.GetImageValue)
	Router = router
}
