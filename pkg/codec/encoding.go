package codec

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

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
