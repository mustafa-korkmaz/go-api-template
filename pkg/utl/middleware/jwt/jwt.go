package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"

	jwt "github.com/dgrijalva/jwt-go"
)

// New generates new JWT service necessery for auth middleware
func New(secret, algo string, d int) *Service {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		key:      []byte(secret),
		algo:     signingMethod,
		duration: time.Duration(d) * time.Minute,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	duration time.Duration

	// Duration for which the jwt refresh token is valid.
	refreshDuration time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod
}

// MWFunc makes JWT implement the Middleware interface.
func (j *Service) MWFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := j.ParseToken(c)
			if token == nil || !token.Valid {
				return c.NoContent(http.StatusUnauthorized)
			}

			claims := token.Claims.(jwt.MapClaims)

			id := claims["id"].(string)
			username := claims["u"].(string)
			email := claims["e"].(string)
			accessLevel := claims["al"].(int)

			c.Set("id", id)
			c.Set("username", username)
			c.Set("email", email)
			c.Set("access_level", accessLevel)

			return next(c)
		}
	}
}

// ParseToken parses token from Authorization header
func (j *Service) ParseToken(c echo.Context) *jwt.Token {

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil
	}

	parsedToken, _ := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, nil
		}
		return j.key, nil
	})

	return parsedToken
}

// GenerateTokens generates new JWT token and refresh token and populates them with user data
func (j *Service) GenerateTokens(u *model.User) (string, string, error) {
	expire := time.Now().Add(j.duration)

	token := jwt.NewWithClaims((j.algo), jwt.MapClaims{
		"id":  u.ID,
		"u":   u.Username,
		"e":   u.Email,
		"al":  u.AccessLevel,
		"exp": expire.Unix(),
	})

	tokenString, err := token.SignedString(j.key)

	if err != nil {
		return "", "", err
	}

	expire = time.Now().Add(j.refreshDuration)

	refreshToken := jwt.NewWithClaims((j.algo), jwt.MapClaims{
		"id":  u.ID,
		"al":  0,
		"exp": expire.Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(j.key)

	//return tokenString, expire.Format(time.RFC3339), err
	return tokenString, refreshTokenString, err
}
