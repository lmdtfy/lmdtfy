package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/lmdtfy/lmdtfy/pkg/database"
	"github.com/lmdtfy/lmdtfy/pkg/handler"
	"github.com/lmdtfy/lmdtfy/pkg/store"
)

var (
	listenAddr     string
	address        string
	dbName         string
	staticDir      string
	insertTestData bool
)

func main() {
	// docker.New()
	// docker.Build()
	flag.StringVar(&address, "dbAddress", "localhost:28015", "")
	flag.StringVar(&dbName, "db", "dev_lmdtfy", "")
	flag.Parse()

	if err := database.Connect(address, dbName); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	if err := store.Connect(address, dbName); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	database.SetupDB(insertTestData)

	r := gin.Default()

	r.POST("/hook", handler.Hook)

	m := r.Group("/api")
	//m.POST("/session", handler.Login)
	m.GET("/repo", handler.CreateRepo)
	m.GET("/auth/login/github", handler.LinkGithub)
	r.Run(":4000")
}
