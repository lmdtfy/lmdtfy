package handler

import (
	"log"

	gh "github.com/lmdtfy/lmdtfy/pkg/vcs"

	"github.com/gin-gonic/gin"
)

func Hook(c *gin.Context) {
	if c.Request.Header.Get("X-Github-Event") == "ping" {
		c.JSON(200, "Working")
		return
	}
	hook := gh.PostReceiveHook{}
	err := c.ParseBody(&hook)

	if err != nil {
		log.Println(err)
		c.JSON(400, err)
		return
	}

	c.JSON(200, hook)
	return
}
