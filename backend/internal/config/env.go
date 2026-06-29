package config
import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	DatabaseURL string
	JWTSecret string
	Port string
}

func Load() (*Config , error ) {
	 err := godotenv.Load()
	 if err!=nil{
		return nil,err
	 }
	 cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port: os.Getenv("PORT"),
	 }
	 return cfg,nil
}
