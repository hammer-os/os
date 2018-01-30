package main

import (
	"context"
	"fmt"
	"log"

	containerd "github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {
	c, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatalf("dialing containerd: %v", err)
	}
	defer c.Close()

	ctx := context.Background()
	ctx = namespaces.WithNamespace(ctx, "hammer")

	image, err := c.Pull(ctx, "docker.io/library/alpine:3.7")
	if err != nil {
		log.Fatalf("retrieving image: %v", err)
	}

	fmt.Println("found image docker.io/library/alpine:3.7: %s", image.Name())

	images, err := c.ListImages(ctx)
	if err != nil {
		log.Fatalf("listing images: %v", err)
	}
	fmt.Printf("%+v\n", images)
}
