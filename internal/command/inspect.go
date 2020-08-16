package command

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
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

	obj, err := ParseAny(content)
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

func ParseAny(bytes []byte) (interface{}, error) {
	// PEM
	if pemBlock, _ := pem.Decode(bytes); pemBlock != nil {
		typ, bytes := pemBlock.Type, pemBlock.Bytes
		switch typ {
		case "CERTIFICATE", "X509 CERTIFICATE":
			return x509.ParseCertificate(bytes)
		case "RSA PRIVATE KEY":
			return x509.ParsePKCS1PrivateKey(bytes)
		case "PRIVATE KEY":
			return x509.ParsePKCS8PrivateKey(bytes)
		}
	}

	// DER
	obj := ParseAnyDer(bytes)
	if obj != nil {
		return obj, nil
	}

	return nil, fmt.Errorf("failed to find familiar object")
}

func ParseAnyDer(bytes []byte) interface{} {
	if key, err := x509.ParsePKCS1PrivateKey(bytes); err == nil {
		return key
	}
	if key, err := x509.ParsePKCS8PrivateKey(bytes); err == nil {
		return key
	}
	if key, err := x509.ParseECPrivateKey(bytes); err == nil {
		return key
	}
	if key, err := x509.ParsePKCS1PublicKey(bytes); err == nil {
		return key
	}
	if key, err := x509.ParsePKIXPublicKey(bytes); err == nil {
		return key
	}
	if cert, err := x509.ParseCertificate(bytes); err == nil {
		return cert
	}

	return nil
}