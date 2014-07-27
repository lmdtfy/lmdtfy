package main

import (
	"github.com/lmdtfy/lmdtfy/pkg/build/docker"

	// "github.com/docker/docker/archive"
	// docker "github.com/fsouza/go-dockerclient"
)

func main() {
	docker.New()
	docker.Build()

}
