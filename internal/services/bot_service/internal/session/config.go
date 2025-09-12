package session

type Config struct {
	SessionURL string `yaml:"session_url" validate:"required"`
}
