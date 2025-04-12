package middlewares

import (
	"context"
	"go-boilerplate/configs"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/services"
	"go-boilerplate/internal/utils"
	"net/http"
)

type contextKey string
const UserClaimsKey contextKey = "userClaims"

func AuthMiddleware(requiredRole string, authService services.AuthServiceInterface, sessionRepo repositories.SessionRepositoryInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			tokenString, err := utils.GetBearerToken(req.Header.Get("Authorization"))
			if err != nil {
				utils.SendResponse(res, "Invalid token format", http.StatusUnauthorized, nil)
				return
			}

			isBlacklisted, err := sessionRepo.IsTokenBlacklisted(req.Context(), tokenString)
			if err != nil {
				utils.SendResponse(res, err.Error(), http.StatusInternalServerError, nil)
				return
			}
			if isBlacklisted {
				utils.SendResponse(res, "Unauthorized: Token is blacklisted", http.StatusUnauthorized, nil)
				return
			}

			userClaims, err := utils.ValidateJWT(tokenString, configs.GetENV("JWT_SECRET"))
			if err != nil {
				utils.SendResponse(res, "Unauthorized: Invalid token", http.StatusUnauthorized, nil)
				return
			}

			hasRole := false
			for _, role := range userClaims.Roles {
				if role.Name == requiredRole {
					hasRole = true
					break
				}
			}

			if !hasRole {
				utils.SendResponse(res, "Forbidden: Insufficient permissions", http.StatusForbidden, nil)
				return
			}

			ctx := context.WithValue(req.Context(), UserClaimsKey, userClaims)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}