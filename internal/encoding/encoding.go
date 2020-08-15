package encoding

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/squ94wk/kli/internal/config"
)

type RSAEncoder interface {
	EncodePrivateKey(*rsa.PrivateKey) ([]byte, error)
}

func GetRSAEncoder(conf config.Config) RSAEncoder {
	return PEM{}
}

type PEM struct {}

func (p PEM) EncodePrivateKey(key *rsa.PrivateKey) ([]byte, error) {
	buf := &bytes.Buffer{}
	pemBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err := pem.Encode(buf, pemBlock)
	return buf.Bytes(), err
}

type DER struct {}