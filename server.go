package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/start", getStart)
	r.GET("/play/:key/:sequence", getPlay)
	r.POST("/play/:key/:sequence", postPlay)
	r.GET("/story/:key/sequence", getStory)
	r.GET("/end/:key/:sequence", getEnd)
	r.Run()
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
