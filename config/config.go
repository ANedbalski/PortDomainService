package config

type Config struct {
	Server struct {
		// HTTP listen addresses
		HTTP struct {
			// public URL
			Service string `mapstructure:"service" json:"service" yaml:"service"`

			// Keep room for other http services.
			//For example Prometheus metrics or gRPC gateway
		} `mapstructure:"http" json:"http" yaml:"http"`
	} `mapstructure:"server" yaml:"server" json:"server"`
}
