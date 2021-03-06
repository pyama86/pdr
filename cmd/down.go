package cmd

import (
	"os"
	"os/exec"

	"github.com/pyama86/pdr/pdr"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	funk "github.com/thoas/go-funk"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "run docker-compose down on each repository dir",
	Long:  `Running docker-compose down command on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runDownCommand(); err != nil {
			logrus.Error(err)
		}
	},
}

func runDownCommand() error {
	return choiceAndRunDockerCommand(recRunDownCommand)
}

func recRunDownCommand(name string, repo *pdr.Repo, done map[string]bool) (map[string]bool, error) {
	if repo == nil {
		logrus.Warnf("%s is notfound", name)
	}
	if !done[name] && repo != nil {
		for dependName, r := range config.Repos {
			if funk.ContainsString(r.Depends, name) && dependName != name {
				d, err := recRunDownCommand(dependName, config.Repos[dependName], done)
				if err != nil {
					return nil, err
				}
				done = d
			}
		}

		if !done[name] {
			logrus.Infof("on %s", repo.Path)
			for _, c := range repo.DownPreHookCommands {
				cmd := exec.Command("/bin/bash", "-l", "-c", c)
				cmd.Dir = repo.Path
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return nil, err
				}
			}

			cmd := exec.Command("docker-compose", "down")
			cmd.Dir = repo.Path
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return nil, err
			}
			done[name] = true
		}
	}
	return done, nil
}

func init() {
	downCmd.PersistentFlags().StringVar(&filterRepo, "repo", "", "target repository(default all)")
	downCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "choose target repository")
	rootCmd.AddCommand(downCmd)
}
