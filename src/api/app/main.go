package main

import (
	"os"
	"task_manager/src/api"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	os.Setenv("APPLICATION_TYPE", "API")
	api := api.NewAPI(&api.Options{})
	api.Serve()
}
