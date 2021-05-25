package pdr

import (
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	Path              string
	UpPreHookCommands []string `mapstructure:"up_prehook_commands"`
}

type Config struct {
	Repos map[string]*Repo
}

func (c *Config) ReplacePath() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, cc := range c.Repos {
		cc.Path = strings.Replace(cc.Path, "~", home, 1)
	}

}
