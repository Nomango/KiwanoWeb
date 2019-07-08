package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/nomango/KiwanoWeb/controllers"
	"github.com/nomango/KiwanoWeb/modules/templates"
)

func Setup(engine *gin.Engine) {
	engine.Delims("{{", "}}")
	engine.SetFuncMap(templates.FuncMap)
	engine.LoadHTMLFiles("./views/index.html", "./views/goPackage.html")
	engine.Static("/static", "./static")

	engine.GET("/", controllers.Home)
	engine.GET("/kiwano", controllers.GoPackageService)
}