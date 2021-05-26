package cmd

import (
	"os"
	"os/exec"

	"github.com/pyama86/pdr/pdr"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	funk "github.com/thoas/go-funk"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "run docker-compose rm on each repository dir",
	Long:  `Running docker-compose rm command on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runRemoveCommand(); err != nil {
			logrus.Error(err)
		}
	},
}

func runRemoveCommand() error {
	return choiceAndRunDockerCommand(recRunRemoveCommand)
}

func recRunRemoveCommand(name string, repo *pdr.Repo, done map[string]bool) (map[string]bool, error) {
	if repo == nil {
		logrus.Warnf("%s is notfound", name)
	}
	if !done[name] && repo != nil {
		for dependName, r := range config.Repos {
			if funk.ContainsString(r.Depends, name) && dependName != name {
				d, err := recRunRemoveCommand(dependName, config.Repos[dependName], done)
				if err != nil {
					return nil, err
				}
				done = d
			}
		}

		if !done[name] {
			logrus.Infof("on %s", repo.Path)
			for _, c := range repo.RemovePreHookCommands {
				cmd := exec.Command("/bin/bash", "-l", "-c", c)
				cmd.Dir = repo.Path
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return nil, err
				}
			}

			cmd := exec.Command("docker-compose", "rm", "-f")
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
	rmCmd.PersistentFlags().StringVar(&filterRepo, "repo", "", "target repository(default all)")
	rmCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "choose target repository")
	rootCmd.AddCommand(rmCmd)
}
