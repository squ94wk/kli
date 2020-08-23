package format

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io"

	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/reference"
)

type Type interface {
	Write(io.Writer) error
}

type Module interface {
	//inspect
	//  -o json
	//  -o yaml
	//  -v
	//  --verbose

	//create
	//  pem
	//  der

	Format(interface{}) (Type, error)
}

func GetModule(conf config.Config) Module {
	return formatter{}
}

type formatter struct{}

func (f formatter) Format(obj interface{}) (Type, error) {
	switch o := obj.(type) {
	case []byte:
		return Binary(o), nil
	case *rsa.PrivateKey:
		return RSAPrivateKey(*o), nil
	case *rsa.PublicKey:
		return RSAPublicKey(*o), nil
	case *x509.Certificate:
		return Certificate(*o), nil
	case *reference.Key:
		return f.Format(o.Key)
	case *reference.Cert:
		return f.Format(o.Cert)
	default:
		return nil, fmt.Errorf("can't format %T", obj)
	}
}
