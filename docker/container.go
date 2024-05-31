package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainerStat struct {
	ID      string
	Image   string
	Command string
	Ports   []types.Port
	State   string
	Status  string
}

func GetContainerStats(cli *client.Client) ([]ContainerStat, error) {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	containerStats := []ContainerStat{}

	for _, ctr := range containers {
		containerStats = append(containerStats, ContainerStat{
			ID:      ctr.ID,
			Image:   ctr.Image,
			Command: ctr.Command,
			Ports:   ctr.Ports,
			State:   ctr.State,
			Status:  ctr.Status,
		})
	}

	return containerStats, nil
}
