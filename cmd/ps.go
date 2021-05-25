package cmd

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "run docker-compose ps on each repository dir",
	Long:  `Running docker-compose ps command on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runPSCommand(); err != nil {
			logrus.Error(err)
		}
	},
}

func runPSCommand() error {
	for _, repo := range config.Repos {
		logrus.Infof("on %s", repo.Path)
		cmd := exec.Command("docker-compose", "ps", "-a")
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

func init() {
	rootCmd.AddCommand(psCmd)
}
