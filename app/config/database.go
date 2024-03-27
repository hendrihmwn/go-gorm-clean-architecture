package config

import (
	"fmt"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/utils"
)

func GetDBConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		utils.Getenv("DATABASE_HOST", "localhost"),
		utils.Getenv("DATABASE_USER", "root"),
		utils.Getenv("DATABASE_PASSWORD", "12345678"),
		utils.Getenv("DATABASE_NAME", "library"),
		utils.Getenv("DATABASE_PORT", "5432"),
		utils.Getenv("DATABASE_SSL_MODE", "disable"),
	)
}
