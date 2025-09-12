package bot

type Config struct {
	Tocken string `yaml:"tocken" validate:"required"`
}
