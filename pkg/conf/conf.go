package conf

import (
	"os"
	"sync"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
)

type Conf struct {
	DbUrl        string       `yaml:"db_url" env:"DB_URL"`
	Tls          TlsConf      `yaml:"tls" embed:"" prefix:"tls-"`
	CustomRoutes CustomRoutes `yaml:"custom_routes" embed:"" prefix:"custom-routes-"`
	FrontEndPath string       `yaml:"front_end_path" env:"FRONT_END_PATH" default:"./frontend/dist"`
}

type CustomRoutes struct {
	Paths []string `yaml:"paths" env:"CUSTOM_ROUTES" default:""`
}

type TlsConf struct {
	CertPath string `yaml:"cert_path" env:"TLS_CERT_PATH"`
	KeyPath  string `yaml:"key_path" env:"TLS_PRIVATE_KEY_PATH"`
	Port     string `yaml:"port" env:"TLS_PORT" default:"0.0.0.0:8443"`
}

var configSingleton Conf
var once sync.Once

func MustGetConfig(configFile ...string) Conf {
	once.Do(func() {
		options := []kong.Option{
			kong.DefaultEnvars(""),
		}

		if len(configFile) > 0 {
			f, err := os.Open(configFile[0])
			if err != nil {
				panic(err)
			}
			defer f.Close()
			options = append(options, kong.Configuration(kongyaml.Loader, configFile[0]))
		}

		parser, err := kong.New(&configSingleton, options...)
		if err != nil {
			panic(err)
		}

		_, err = parser.Parse([]string{})
		if err != nil {
			panic(err)
		}
	})
	return configSingleton
}
