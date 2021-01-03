package main

import (
	"Doodlering/config"
	"Doodlering/controllers"
	"Doodlering/routers"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("start!")
	err := config.SetupEnv()
	if err != nil {
		fmt.Println("fail to parse")
		return
	}

	err = controllers.InitDB()
	if err != nil {
		fmt.Println("fail to open db")
		return
	}
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// gin.SetMode(gin.ReleaseMode)
	r := routers.InitRouter()
	// r.Run()
	server := &http.Server{
		Addr:    "",
		Handler: r,
	}
	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()
}
