package cmd

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/peco/peco"
)

func getContainerName() (string, error) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		return "", err
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: false,
	})
	if len(containers) == 0 {
		return "", errors.New("runnning container is notfound")
	}

	containerNames := []string{}
	for _, c := range containers {
		containerNames = append(containerNames, c.Names[0][1:])
	}

	var choice = bytes.Buffer{}
	pecoCli := peco.New()
	pecoCli.Argv = nil
	pecoCli.Stdin = bytes.NewBufferString(strings.Join(containerNames, "\n"))
	pecoCli.Stdout = &choice

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := pecoCli.Run(ctx); err != nil {
		if reflect.TypeOf(err).Name() == "errCollectResults" {
			pecoCli.PrintResults()
		} else {
			return "", err
		}
	}
	return strings.Replace(choice.String(), "\n", "", 1), nil
}