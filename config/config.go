package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	// Minio    MinioConfig
	// JWT      JWTConfig
}

func LoadConfig() *Config {
	v := viper.New()
	// run main
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".") // หาใน root project
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cfg := &Config{
		Database: DatabaseConfig{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetInt("DB_PORT"),
			User:     v.GetString("DB_USER"),
			Password: v.GetString("DB_PASSWORD"),
			Name:     v.GetString("DB_NAME"),
			SSLMode:  v.GetString("DB_SSLMODE"),
		},
		// Minio: MinioConfig{
		// 	Endpoint:  v.GetString("MINIO_ENDPOINT"),
		// 	AccessKey: v.GetString("MINIO_ACCESS_KEY"),
		// 	SecretKey: v.GetString("MINIO_SECRET_KEY"),
		// 	Bucket:    v.GetString("MINIO_BUCKET"),
		// },
		// JWT: JWTConfig{
		// 	Secret: v.GetString("JWT_SECRET"),
		// },
	}

	// generate DSN
	// cfg.Database.BuildDSN()
	// log.Println("✅ Config loaded", cfg)
	return cfg
}
