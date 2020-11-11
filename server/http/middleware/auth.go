package middleware

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"net/http"
	"strings"
)

const userKey = "user"

type permissionResult int

const (
	allowed permissionResult = iota
	forbidden
	unauthorized
)

func contains(permissions []common.Permission, permission common.Permission) bool {
	for _, perm := range permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

func getUserPermission(user *domain.User) common.Permission {
	var userPerm = common.Public
	if user != nil {
		if user.IsAdmin {
			userPerm = common.PermitAdmin
		} else {
			userPerm = common.PermitUser
		}
	}
	return userPerm
}

func checkPermission(request *http.Request, userService domain.UserService) (permissionResult, *http.Request, error) {
	logger := common.GetRequestLogger()
	accessToken := request.Header.Get("Authorization")

	fullPath := request.URL.Path

	if strings.HasSuffix(fullPath, "/") {
		fullPath = fullPath[:len(fullPath)-1]
	}

	var permissions []common.Permission
	for key, perm := range common.Permissions {
		pattern, ok := common.PathPatterns[key]
		if !ok {
			logger.Printf(request, "path pattern not found.")
		}
		// logger.Printf(request, "key - %s, fullPath - %s, %v\n", key, fullPath, pattern.MatchString(fullPath))
		if pattern.MatchString(fullPath) {
			permissions = perm[request.Method]
			break
		}
	}

	logger.Printf(request, "permission - %v\n", permissions)
	if contains(permissions, common.Public) {
		return allowed, request, nil
	}

	user, err := userService.GetByAccessToken(request.Context(), accessToken)

	if err != nil {
		return unauthorized, request, common.NewUnauthorizedError("invalid token", err)
	} else {
		logger.Printf(request, "user - %v(%v)\n", user.Email, user.ID)
		var userPerm = getUserPermission(user)
		req := request.WithContext(context.WithValue(request.Context(), userKey, user))
		logger.Printf(request, "permissions - %v, userPerm - %v\n", permissions, userPerm)
		if contains(permissions, userPerm) {
			return allowed, req, nil
		} else {
			return forbidden, req, common.NewForbiddenError("forbidden", err)
		}
	}
}

func GetUser(request *http.Request) *domain.User {
	user := request.Context().Value(userKey)
	if user != nil {
		return user.(*domain.User)
	} else {
		return nil
	}
}

func AuthMiddleware(userService domain.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			logger := common.GetRequestLogger()

			accessToken := request.Header.Get("Authorization")
			// req := request.WithContext(context.WithValue(request.Context(), databaseKey, db))

			result, req, err := checkPermission(request, userService)
			logger.Printf(request, "check permission - %v, (accessToken: %s)", result, accessToken)
			switch result {
			case allowed:
				next.ServeHTTP(writer, req)
			case forbidden:
				logger.Printf(request, "forbidden(accessToken: %s)", accessToken)
				common.WriteError(writer, err, nil)
			case unauthorized:
				logger.Printf(request, "unauthorized(accessToken: %s)", accessToken)
				common.WriteError(writer, err, nil)
			}
		})
	}
}
