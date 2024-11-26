package postgres

import (
	"fmt"
	"os"
	"task_manager/src/core/utils"
)

func getDatabaseURI() string {
	schema := utils.GetenvWithDefault("DATABASE_SCHEMA", "postgres")
	user := os.Getenv("DATABASE_USER")
	pwd := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	authentication := fmt.Sprintf("%s:%s", user, pwd)
	dst := fmt.Sprintf("%s:%s/%s", host, port, name)
	sslMode := utils.GetenvWithDefault("DATABASE_SSL_MODE", "disable")
	return fmt.Sprintf("%s://%s@%s?sslmode=%s", schema, authentication, dst, sslMode)
}
