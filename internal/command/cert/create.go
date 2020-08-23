package cert

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"time"

	"github.com/squ94wk/kli/internal/cli"
	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/signing"
)

type Create struct{}

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
	var key crypto.Signer
	signers, leftover, err := signing.MatchKeys(cli.Inputs.Keys, cli.Inputs.Certs)
	if err != nil {
		log.Fatal(err)
	}

	// choose key
	switch len(leftover) {
	case 0:
		var err error
		key, err = cli.Crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		key = leftover[0].Key
	default:
		log.Fatal("choosing key for new certificate is ambiguous: found 2 or more keys that don't belong to a certificate")
	}

	// create cert
	var cert *x509.Certificate
	var parentCert *x509.Certificate
	switch {
	case len(cli.Inputs.Certs) > 1:
		log.Fatal("cross signing cert is not supported (yet)")
	case len(cli.Inputs.Certs) > len(signers):
		log.Fatal("missing key for some certificate(s)")

	case len(cli.Inputs.Certs) == 0:
		// self sign
		rootTemp := x509.Certificate{
			SignatureAlgorithm: x509.SHA256WithRSAPSS,
			PublicKeyAlgorithm: x509.RSA,
			PublicKey:          key.Public(),
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

		rootCARaw, err := x509.CreateCertificate(cli.Crypto.Rand(), &rootTemp, &rootTemp, key.Public(), key)
		if err != nil {
			log.Fatal(err)
		}

		parentCert, err = x509.ParseCertificate(rootCARaw)
		if err != nil {
			log.Fatal(err)
		}

		certDer, err := x509.CreateCertificate(cli.Crypto.Rand(), parentCert, parentCert, key.Public(), key)
		if err != nil {
			log.Fatal(err)
		}
		cert, err = x509.ParseCertificate(certDer)
		if err != nil {
			log.Fatal(err)
		}

	case len(cli.Inputs.Certs) == 1:
		var signingKey crypto.Signer
		for cert, key := range signers {
			parentCert = cert.Cert
			signingKey = key.Key
		}

		certDer, err := x509.CreateCertificate(cli.Crypto.Rand(), parentCert, parentCert, key.Public(), signingKey)
		if err != nil {
			log.Fatal(err)
		}
		cert, err = x509.ParseCertificate(certDer)
		if err != nil {
			log.Fatal(err)
		}
	}

	buf, err := cli.Encoder.Encode(cert)
	if err != nil {
		log.Fatal(err)
	}

	pretty, err := cli.Format.Format(buf)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Output.WriteType(pretty)
	if err != nil {
		log.Fatal(err)
	}
}
