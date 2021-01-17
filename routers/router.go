package routers

import (
	"Doodlering/controllers"
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.HTMLRender = ginview.Default()
	r.Static("/static", "./static")
	r.StaticFS("/model", http.Dir("model"))
	r.GET("/start", controllers.GetStart)
	r.GET("/playing/:key/:sequence", controllers.GetPlayingks)
	r.POST("/play/:key/:sequence", controllers.PostPlayks)
	r.GET("/story/:key/:sequence", controllers.GetStoryks)
	r.GET("/ending/:key", controllers.GetEndingk)
	r.GET("/ending/:key/:sequence", controllers.GetEndingks)
	r.GET("/home", controllers.GetHome)
	r.GET("/howto", controllers.GetHowto)
	r.GET("/end/:key/:sequence", controllers.GetEndks)
	r.GET("/tale/:key/:sequence", controllers.Taleks)
	r.GET("/play/:key/:sequence", controllers.GetPlayks)
	r.GET("/", controllers.RedirectHome)
	r.GET("/tales", controllers.GetTales)
	r.GET("/stories", controllers.GetStories)
	r.GET("/sentence/:key/:sequence", controllers.GetSentence)

	return r
}
