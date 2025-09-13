package server

type Config struct {
	Port string `yaml:"port" validate:"required" env:"PORT" env-default:"8080"`
}
