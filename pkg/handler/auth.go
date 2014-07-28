package handler

import (
	"net/http"

	"github.com/lmdtfy/lmdtfy/pkg/model"

	"code.google.com/p/goauth2/oauth"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

func LinkGithub(c *gin.Context) {
	settings := &model.Settings{
		GitHubKey:    "33743eabe5768e51613d",
		GitHubSecret: "2fdae2837af8c7851d8c6a641f9d4090f8de82ff",
		GitHubURL:    "github.com",
	}
	// err := db.Settings.Get(settings)
	// if err != nil {
	// 	c.JSON(400, "Error parsing settings")
	// 	return
	// }
	var config = &oauth.Config{
		RedirectURL:  "https://7111ad68.ngrok.com/auth/login/github",
		TokenURL:     "https://" + settings.GitHubURL + "/login/oauth/access_token",
		AuthURL:      "https://" + settings.GitHubURL + "/login/oauth/authorize",
		ClientId:     settings.GitHubKey,
		ClientSecret: settings.GitHubSecret,
	}

	code := c.Request.FormValue("code")
	if len(code) == 0 {
		//scope := "repo,repo:status,user:email"
		redirect := config.AuthCodeURL("foo")
		//redirect.params.Set("scope", scope)
		http.Redirect(c.Writer, c.Request, redirect, http.StatusSeeOther)
		return
	}

	t := &oauth.Transport{Config: config}
	t.Exchange(code)

	client := github.NewClient(t.Client())
	repos, _, _ := client.Repositories.List("", nil)

	c.JSON(200, repos)
	return
}
