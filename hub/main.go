package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"os"
)

type HubInput struct {
	GitURL    string `json:"git_url" binding:"required"`
	GitBranch string `json:"git_branch" binding:"required"`
	RepoName  string `json:"repo_name" binding:"required"`
}

func GitClone() gin.HandlerFunc {
	return func(context *gin.Context) {
		var GitRequest HubInput

		err := context.BindJSON(&GitRequest)
		if err != nil {
			log.Println(err.Error())
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = git.PlainClone("/tmp/test/"+GitRequest.RepoName, false, &git.CloneOptions{
			URL:           GitRequest.GitURL,
			ReferenceName: plumbing.NewBranchReferenceName(GitRequest.GitBranch),
			SingleBranch:  true,
			Progress:      os.Stdout,
		})
		if err != nil {
			log.Println(err.Error())
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "Successfully cloned " + GitRequest.GitURL,
		})
	}
}

func main() {
	router := gin.Default()
	router.POST("/clone", GitClone())
	router.Run(":3000")
}
