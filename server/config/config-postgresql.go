package config

import "strconv"

type PostgreSQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log-level"`
}

func (p *PostgreSQL) Dsn() string {
	return "host=" + p.Host + " user=" + p.User + " password=" + p.Password + " dbname=" + p.DB + " port=" + strconv.Itoa(p.Port) + " sslmode=disable"
}
