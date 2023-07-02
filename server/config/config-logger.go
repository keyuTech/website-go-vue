package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	Show         string `yaml:"show"`
	ShowLine     bool   `yaml:"show-line"`
	LogInConsole bool   `yaml:"log-in-console"`
}
