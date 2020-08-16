package format

import "io"

type Binary []byte

func (b Binary) Write(w io.Writer) error {
	_, err := w.Write(b)
	return err
}
