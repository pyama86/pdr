package cmd

import (
	"os"
	"os/exec"

	"github.com/pyama86/pdr/pdr"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "run docker-compose up on each repository dir",
	Long:  `Running docker-compose up command on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runUpCommand(); err != nil {
			logrus.Error(err)
		}
	},
}

func runUpCommand() error {
	return choiceAndRunDockerCommand(recRunUpCommand)
}

func recRunUpCommand(name string, repo *pdr.Repo, done map[string]bool) (map[string]bool, error) {
	if repo == nil {
		logrus.Warnf("%s is notfound", name)
	}
	if !done[name] && repo != nil {
		for _, depend := range repo.Depends {
			if depend != name {
				d, err := recRunUpCommand(depend, config.Repos[depend], done)
				if err != nil {
					return nil, err
				}
				done = d
			}
		}

		if !done[name] {
			logrus.Infof("on %s", repo.Path)
			for _, c := range repo.UpPreHookCommands {
				cmd := exec.Command("/bin/bash", "-l", "-c", c)
				cmd.Dir = repo.Path
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return nil, err
				}
			}

			cmd := exec.Command("docker-compose", "up", "-d")
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

var filterRepo string
var interactive bool

func init() {
	upCmd.PersistentFlags().StringVar(&filterRepo, "repo", "", "target repository(default all)")
	upCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "choose target repository")
	rootCmd.AddCommand(upCmd)
}
