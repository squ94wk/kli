// package crypt defines the Module for creating keys
package crypt

import (
	"crypto/rsa"
	"math/rand"

	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/encoding"
)

type Module interface {
	encoding.RSAEncoder
	GenerateKey() (*rsa.PrivateKey, error)
}

func GetModule(conf config.Config) Module {
	return RSA{
		rng:    rand.New(rand.NewSource(0)),
		length: 4096,
		enc:    encoding.GetRSAEncoder(conf),
	}
}

type RSA struct {
	rng    *rand.Rand
	length int
	enc    encoding.RSAEncoder
}

func (r RSA) GenerateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(r.rng, r.length)
}

func (r RSA) EncodePrivateKey(key *rsa.PrivateKey) ([]byte, error) {
	return r.enc.EncodePrivateKey(key)
}
