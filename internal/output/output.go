// This package defines possible OutputModules for the CLI
// Those are:
//   - commandline
//   - file
package output

import "github.com/squ94wk/kli/internal/config"

type Module interface {
}

func GetModule(conf config.Config) Module {
	return Output{}
}

type Output struct {}
