package command

import (
	"fmt"
	"log"

	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/reference"
)

type Inspect struct {}

func NewInspectCmd() Inspect {
	return Inspect{}
}

func (i Inspect) Match(conf config.Config) bool {
	args := conf.Args
	if len(args) < 2 {
		return false
	}
	if args[0] != "inspect" && args[1] != "inspect" {
		return false
	}
	return true
}

func (i Inspect) Run(conf config.Config, cli *cli.CLI) {
	args := conf.Args[1:]
	if len(args) != 1 {
		log.Fatal("only accept exactly one argument")
	}
	target := args[0]
	obj, err := reference.ResolveAny(target, cli.Resolver)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to resolve reference '%s': %w", target, err))
	}
	if obj == nil {
		log.Fatal(fmt.Errorf("couldn't resolve reference"))
	}

	pretty, err := cli.Format.Format(obj)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Output.WriteType(pretty)
	if err != nil {
		log.Fatal(err)
	}
}

