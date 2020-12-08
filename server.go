package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
)

var DB *sql.DB

func main() {
	fmt.Println("start")
	var err error
	DB, err = sql.Open("mysql", "root:1q2w3e4r5T!@@tcp(localhost:3306)/DOODLERING")

	if err != nil {
		fmt.Println("fail to open db")
		return
	}

	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.Static("/static", "./static")
	r.GET("/start", getStart)
	r.GET("/playing/:key/:sequence", getPlayingks)
	r.POST("/play/:key/:sequence", postPlayks)
	r.GET("/story/:key/:sequence", getStoryks)
	r.GET("/ending/:key", getEndingk)
	r.GET("/ending/:key/:sequence", getEndingks)
	r.GET("/home", getHome)
	r.GET("/howto", getHowto)
	r.GET("/end/:key/:sequence", getEndks)
	r.GET("/tale/:key/:sequence", taleks)
	r.GET("/play/:key/:sequence", getPlayks)
	r.GET("/", redirectHome)
	r.Run(":8080")
}
func redirectHome(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/home")
}
func endingkEnd(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "ending-real.html", gin.H{})
}
func getEndingks(c *gin.Context) {
	if c.Param("sequence") == "end" {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "ending-real.html", gin.H{})
	} else {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "ending.html", gin.H{})
	}
}
func getStoryks(c *gin.Context) {
	if c.Param("sequence") == "end" {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "story-end.html", gin.H{})
	} else {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "story.html", gin.H{})
	}
}
func getEndingk(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing-end.html", gin.H{})
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
	DB.Exec("INSERT INTO `DOODLERING`.`Games` (`key`) VALUES ('" + uuid.String() + "');")
}
func getPlayingks(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing.html", gin.H{})
}

type Game struct {
	Prediction []*Play     `json: "prediction"`
	Coordinate [][]float64 `json: "coordinate"`
}
type Play struct {
	Label      string  `json : "label"`
	Confidence float64 `json : "confidence"`
}
type End struct {
	Coordinate [][]float64 `json: "coordinate"`
	Answer     []*Play     `json: "answer"`
}
type Sentences struct {
	Sentence string `json: "sentence"`
}

func postPlayks(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	var input Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	DB.Exec("INSERT INTO `DOODLERING`.`Play` (`Games_key`, `sequence`) " +
		"VALUES ('" + key + "', '" + sequence + "');")
	var id string
	DB.QueryRow("SELECT id FROM DOODLERING.Play " +
		"where Games_key = '" + key + "'AND sequence = '" + sequence + "';").Scan(&id)
	var query = ""
	for i, _ := range input.Coordinate {
		query = " INSERT INTO `DOODLERING`.`Coordinate` (`Play_id`, `Play_Games_key`, `x`, `y`, `dx`, `dy`) VALUES ('" + id + "', '" + key + "', '" + fmt.Sprintf("%f", input.Coordinate[i][0]) + "', '" + fmt.Sprintf("%f", input.Coordinate[i][1]) +
			"', '" + fmt.Sprintf("%f", input.Coordinate[i][2]) + "', '" + fmt.Sprintf("%f", input.Coordinate[i][3]) + "');"
		DB.Exec(query)
	}
	for i, _ := range input.Prediction {
		query = "INSERT INTO `Doodlering`.`Words` (`label`, `Play_id`, `Play_Games_key`, `confidence`)" +
			"VALUES ('" + input.Prediction[i].Label + "', '" + id + "', '" + key + "', '" + fmt.Sprintf("%f", input.Prediction[i].Confidence) + "');"
		DB.Exec(query)
	}
}

func getPlayks(c *gin.Context) {
	// key := c.Param("key")
	sequence := c.Param("sequence")

	var sentence string
	DB.QueryRow("SELECT sentence FROM Sentences WHERE id =" + sequence + ";").Scan(&sentence)
	c.JSON(200, Sentences{
		Sentence: sentence,
	})
}

type Tale struct {
	Candidate  []string    `json:"candidate"`
	Coordinate [][]float64 `json:"coordinate"`
}

func taleks(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	var id string
	DB.QueryRow("SELECT id FROM DOODLERING.Play " +
		"where Games_key = '" + key + "'AND sequence = '" + sequence + "';").Scan(&id)

	query := "SELECT label, confidence FROM DOODLERING.Words " +
		"where Play_id = " + id + " AND Play_Games_key = '" + key + "';"
	rows, err := DB.Query(query)
	if err != nil {
		return
	}
	var words []Play
	var label string
	var confidence float64
	for rows.Next() {
		rows.Scan(&label, &confidence)
		words = append(words, Play{Label: label, Confidence: confidence})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].Confidence > words[j].Confidence
	})
	for _, tmp := range words {
		fmt.Println(tmp.Confidence)
	}
	var tale Tale
	var tmp []string
	for i := 0; i < 3; i++ {
		tmp = append(tmp, words[i].Label)
	}
	tale.Candidate = tmp

	query = "SELECT x, y, dx, dy FROM DOODLERING.Coordinate " +
		"WHERE Play_Games_key = '" + key + "' AND Play_id = " + id + ";"
	rows, err = DB.Query(query)
	var x, y, dx, dy float64
	var cordiOutput [][]float64
	for rows.Next() {
		rows.Scan(&x, &y, &dx, &dy)
		var tmp []float64
		tmp = append(tmp, x, y, dx, dy)
		cordiOutput = append(cordiOutput, tmp)
	}
	tale.Coordinate = cordiOutput
	c.JSON(200, tale)
}
func getEndks(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	var id string
	DB.QueryRow("SELECT id FROM DOODLERING.Play " +
		"where Games_key = '" + key + "'AND sequence = '" + sequence + "';").Scan(&id)
	var output End
	fmt.Println("id:" + id)
	var query string
	query = "SELECT label, confidence FROM DOODLERING.Words " +
		"where Play_id = " + id + " AND Play_Games_key = '" + key + "';"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("fail to fetch")
		return
	}
	var label string
	var confidence float64
	var answers []*Play
	for rows.Next() {
		rows.Scan(&label, &confidence)
		row := Play{
			Confidence: confidence,
			Label:      label,
		}
		answers = append(answers, &row)
	}
	output.Answer = answers
	query = "SELECT x, y, dx, dy FROM DOODLERING.Coordinate " +
		"WHERE Play_Games_key = '" + key + "' AND Play_id = " + id + ";"
	rows, err = DB.Query(query)
	var x, y, dx, dy float64
	var cordiOutput [][]float64
	for rows.Next() {
		rows.Scan(&x, &y, &dx, &dy)
		var tmp []float64
		tmp = append(tmp, x, y, dx, dy)
		cordiOutput = append(cordiOutput, tmp)
	}
	output.Coordinate = cordiOutput
	c.JSON(200, output)
}
func getEndk(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing-end.html", gin.H{})
}
