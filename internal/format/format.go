package format

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io"

	"github.com/squ94wk/kli/internal/config"
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
	switch obj.(type) {
	case []byte:
		return Binary(obj.([]byte)), nil
	case *rsa.PrivateKey:
		return RSAPrivateKey(*obj.(*rsa.PrivateKey)), nil
	case *rsa.PublicKey:
		return RSAPublicKey(*obj.(*rsa.PublicKey)), nil
	case *x509.Certificate:
		return Certificate(*obj.(*x509.Certificate)), nil
	default:
		return nil, fmt.Errorf("can't format %T", obj)
	}
}
