package key

import (
	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
)

func NewCreateCmd() Create {
	return Create{}
}

type Create struct {

}

func (c Create) Match(conf config.Config) bool {
	args := conf.Args
	if len(args) < 2 {
		return false
	}
	if args[0] != "key" && args[1] != "key" {
		return false
	}
	if args[0] != "create" && args[1] != "create" {
		return false
	}
	return true
}

func (c Create) Run(conf config.Config, cli cli.CLI) {
	_, _ = cli.Write([]byte("create key"))
	// create key
}