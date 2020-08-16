package command

import (
	"github.com/squ94wk/kli/internal/command/cert"
	"github.com/squ94wk/kli/internal/command/key"
)

func Init() Commander {
	var cmder Commander
	cmder.AddCommand(key.NewCreateCmd())
	cmder.AddCommand(cert.NewCreateCmd())
	cmder.AddCommand(NewInspectCmd())

	return cmder
}
