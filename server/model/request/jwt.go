package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID         int
	Mobile     string
	Nickname   string
	BufferTime int64
	jwt.StandardClaims
}
