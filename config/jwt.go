package config

import (
	"fmt"
	"time"
)

type JWTConfig struct {
	Secret   string
	Issuer   string
	Audience string
	TTL      time.Duration
	Alg      string
}

// สร้างค่า default หรือ validate config
func (c *JWTConfig) BuildJWTInfo() string {
	return fmt.Sprintf("alg=%s issuer=%s audience=%s ttl=%s",
		c.Alg, c.Issuer, c.Audience, c.TTL.String(),
	)
}
