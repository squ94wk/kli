package config

import "flag"

type Config struct {
	Args []string
}

func Configure() Config {
	var conf Config
	ParseArgs(&conf)

	return conf
}

func ParseArgs(conf *Config) {
	flag.Parse()
	conf.Args = flag.Args()
}