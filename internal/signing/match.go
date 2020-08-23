package signing

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/pkg/codec"
)

type CertSigner struct {
	Cert *x509.Certificate
	Key  *rsa.PrivateKey
}

func MatchKeys(conf config.Config) (map[*x509.Certificate]*rsa.PrivateKey, []*rsa.PrivateKey, error) {
	var certs []*x509.Certificate
	var keys []*rsa.PrivateKey
	for _, certFile := range conf.CA {
		cert, err := parseCert(certFile)
		if err != nil {
			return nil, nil, err
		}
		_, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, nil, fmt.Errorf("public key format not (yet) supported")
		}
		certs = append(certs, cert)
	}
	for _, keyFile := range conf.Keys {
		key, err := parseKey(keyFile)
		if err != nil {
			return nil, nil, err
		}
		keys = append(keys, key)
	}

	pairs := make(map[*x509.Certificate]*rsa.PrivateKey, 0)
	leftover := make([]*rsa.PrivateKey, 0)
	for _, key := range keys {
		for _, cert := range certs {
			pub := cert.PublicKey.(*rsa.PublicKey)
			if key.N.Cmp(pub.N) == 0 && key.E == pub.E {
				pairs[cert] = key
				break
			}
		}
		leftover = append(leftover, key)
	}
	return pairs, leftover, nil
}

func parseKey(file string) (*rsa.PrivateKey, error) {
	pem, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	any, err := codec.ParseAny(pem)
	if err != nil {
		return nil, err
	}
	key, ok := any.(*rsa.PrivateKey)
	if !ok {
		log.Fatal(fmt.Errorf("parent key is not a valid private key"))
	}
	return key, err
}

func parseCert(file string) (*x509.Certificate, error) {
	pem, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	any, err := codec.ParseAny(pem)
	if err != nil {
		return nil, err
	}
	cert, ok := any.(*x509.Certificate)
	if !ok {
		log.Fatal(fmt.Errorf("parent is not a valid cert"))
	}
	return cert, err
}
