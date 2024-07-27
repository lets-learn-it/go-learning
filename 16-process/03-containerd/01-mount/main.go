package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	CONTAINERD_SOCKET = "/run/containerd/containerd.sock"
	NAMESPACE         = "default"
)

func main() {
	if len(os.Args) < 2 {
		panic("numb of args < 2")
	}

	IMAGE := os.Args[1]
	fmt.Printf("Image: %s\n", IMAGE)
	fmt.Printf("Args: %v\n", os.Args[2:])

	client, err := containerd.New(CONTAINERD_SOCKET)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), NAMESPACE)

	version, err := client.Version(ctx)
	if err != nil {
		fmt.Printf("error retrieving containerd version: %v\n", err)
		return
	}

	fmt.Printf("containerd version: %s %s\n", version.Version, version.Revision)

	// get image
	image, err := client.GetImage(ctx, IMAGE)
	if err != nil {
		fmt.Println("image locally not found, pulling.")
		image, err = client.Pull(ctx, IMAGE, containerd.WithPullUnpack)
		if err != nil {
			fmt.Printf("couldn't pull image %s, %s\n", IMAGE, err)
			return
		}
	}

	i, err := image.Spec(ctx)
	if err != nil {
		fmt.Printf("error while getting config from image: %s", err)
	}

	iconfig := i.Config

	// mount
	mount := specs.Mount{
		Source:      "/tmp",
		Destination: "/mnt/models",
		Options:     []string{"rbind", "rw"},
		Type:        "rbind",
	}

	container, err := client.NewContainer(ctx, "alpine",
		containerd.WithNewSnapshot("alpine", image),
		containerd.WithNewSpec(
			oci.WithImageConfig(image),
			oci.WithProcessArgs(append(iconfig.Entrypoint, os.Args[2:]...)...),
			oci.WithMounts([]specs.Mount{mount}),
		),
	)
	if err != nil {
		fmt.Printf("error creating container: %s\n", err)
	}

	// start container AKS create task
	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		fmt.Printf("error creating task: %s\n", err)
	}

	start := time.Now()
	var end time.Time

	// start task
	if err = task.Start(ctx); err != nil {
		task.Delete(ctx)
		fmt.Printf("error starting task: %s\n", err)
	}

	for {
		status, err := task.Status(ctx)
		if err != nil {
			fmt.Printf("error while getting status: %s\n", err)
			return
		}

		fmt.Println(status.Status)
		if status.Status == containerd.Stopped {
			end = status.ExitTime
			break
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Time taken: %v\n", end.Sub(start))

	// clean task & container
	exit, err := task.Delete(ctx)
	if err != nil {
		fmt.Printf("error while deleting task: %s", err)
	}

	fmt.Println(exit.ExitCode())

	if container.Delete(ctx) != nil {
		fmt.Printf("error while deleting container: %s", err)
	}
}
