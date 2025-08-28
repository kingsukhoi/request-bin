package conf

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Conf struct {
	DbUrl        string  `yaml:"db_url" env:"DB_URL"`
	Tls          TlsConf `yaml:"tls"`
	CustomRoutes CustomRoutes
}

type CustomRoutes struct {
	Paths []string `yaml:"paths" env:"CUSTOM_ROUTES" env-default:""`
}
type TlsConf struct {
	CertPath string `yaml:"cert_path" env:"TLS_CERT_PATH"`
	KeyPath  string `yaml:"key_path" env:"TLS_PRIVATE_KEY_PATH"`
	Port     string `yaml:"port" env:"TLS_PORT" env-default:"0.0.0.0:8443"`
}

var configSingleton Conf
var once sync.Once

func MustGetConfig(configFile ...string) Conf {
	once.Do(func() {
		if len(configFile) > 0 {
			err := cleanenv.ReadConfig(configFile[0], &configSingleton)
			if err != nil {
				panic(err)
			}
		} else {
			err := cleanenv.ReadEnv(&configSingleton)
			if err != nil {
				panic(err)
			}
		}
	})
	return configSingleton
}
