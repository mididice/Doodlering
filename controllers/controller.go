package controllers

import (
	"Doodlering/config"
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dbConnectInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DBEnv.User,
		config.DBEnv.Password,
		config.DBEnv.Host,
		config.DBEnv.Name)
	DB, err = sql.Open("mysql", dbConnectInfo)
	if err != nil {
		return err
	}
	return nil
}
func RedirectHome(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/home")
}
func EndingkEnd(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "ending-real.html", gin.H{})
}
func GetEndingks(c *gin.Context) {
	if c.Param("sequence") == "end" {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "ending-real.html", gin.H{})
	} else {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "ending.html", gin.H{})
	}
}
func GetStoryks(c *gin.Context) {
	if c.Param("sequence") == "end" {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "story-end.html", gin.H{})
	} else {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "story.html", gin.H{})
	}
}
func GetEndingk(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing-end.html", gin.H{})
}
func GetHome(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func GetHowto(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "howto.html", gin.H{})
}
func GetStart(c *gin.Context) {
	uuid, _ := uuid.NewV4()
	c.JSON(200, gin.H{
		"key": uuid,
	})
	_, err := DB.Exec("INSERT INTO `doodlering`.`Games` (`key`) VALUES ('" + uuid.String() + "');")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func GetPlayingks(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err := DB.Exec("INSERT INTO `doodlering`.`Play` (`Games_key`, `sequence`, `gen_date`) " +
		"VALUES ('" + key + "', '" + sequence + "', '" + now + "');")
	if err != nil {
		fmt.Println(err)
		return
	}
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
type Stories struct {
	Key      string
	Sentence string
	Date     string
}

func PostPlayks(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	var input Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	var id string
	DB.QueryRow("SELECT id FROM doodlering.Play " +
		"where Games_key = '" + key + "'AND sequence = '" + sequence + "';").Scan(&id)
	var query = ""
	for i, _ := range input.Coordinate {
		query = " INSERT INTO `doodlering`.`Coordinate` (`Play_id`, `Play_Games_key`, `x`, `y`, `dx`, `dy`) VALUES ('" + id + "', '" + key + "', '" + fmt.Sprintf("%f", input.Coordinate[i][0]) + "', '" + fmt.Sprintf("%f", input.Coordinate[i][1]) +
			"', '" + fmt.Sprintf("%f", input.Coordinate[i][2]) + "', '" + fmt.Sprintf("%f", input.Coordinate[i][3]) + "');"
		DB.Exec(query)
	}
	for i, _ := range input.Prediction {
		query = "INSERT INTO `doodlering`.`Words` (`label`, `Play_id`, `Play_Games_key`, `confidence`)" +
			"VALUES ('" + input.Prediction[i].Label + "', '" + id + "', '" + key + "', '" + fmt.Sprintf("%f", input.Prediction[i].Confidence) + "');"
		DB.Exec(query)
	}
}

func GetPlayks(c *gin.Context) {
	var sentence string
	var sentenceId, playId string
	key := c.Param("key")
	sequence := c.Param("sequence")
	query := "SELECT id FROM Play WHERE Games_key = '" + key + "' AND sequence = " + sequence + ";"
	DB.QueryRow(query).Scan(&playId)
	query = "SELECT Sentences_id FROM doodlering.Play_has_Sentences where Play_id = '" +
		playId + "' AND Play_Games_key = '" + key + "';"
	err := DB.QueryRow(query).Scan(&sentenceId)
	if err == sql.ErrNoRows {
		fmt.Println("no data")
		var playId string
		DB.QueryRow("SELECT id, sentence FROM doodlering.Sentences ORDER BY RAND() LIMIT 1;").Scan(&sentenceId, &sentence)
		query = "SELECT id FROM Play WHERE Games_key = '" + key + "' AND sequence = " + sequence + ";"
		DB.QueryRow(query).Scan(&playId)
		fmt.Println("you wn:")
		fmt.Println(playId)
		query = "INSERT INTO `doodlering`.`Play_has_Sentences` (`Play_id`, `Play_Games_key`, `Sentences_id`) " +
			"VALUES (" + playId + ", '" + key + "', " + sentenceId + ");"
		DB.Exec(query)
		c.JSON(200, Sentences{
			Sentence: sentence,
		})
		return
	}
	err = DB.QueryRow("SELECT sentence FROM doodlering.Sentences WHERE id = " + sentenceId + ";").Scan(&sentence)
	c.JSON(200, Sentences{
		Sentence: sentence,
	})
}

type Tale struct {
	Candidate  []string    `json:"candidate"`
	Coordinate [][]float64 `json:"coordinate"`
}

func Taleks(c *gin.Context) {
	key := c.Param("key")
	fmt.Println(key)
	sequence := c.Param("sequence")
	var id string
	DB.QueryRow("SELECT id FROM doodlering.Play " +
		"where Games_key = '" + key + "'AND sequence = '" + sequence + "';").Scan(&id)

	query := "SELECT label, confidence FROM doodlering.Words " +
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

	query = "SELECT x, y, dx, dy FROM doodlering.Coordinate " +
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
func GetEndks(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	var id string
	DB.QueryRow("SELECT id FROM doodlering.Play " +
		"where Games_key = '" + key + "'AND sequence = '" + sequence + "';").Scan(&id)
	var output End
	fmt.Println("id:" + id)
	var query string
	query = "SELECT label, confidence FROM doodlering.Words " +
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
	query = "SELECT x, y, dx, dy FROM doodlering.Coordinate " +
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
func GetEndk(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "playing-end.html", gin.H{})
}
func GetStories(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "stories.html", gin.H{})
}
func GetTales(c *gin.Context) {
	query := //"SELECT Games_key, gen_date, sentence FROM Play as p left join Play_has_Sentences as ps on p.id = ps.Play_id left join Sentences as s on ps.Sentences_id = s.id where sequence = 1 order by gen_date desc limit 100;"
		"SELECT Games_key, gen_date, sentence FROM Play as p " +
			"INNER JOIN Play_has_Sentences as ps ON p.id = ps.Play_id " +
			"LEFT JOIN Sentences as s ON ps.Sentences_id = s.id " +
			"GROUP BY p.Games_key HAVING count(p.Games_key) > 9 ORDER BY gen_date DESC LIMIT 100;"
	result, err := DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	var stories []Stories
	var key, sentence, date string

	for result.Next() {
		err = result.Scan(&key, &date, &sentence)
		slice := strings.Split(date, " ")
		if err != nil {
			fmt.Println(err)
		}
		tmp := Stories{
			Key:      key,
			Sentence: sentence,
			Date:     slice[0],
		}
		stories = append(stories, tmp)
	}
	c.JSON(200, stories)
}
func GetSentence(c *gin.Context) {
	key := c.Param("key")
	sequence := c.Param("sequence")
	fmt.Println(key + ":" + sequence)
	var playId, sentenceId, sentence string
	query := "SELECT id FROM Play WHERE Games_key = '" + key + "' AND sequence = " + sequence + ";"
	DB.QueryRow(query).Scan(&playId)

	query = "SELECT Sentences_id FROM Play_has_Sentences WHERE Play_id = " +
		playId + " AND Play_Games_key = '" + key + "';"
	DB.QueryRow(query).Scan(&sentenceId)

	query = "SELECT sentence FROM Sentences WHERE id = " + sentenceId + ";"
	DB.QueryRow(query).Scan(&sentence)

	c.JSON(200, Sentences{Sentence: sentence})
}
