package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/2marks/go-expense-tracker-api/utils"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

type userKey string

var userWithKey userKey = "user"

func (a *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := getTokenFromHeader(r)
		if authToken == "" {
			utils.WriteErroresponseToJson(w, http.StatusUnauthorized, fmt.Errorf("no authorization sent"))
			return
		}

		userId, err := utils.ValidateAuthToken(authToken)
		if err != nil {
			fmt.Println(err)
			utils.WriteErroresponseToJson(w, http.StatusUnauthorized, fmt.Errorf("invalid authorization token sent"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(
			ctx,
			userWithKey,
			userId,
		)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromHeader(r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization")

	if tokenHeader == "" {
		return ""
	}

	tokenSlice := strings.Split(tokenHeader, " ")

	return tokenSlice[1]
}

func GetUserFromContext(ctx context.Context) (int, error) {
	userId, ok := ctx.Value(userWithKey).(int)

	if !ok {
		return 0, fmt.Errorf("error while fetching user from context")
	}

	return userId, nil
}
