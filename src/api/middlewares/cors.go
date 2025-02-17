package middlewares

import (
	"strings"
	"task_manager/src/core/utils"

	"github.com/labstack/echo/v4"
)

// VerifyOrigin verifies if the request origin is included on the defined server
// allowed hosts.
func VerifyOrigin(origin string) (bool, error) {
	allowedOrigins := strings.Split(utils.GetenvWithDefault("SERVER_ALLOWED_HOSTS", "*"), ",")
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == "*" || origin == allowedOrigin {
			return true, nil
		}
	}
	return false, &echo.HTTPError{Code: 401, Message: "you're not allowed to access this API"}
}

// OriginInspectSkipper verifies the request context and skip the origin verification.
// It's useful to allow access for any origin when a route (e.g. public images routes)
// should be accessed by anyone.
func OriginInspectSkipper(context echo.Context) bool {
	return false
}
