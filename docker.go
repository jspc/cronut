package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
)

func Start(j Job) {
	cmd := strslice.StrSlice(j.Args)
	name := fmt.Sprintf("%s-%d", j.Name, time.Now().Unix())

	cb, err := Docker.ContainerCreate(
		context.Background(),
		&container.Config{
			Hostname: j.Name,
			Env:      j.Env,
			Cmd:      cmd,
			Image:    j.Container,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		name,
	)

	if err != nil {
		panic(err)
	}

	err = Docker.ContainerStart(
		context.Background(),
		cb.ID,
		types.ContainerStartOptions{},
	)

	L.Print(err)
}
