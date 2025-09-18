package bot

type Config struct {
	Tocken      string `yaml:"tocken" validate:"required"`
	AuthAddress string `yaml:"auth_address" validate:"required"`
}
