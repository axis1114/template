package config

type Jwt struct {
	Secret  string `mapstructure:"secret"`
	Expires int    `mapstructure:"expires"`
	Issuer  string `mapstructure:"issuer"`
}
