package config

import "github.com/caarlos0/env/v6"

type SetupServer struct {
	Protocol string `env:"PROTOCOL" envDefault:"tcp"`
	Hostname string `env:"WEB_HOST" envDefault:""`
	Port     string `env:"WEB_PORT" envDefault:"8080"`
}

type Config struct {
	Server *SetupServer
}

func Envread() (*Config, error) {
	server_cfg := &SetupServer{}
	if err := env.Parse(server_cfg); err != nil {
		return nil, err
	}
	return &Config{
		Server: server_cfg,
	}, nil

}
