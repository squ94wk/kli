package format

import "github.com/squ94wk/kli/internal/config"

type Module interface {
	Format(interface{}) ([]byte, error)
}

func GetModule(conf config.Config) Module {
	return JSON{}
}
