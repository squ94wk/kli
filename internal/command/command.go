package command

import "github.com/squ94wk/kli/internal/command/key"

func Init() Commander {
	var cmder Commander
	cmder.AddCommand(key.NewCreateCmd())

	return cmder
}

