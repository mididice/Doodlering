package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
)

func main() {
	DB, err := sql.Open("mysql", "root:pw@tcp(localhost:3306)/Doodlering")

	if err != nil {
		return
	}
	var key string
	result, err := DB.Query("SELECT * FROM Doodlering.Games;")
	if err != nil {
		return
	}
	for result.Next() {
		err = result.Scan(&key)
		fmt.Println(key)
	}
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.GET("/start", getStart)
	r.GET("/playing/:key/:sequence", getPlayingks)
	r.POST("/play/:key/:sequence", postPlayks)
	r.GET("/story/:key/sequence", getStoryks)
	r.GET("/ending/:key", getEndingk)
	r.GET("/ending/:key/:sequence", getEndingks)
	r.GET("/home", getHome)
	r.GET("/howto", getHowto)
	r.GET("/end/:key/:sequence", getEndks)
	r.GET("/tale/:key/:sequence", taleks)
	r.Run(":8080")
}

func getEndks(c *gin.Context) {

}
func getEndingks(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "ending.html", gin.H{})
}
func getEndingk(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing-end.html", gin.H{})
}
func taleks(c *gin.Context) {

}
func getHome(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func getHowto(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "howto.html", gin.H{})
}
func getStart(c *gin.Context) {
	uuid, _ := uuid.NewV4()
	c.JSON(200, gin.H{
		"key": uuid,
	})
	//insert DB
}
func getPlayingks(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing.html", gin.H{})
}
func postPlayks(c *gin.Context) {

}
func getStoryks(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "story.html", gin.H{})
}
func getEndk(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "end",
	})
}
