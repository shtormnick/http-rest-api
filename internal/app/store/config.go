package store

type Config struct {
	DisableURL string `toml:"database_url"`
}

func NewConfig() *Config{
	return &Config{}
}