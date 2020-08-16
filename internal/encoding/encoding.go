package encoding

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/squ94wk/kli/internal/config"
)

type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

func GetEncoder(conf config.Config) Encoder {
	switch conf.Encoding {
	case "pem":
		return PEM{}
	case "der":
		return DER{}
	}
	return PEM{}
}

type PEM struct{}

func (p PEM) Encode(doc interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	switch doc.(type) {
	case *rsa.PrivateKey:
		key := doc.(*rsa.PrivateKey)
		pemBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		}
		err := pem.Encode(buf, pemBlock)
		return buf.Bytes(), err
	case *x509.Certificate:
		cert := doc.(*x509.Certificate)
		pemBlock := &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		}
		err := pem.Encode(buf, pemBlock)
		return buf.Bytes(), err
	default:
		panic(fmt.Sprintf("no PEM encoder for type %T", doc))
	}
}

type DER struct{}

func (d DER) Encode(doc interface{}) ([]byte, error) {
	switch doc.(type) {
	case *rsa.PrivateKey:
		key := doc.(*rsa.PrivateKey)
		der := x509.MarshalPKCS1PrivateKey(key)
		return der, nil
	case *x509.Certificate:
		cert := doc.(*x509.Certificate)
		return cert.Raw, nil
	default:
		panic(fmt.Sprintf("no PEM encoder for type %T", doc))
	}
}
