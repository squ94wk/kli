package main

import (
	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/command"
	"github.com/squ94wk/kli/internal/config"
)

func main() {
	conf := config.Configure()

	commander := command.Init()
	cli := cli.GetCLI(conf)
	cmd := commander.GetCommand(conf)
	cmd.Run(conf, cli)
}
