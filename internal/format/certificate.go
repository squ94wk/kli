package format

import (
	"crypto/x509"
	"fmt"
	"io"
)

type Certificate x509.Certificate

func (c Certificate) Short(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(
		"x.509 Certificate",
	)))
	return err
}

func (c Certificate) Detail(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(
		`
Type: Certificate
Format: x.509
Subject: %s
Issuer: %s
`,
		c.Issuer.String(),
		c.Subject.String(),
	)))
	return err
}

func (c Certificate) Write(w io.Writer) error {
	return c.Detail(w)
}
