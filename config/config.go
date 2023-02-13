package config

type Config struct {
	Path string

	UserFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/customer.json"

	return cfg
}
