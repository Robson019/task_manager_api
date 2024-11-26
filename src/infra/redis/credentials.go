package redis

import (
	"fmt"
	"task_manager/src/core/utils"
)

func getAddress() string {
	host := utils.GetenvWithDefault("REDIS_HOST", "redis")
	port := utils.GetenvWithDefault("REDIS_PORT", "6379")

	return fmt.Sprintf("%s:%s", host, port)
}
