package cmd

import (
	"os"
	"os/exec"
	"strings"

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

func runUporDown(upordown []string) error {
	for _, repo := range config.Repos {
		if filterRepo != "" && !strings.Contains(repo.Path, filterRepo) {
			continue
		}
		logrus.Infof("on %s", repo.Path)

		if upordown[0] == "up" {
			for _, c := range repo.UpPreHookCommands {
				cmd := exec.Command("/bin/bash", "-l", "-c", c)
				cmd.Dir = repo.Path
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return err
				}
			}
		}
		cmd := exec.Command("docker-compose", upordown...)
		cmd.Dir = repo.Path
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
func runUpCommand() error {
	return runUporDown([]string{"up", "-d"})
}

var filterRepo string

func init() {
	upCmd.PersistentFlags().StringVar(&filterRepo, "repo", "", "target repository(default all)")
	rootCmd.AddCommand(upCmd)
}
