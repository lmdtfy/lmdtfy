package handler

import (
	"github.com/gin-gonic/gin"

	db "github.com/lmdtfy/lmdtfy/pkg/store"
)

func CreateRepo(c *gin.Context) {
	query := db.Repos().Insert(map[string]string{"Test": "test"})
	store.RunWrite(query)

	c.JSON(200, "")
	return
}

func GetAllRepos(c *gin.Context) {
	in := []map[string]interface{}{}
	query := db.Repos().Term
	store.All(&in, query)

	c.JSON(200, in)
}
