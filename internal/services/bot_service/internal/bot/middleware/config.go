package middleware

type Config struct {
	AuthURL string `yaml:"auth_url" validate:"required"`
}
