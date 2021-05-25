package cmd

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Show login to docker container",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runLoginCommand(); err != nil {
			logrus.Error(err)
		}
	},
}

func runLoginCommand() error {
	cn, err := getContainerName()
	if err != nil {
		return err
	}

	cmd := exec.Command("docker", "exec", "-it", cn, loginShell)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var loginShell string

func init() {
	loginCmd.PersistentFlags().StringVar(&loginShell, "shell", "/bin/bash", "container login shell")
	rootCmd.AddCommand(loginCmd)
}
