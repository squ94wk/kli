package format

import (
	"crypto/rsa"
	"fmt"
	"io"
)

type RSAPublicKey rsa.PublicKey

func (k RSAPublicKey) Short(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(
		"RSA Public Key",
	)))
	return err
}

func (k RSAPublicKey) Detail(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(
		`
Type: Private Key
Algorithm: RSA
`,
	)))
	return err
}

func (k RSAPublicKey) Write(w io.Writer) error {
	return k.Detail(w)
}