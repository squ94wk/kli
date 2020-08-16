package cert

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
)

type Create struct {}

func NewCreateCmd() Create {
	return Create{}
}

func (c Create) Match(conf config.Config) bool {
	args := conf.Args
	if len(args) < 2 {
		return false
	}
	if args[0] != "cert" && args[1] != "cert" {
		return false
	}
	if args[0] != "create" && args[1] != "create" {
		return false
	}
	return true
}

func (c Create) Run(conf config.Config, cli *cli.CLI) {
	key, err := cli.Crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	rootTemp := x509.Certificate{
		SignatureAlgorithm: x509.SHA256WithRSAPSS,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          key.PublicKey,
		SerialNumber:       big.NewInt(time.Now().UnixNano()),
		Issuer: pkix.Name{
			CommonName: "Root CA",
		},
		Subject:        pkix.Name{},
		NotBefore:      time.Time{},
		NotAfter:       time.Time{},
		KeyUsage:       x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		IsCA:           true,
		MaxPathLenZero: false,
	}

	rootCARaw, err := x509.CreateCertificate(cli.Crypto.Rand(), &rootTemp, &rootTemp, &key.PublicKey, key)
	if err != nil {
		log.Fatal(err)
	}

	rootCA, err := x509.ParseCertificate(rootCARaw)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := cli.Encoder.Encode(rootCA)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", string(buf))
}