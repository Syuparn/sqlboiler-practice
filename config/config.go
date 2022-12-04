package config

type Config struct {
	DBPort int
}

func NewConfig() *Config {
	// TODO: enable to overridden by envvars
	return &Config{
		DBPort: 3306,
	}
}
