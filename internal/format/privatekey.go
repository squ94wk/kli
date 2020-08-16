package format

import (
	"crypto/rsa"
	"fmt"
	"io"
)

type RSAPrivateKey rsa.PrivateKey

func (k RSAPrivateKey) Short(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(
		"RSA %dbit Private Key",
		k.Size(),
	)))
	return err
}

func (k RSAPrivateKey) Detail(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(
		`
Type: Private Key
Algorithm: RSA
Size: %dbit
`,
		k.Size(),
	)))
	return err
}

func (k RSAPrivateKey) Write(w io.Writer) error {
	return k.Detail(w)
}