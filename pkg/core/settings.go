package core

import (
	"os"
)

func SetEnv() {
	os.Setenv("DB_DIALECT", "postgres")
	os.Setenv("DB_NAME", "senao")
	os.Setenv("DB_SCHEMA", "senao")
	os.Setenv("DB_USER", "usr")
	os.Setenv("DB_PASSWORD", "usr")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5440")
}
