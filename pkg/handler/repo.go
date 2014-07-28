package handler

import (
	"github.com/gin-gonic/gin"

	db "github.com/lmdtfy/lmdtfy/pkg/store"
)

func CreateRepo(c *gin.Context) {
	query := db.Repos().Insert(map[string]string{"Test": "test"})
	store.RunWrite("sad", "asd")

	c.JSON(200, "")
	return
}
