package docker

import (
	"log"
	"os"

	d "github.com/fsouza/go-dockerclient"
)


var client *d.Client

// New creates an instance of the Docker Client
func New() *d.Client {
	var err error
	if os.Getenv("DOCKER_HOST") != "" {
		client, err = d.NewClient(os.Getenv("DOCKER_HOST"))
	} else {
		client, err = d.NewClient("DEFAULTHERE")
	}

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return client
}
