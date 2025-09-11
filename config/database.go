package config

import "fmt"

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	DSN      string
}

func (c *DatabaseConfig) BuildDSN() {
	c.DSN = fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Bangkok",
		c.Host, c.User, c.Password, c.Name, c.Port, c.SSLMode,
	)
}
