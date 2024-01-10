package config

type JWTConfig struct {
	SecretKey string `env:"SECRET_KEY, required"`
}
