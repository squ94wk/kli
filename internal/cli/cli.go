package cli

import (
	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/crypt"
	"github.com/squ94wk/kli/internal/encoding"
	"github.com/squ94wk/kli/internal/format"
	"github.com/squ94wk/kli/internal/output"
)

type CLI struct {
	Output  output.Module
	Crypto  crypt.Module
	Encoder encoding.Encoder
	Format  format.Module
}

func GetCLI(conf config.Config) *CLI {
	outMod := output.GetModule(conf)
	cryptMod := crypt.GetModule(conf)
	formatMod := format.GetModule(conf)
	encoder := encoding.GetEncoder(conf)

	return &CLI{
		Output:  outMod,
		Crypto:  cryptMod,
		Format:  formatMod,
		Encoder: encoder,
	}
}
