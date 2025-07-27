package entrypoint

import (
	"context"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

//go:generate go run go.uber.org/mock/mockgen -package=mock -destination=./mock/auth_middleware.go . IdentityProvider
type IdentityProvider interface {
	Auth(next http.Handler) http.Handler
}

type identityProvider struct {
	clientId    string
	redirectUri string
	tokenType   string
}

func NewIdentityProvider(clientId, redirectUri, tokenType string) IdentityProvider {
	return &identityProvider{
		clientId:    clientId,
		redirectUri: redirectUri,
		tokenType:   tokenType,
	}
}

type claimsKey string

var claimEmail claimsKey = "email"

func (i *identityProvider) Auth(next http.Handler) http.Handler {
	adjustedTokenType := i.tokenType + " "
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if len(accessToken) == 0 {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not have an access token"})
			return
		}

		accessToken = strings.Replace(accessToken, adjustedTokenType, "", 1)
		provider, err := oidc.NewProvider(r.Context(), i.redirectUri)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "failed to connect to OIDC provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: i.clientId})
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
