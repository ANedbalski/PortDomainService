package config

import "github.com/spf13/viper"

func ApplyDefaults(v *viper.Viper) error {
	v.SetDefault("server.http.service", "0.0.0.0:8282")

	return nil
}
