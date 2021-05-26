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
	"github.com/pyama86/pdr/pdr"
)

func choiceAndRunDockerCommand(f func(string, *pdr.Repo, map[string]bool) (map[string]bool, error)) error {
	if interactive {
		f, err := getChoiceRepo()
		if err != nil {
			return err
		}
		filterRepo = f
	}
	done := map[string]bool{}
	for name, repo := range config.Repos {
		if filterRepo != "" && !strings.Contains(repo.Path, filterRepo) {
			continue
		}

		_, err := f(name, repo, done)
		if err != nil {
			return err
		}
	}
	return nil
}
func getChoiceRepo() (string, error) {
	dirs := []string{}
	for _, r := range config.Repos {
		dirs = append(dirs, r.Path)
	}

	var choice = bytes.Buffer{}
	pecoCli := peco.New()
	pecoCli.Argv = nil
	pecoCli.Stdin = bytes.NewBufferString(strings.Join(dirs, "\n"))
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
