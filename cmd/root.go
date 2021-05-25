/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pyama86/pdr/pdr"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var config pdr.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pdr",
	Short: "docker command wrapper",
	Long:  `This is docker command wrapper CLI.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", filepath.Join(home, ".pdr"), "config file (default is $HOME/.pdr)")
}

func initConfig() {
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".pdr")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
	config.ReplacePath()
}
