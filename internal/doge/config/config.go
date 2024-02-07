package config

import (
	"errors"
	"net/netip"

	"my-go/internal/doge/validator"

	"github.com/chelpis/member-center/pkg/config"
	"github.com/chelpis/member-center/pkg/db"
)

var ErrLoadConfig = errors.New("load config error")

type Config struct {
	Database   db.DatabaseConfig `koanf:"database" validate:"required"`
	GrpcServer GrpcServerConfig  `koanf:"grpc_server" validate:"required"`
	HttpServer HttpServerConfig  `koanf:"http_server" validate:"required"`
}

type GrpcServerConfig struct {
	Addr netip.AddrPort `koanf:"addr_port" validate:"required"`
}

type HttpServerConfig struct {
	Addr netip.AddrPort `koanf:"addr_port" validate:"required"`
}

func Load(file string) (c *Config, err error) {
	c = &Config{}
	_, err = config.New(file, "", "__").GetConfig(c)
	if err != nil {
		err = errors.Join(err, ErrLoadConfig)
		return
	}

	err = c.Validate()
	if err != nil {
		err = errors.Join(err, ErrLoadConfig)
		return
	}

	return
}

func (c *Config) Validate() error {
	validator := validator.New()
	err := validator.Validate(c)
	return err
}

func (c *Config) Default() config.ConfigInterface {
	return &Config{
		GrpcServer: GrpcServerConfig{
			Addr: netip.AddrPortFrom(netip.MustParseAddr("127.0.0.1"), 50051),
		},
		HttpServer: HttpServerConfig{
			Addr: netip.AddrPortFrom(netip.MustParseAddr("127.0.0.1"), 8080),
		},
	}
}
