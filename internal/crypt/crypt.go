// package crypt defines the Module for creating keys
package crypt

import (
	"crypto/rsa"
	"math/rand"

	"github.com/squ94wk/kli/internal/config"
)

type Module interface {
	GenerateKey() (*rsa.PrivateKey, error)
	Rand() *rand.Rand
}

func GetModule(conf config.Config) Module {
	if conf.Algorithm == "rsa" {
		return RSA{
			rng:    rand.New(rand.NewSource(0)),
			length: 4096,
		}
	}
	panic("invalid algorithm")
}

type RSA struct {
	rng    *rand.Rand
	length int
}

func (r RSA) GenerateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(r.rng, r.length)
}

func (r RSA) Rand() *rand.Rand {
	return r.rng
}
