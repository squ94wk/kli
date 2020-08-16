// This package defines possible OutputModules for the CLI
// Those are:
//   - commandline
//   - file
package output

import (
	"io"
	"os"

	"github.com/squ94wk/kli/internal/config"
)

func GetModule(conf config.Config) io.Writer {
	return os.Stdout
}
