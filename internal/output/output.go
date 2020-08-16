// This package defines possible OutputModules for the CLI
// Those are:
//   - commandline
//   - file
package output

import (
	"io"
	"os"

	"github.com/squ94wk/kli/internal/config"
	"github.com/squ94wk/kli/internal/format"
)

type Module interface {
	io.Writer
	WriteType(format.Type) error
}

func GetModule(conf config.Config) Module {
	return wrapper{os.Stdout}
}

type wrapper struct{
	io.Writer
}

func (w wrapper) WriteType(any format.Type) error {
	return any.Write(w.Writer)
}