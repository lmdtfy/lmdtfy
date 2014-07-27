package docker

import (
	"log"
	"os"

	"github.com/docker/docker/archive"
	d "github.com/fsouza/go-dockerclient"
)

func Build() {

	tr, err := archive.Tar("testdata", archive.Uncompressed)
	if err != nil {
		log.Fatal(err)
	}
	opts := d.BuildImageOptions{
		Name:         "test",
		InputStream:  tr,
		OutputStream: os.Stdout,
	}

	if err := client.BuildImage(opts); err != nil {
		log.Fatal(err)
	}

}
