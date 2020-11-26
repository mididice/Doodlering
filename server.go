package main

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.GET("/start", getStart)
	r.GET("/playing/:key/:sequence", getPlay)
	r.POST("/play/:key/:sequence", postPlay)
	r.GET("/story/:key/sequence", getStory)
	r.GET("/ending/:key", getEnd)
	r.GET("/ending/:key/:sequence", getEnd)
	r.GET("/home", getHome)
	r.GET("/howto", getHowto)
	r.GET("/end/:key/:sequence")
	r.GET("/tale/:key/:sequence")
	r.Run(":8080")
}

func getHome(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
func getHowto(c *gin.Context) {

}
func getStart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "start",
	})
}
func getPlay(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "play",
	})
}
func postPlay(c *gin.Context) {

}
func getStory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "story",
	})
}
func getEnd(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "end",
	})
}
