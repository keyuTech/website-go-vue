package config

type Config struct {
	PostgreSQL PostgreSQL `yaml:"postgresql"`
	Logger     Logger     `yaml:"logger"`
	System     System     `yaml:"system"`
}
