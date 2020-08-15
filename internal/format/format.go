package format

import "github.com/squ94wk/kli/internal/config"

type Module interface {

}

func GetModule(conf config.Config) Module {
	return Format{}
}

type Format struct {}
