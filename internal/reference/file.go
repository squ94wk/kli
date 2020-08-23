package reference

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/squ94wk/kli/pkg/codec"
)

type FileResolver struct {
	tryHard bool
}

func (r FileResolver) ResolveKey(ref string) (*Key, error) {
	ext := path.Ext(ref)

	if ext == "" && !r.tryHard {
		return nil, nil
	}

	content, err := ioutil.ReadFile(ref)
	if err != nil {
		return nil, nil
	}

	any, err := codec.ParseAny(content)
	key, ok := any.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("found file '%s', but content is no valid rsa private key", ref)
	}

	return &Key{
		Name: ref,
		Key:  key,
	}, err
}

func (r FileResolver) ResolveCert(ref string) (*Cert, error) {
	ext := path.Ext(ref)

	if ext == "" && !r.tryHard {
		return nil, nil
	}

	content, err := ioutil.ReadFile(ref)
	if err != nil {
		return nil, nil
	}

	any, err := codec.ParseAny(content)
	cert, ok := any.(*x509.Certificate)
	if !ok {
		return nil, fmt.Errorf("found file '%s', but content is no valid x509 certificate", ref)
	}

	return &Cert{
		Name: ref,
		Cert: cert,
	}, err
}
