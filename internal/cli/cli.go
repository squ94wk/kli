package cli

import (
	"fmt"
	"log"

	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/crypt"
	"github.com/squ94wk/kli/internal/encoding"
	"github.com/squ94wk/kli/internal/format"
	"github.com/squ94wk/kli/internal/output"
	"github.com/squ94wk/kli/internal/reference"
)

type CLI struct {
	Output   output.Module
	Crypto   crypt.Module
	Encoder  encoding.Encoder
	Format   format.Module
	Resolver reference.Resolver
	Inputs   Inputs
}

type Inputs struct {
	Keys  []*reference.Key
	Certs []*reference.Cert
}

func GetCLI(conf config.Config) *CLI {
	var cli CLI
	cli.Output = output.GetModule(conf)
	cli.Crypto = crypt.GetModule(conf)
	cli.Format = format.GetModule(conf)
	cli.Encoder = encoding.GetEncoder(conf)
	resolver := reference.GetResolver()
	cli.Resolver = resolver

	err := ResolveInputs(conf, resolver, &cli.Inputs)
	if err != nil {
		log.Fatal(err)
	}

	return &cli
}

func ResolveInputs(conf config.Config, r reference.Resolver, inputs *Inputs) error {
	for _, ref := range conf.Keys {
		key, err := r.ResolveKey(ref)
		if err != nil {
			return fmt.Errorf("failed to resolve key reference '%s'", ref)
		}
		inputs.Keys = append(inputs.Keys, key)
	}
	for _, ref := range conf.Certs {
		cert, err := r.ResolveCert(ref)
		if err != nil {
			return fmt.Errorf("failed to resolve cert reference '%s'", ref)
		}
		inputs.Certs = append(inputs.Certs, cert)
	}
	return nil
}