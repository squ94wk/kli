package command

import (
	"io/ioutil"
	"log"

	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/pkg/codec"
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
	content, err := ioutil.ReadFile(target)
	if err != nil {
		log.Fatal(err)
	}

	obj, err := codec.ParseAny(content)
	if err != nil {
		log.Fatal(err)
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

