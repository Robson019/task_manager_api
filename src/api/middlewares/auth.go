package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"task_manager/src/api/dicontainer"
	"task_manager/src/api/handlers/dto/response"
	"task_manager/src/core/domain/authorization"
	"task_manager/src/core/messages"
	"task_manager/src/core/utils"
)

var logger = Logger()

var (
	unauthorizedError = response.ErrorMessage{
		StatusCode: http.StatusUnauthorized,
		Message:    messages.UnauthorizedErrorMessage,
	}
	forbiddenError = response.ErrorMessage{
		StatusCode: http.StatusForbidden,
		Message:    messages.ForbiddenErrorMessage,
	}
	sessionExpiredError = response.ErrorMessage{
		StatusCode: http.StatusUnauthorized,
		Message:    messages.InvalidSessionErrorMessage,
	}
)

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	authModel := os.Getenv("SERVER_CASBIN_AUTH_MODEL")
	authPolicy := os.Getenv("SERVER_CASBIN_AUTH_POLICY")
	enforcer, err := casbin.NewEnforcer(authModel, authPolicy)
	if err != nil {
		fmt.Println("Error when building enforcer:", err)
		log.Fatal().Err(err)
	}
	authService := dicontainer.GetAuthServices()
	return func(context echo.Context) error {
		authHeader := context.Request().Header.Get("Authorization")
		method := context.Request().Method
		path := context.Request().URL.Path
		role, ok := utils.ExtractAuthorizationAccountRole(authHeader)
		if !ok {
			return context.JSON(unauthorizedError.StatusCode, unauthorizedError)
		} else if ok, err = enforcer.Enforce(role, path, method); err != nil {
			return context.NoContent(http.StatusInternalServerError)
		} else if role == authorization.AnonymousRoleCode && !ok {
			return context.NoContent(http.StatusUnauthorized)
		} else if !ok {
			logger.Warn().Fields(map[string]interface{}{
				"path":   path,
				"method": method,
				"role":   role,
			}).Msg("FORBIDDEN ACCESS")
			return context.JSON(forbiddenError.StatusCode, forbiddenError)
		} else if role != authorization.AnonymousRoleCode {
			_, authToken := utils.ExtractToken(authHeader)
			if claims, err := utils.ExtractTokenClaims(authToken); err != nil {
				return context.NoContent(http.StatusUnauthorized)
			} else if uID, err := uuid.Parse(claims.AccountID); err != nil {
				return context.NoContent(http.StatusUnauthorized)
			} else if exists, err := authService.SessionExists(uID, authToken); err != nil {
				return context.JSON(sessionExpiredError.StatusCode, sessionExpiredError)
			} else if !exists {
				return context.JSON(unauthorizedError.StatusCode, unauthorizedError)
			}
		}
		return next(context)
	}
}
