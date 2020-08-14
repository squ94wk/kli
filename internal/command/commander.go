package command

import (
	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
)

type Commander struct {
	cmds []Command
}

type Command interface {
	Match(config.Config) bool
	Run(config.Config, cli.CLI)
}

func (c *Commander) GetCommand(conf config.Config) Command {
	for _, cmd := range c.cmds {
		if cmd.Match(conf) {
			return cmd
		}
	}
	panic("no command matches")
}

func (c *Commander) AddCommand(cmd Command) {
	c.cmds = append(c.cmds, cmd)
}
