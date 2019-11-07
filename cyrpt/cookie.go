package cyrpt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"regexp"
	"time"
)

var jwtKey = []byte(os.Getenv("jwt_key"))
var regex = regexp.MustCompile(`(token=)(.*)(;.*)`)

// ƒ responsible for key interpretation
func keyFunc(token *jwt.Token) (interface{}, error) { return jwtKey, nil }

// data structure representing a parsed JWT string.
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type CookieBuilder interface {
	Expiry(time.Time) CookieBuilder
	Email(string) CookieBuilder
	Build() string
}

type cookieBuilder struct {
	email  string
	expiry time.Time
}

func (cb *cookieBuilder) Expiry(expiry time.Time) CookieBuilder {
	cb.expiry = expiry
	return cb
}

func (cb *cookieBuilder) Email(email string) CookieBuilder {
	cb.email = email
	return cb
}

func New() CookieBuilder { return &cookieBuilder{} }

func (cb *cookieBuilder) Build() string {

	if cb.expiry == (time.Time{}) {
		// Declare a default expiration time of the token
		cb.expiry = time.Now().Add(30 * time.Minute)
	}

	// Declare the token with the algorithm used for signing, and JWT claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email: cb.email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: cb.expiry.Unix(),
		},
	})

	// Create the JWT string
	if tokenString, err := token.SignedString(jwtKey); err != nil {
		return err.Error()
	} else {
		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		cookie := &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  cb.expiry,
			HttpOnly: false,
		}
		return cookie.String()
	}
}

func Validate(cookie string) (*Claims, error) {
	c := &Claims{}
	// find the token in the cookie, may not exist.
	token := regex.ReplaceAllString(cookie, `$2`)
	if tkn, err := jwt.ParseWithClaims(token, c, keyFunc); err != nil {
		// either the token expired or the signature doesn't match.
		return c, err
	} else if !tkn.Valid {
		return c, errors.New("invalid token")
	} else {
		return c, nil
	}
}
