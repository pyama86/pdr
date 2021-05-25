package cmd

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show log in docker container",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runLogCommand(); err != nil {
			logrus.Error(err)
		}
	},
}

func runLogCommand() error {
	cn, err := getContainerName()
	if err != nil {
		return err
	}

	command := []string{"logs", cn}
	if tailNumber != 0 {
		command = append(command, []string{"--tail", strconv.Itoa(tailNumber)}...)
	}
	cmd := exec.Command("docker", command...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var tailNumber int

func init() {
	logCmd.PersistentFlags().IntVar(&tailNumber, "tail", 0, "tail number")
	rootCmd.AddCommand(logCmd)
}
