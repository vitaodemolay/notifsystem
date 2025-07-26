package entrypoint

import (
	"context"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

type claimsKey string

var claimEmail claimsKey = "email"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if len(accessToken) == 0 {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not have an access token"})
			return
		}

		accessToken = strings.Replace(accessToken, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8070/realms/dummy-provider")
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "failed to connect to OIDC provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "notifysystem"})
		_, err = verifier.Verify(r.Context(), accessToken)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "invalid access token"})
			return
		}

		token, _ := jwtgo.Parse(accessToken, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		email := claims["email"]

		ctx := context.WithValue(r.Context(), claimEmail, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
