package config

import flag "github.com/spf13/pflag"

type Config struct {
	Args      []string
	Algorithm string
	Encoding  string

	Type string
}

func Configure() Config {
	var conf Config
	InitDefaults(&conf)
	ParseArgs(&conf)

	return conf
}

func InitDefaults(conf *Config) {
	*conf = Config{
		Algorithm: "rsa",
		Encoding:  "pem",
		Type:      "ca",
	}
}

func ParseArgs(conf *Config) {
	flag.StringVar(&conf.Algorithm, "alg", conf.Algorithm, "Cryptography algorithm [rsa]")
	flag.StringVar(&conf.Encoding, "enc", conf.Encoding, "Key encoding [pem, der]")
	flag.StringVar(&conf.Type, "type", conf.Type, "Certificate type")
	flag.Parse()
	conf.Args = flag.Args()
}
