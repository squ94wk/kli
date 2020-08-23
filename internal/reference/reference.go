package reference

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

type Key struct {
	Name string
	Key  crypto.Signer
}

type Cert struct {
	Name string
	Cert *x509.Certificate
}

func GetResolver() Resolver {
	return ChainedResolver{
		[]Resolver{
			FileResolver{tryHard: false},
			FileResolver{tryHard: true},
		},
	}
}

type Resolver interface {
	ResolveKey(string) (*Key, error)
	ResolveCert(string) (*Cert, error)
}

func (k Key) BelongsToCert(cert *Cert) bool {
	switch pub := k.Key.Public().(type) {
	case *rsa.PublicKey:
		if cert.Cert.PublicKeyAlgorithm != x509.RSA {
			return false
		}
		certPub, ok := cert.Cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			//invalid
			return false
		}
		return pub.N.Cmp(certPub.N) == 0 && pub.E == certPub.E
	default:
		return false
	}
}

type AnyResolver interface {
	ResolveAny(string) (interface{}, error)
}

func ResolveAny(ref string, r Resolver) (interface{}, error) {
	if anyR, ok := r.(AnyResolver); ok {
		return anyR.ResolveAny(ref)
	}

	var errs []error
	key, err := r.ResolveKey(ref)
	if err != nil {
		errs = append(errs, err)
	}

	cert, err := r.ResolveCert(ref)
	if err != nil {
		errs = append(errs, err)
	}

	switch {
	case key != nil && cert != nil:
		return nil, fmt.Errorf("ambiguous reference '%s'", ref)
	case key != nil:
		return key, nil
	case cert != nil:
		return cert, nil
	default:
		if len(errs) != 0 {
			return nil, fmt.Errorf("failed to resolve reference '%s': %v", ref, errs)
		}
		return nil, nil
	}
}