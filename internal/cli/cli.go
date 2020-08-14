package cli

import (
	"io"

	"github.com/squ94wk/kli/internal/config"
)

type CLI interface {
	io.ReadWriter
}

func GetCLI(conf config.Config) CLI {
	return &CommandLine{}
}

