package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"os"
	//"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	//time.Sleep(30 * time.Second) // in case of debugging uncomment this line.
	pulumi.Run(DockerBuildAndPublish)
}

func DockerBuildAndPublish(ctx *pulumi.Context) error {
	// Populate the registry info (creds and endpoint).
	registryInfo := docker.ImageRegistryArgs{
		Server:   pulumi.String("docker.io"),
		Username: pulumi.String(os.Getenv("DOCKER_USERNAME")),
		Password: pulumi.String(os.Getenv("DOCKER_PASSWORD")),
	}

	// Build and publish the container image.
	image, err := docker.NewImage(ctx, "clientapi", &docker.ImageArgs{
		Build:     &docker.DockerBuildArgs{Context: pulumi.String("../clientapi")}, //path to the dockerfile
		ImageName: pulumi.String("fafg/clientapi"),
		Registry:  registryInfo,
	})

	// Export the base and specific version image name.
	if image != nil {
		ctx.Export("baseImageName", image.BaseImageName)
		ctx.Export("fullImageName", image.ImageName)
	}

	return err
}