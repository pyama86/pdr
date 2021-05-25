package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	return runUporDown([]string{"down"})
}

func init() {
	downCmd.PersistentFlags().StringVar(&filterRepo, "repo", "", "target repository(default all)")
	downCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "choose target repository")
	rootCmd.AddCommand(downCmd)
}
