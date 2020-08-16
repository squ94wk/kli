package format

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
)

type JSON struct {}

func (j JSON) Format(obj interface{}) ([]byte, error){
	switch obj.(type) {
	case *x509.Certificate:
		cert := obj.(*x509.Certificate)
		return json.Marshal(cert)
	default:
		panic(fmt.Sprintf("no json formatter for type %T", obj))
	}
}