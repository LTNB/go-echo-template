package config

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/**
 * @author LTNB (baolam0307@gmail.com)
 * @since
 *
 */
type Claims struct {
	Data map[string]interface{}
	jwt.StandardClaims
}

var jwtConf *JWTConf

type JWTConf struct {
	expireTime time.Duration
	secret []byte
}

var jwtKey = []byte("my_secret_key")

func InitJWTConf(expireTime time.Duration, secret string){
	jwtConf = &JWTConf{
		expireTime: expireTime,
		secret:     []byte(secret),
	}
}
func GenerateToken(data map[string]interface{}) (string, error) {
	expirationTime := time.Now().Add(jwtConf.expireTime)

	claims := &Claims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtConf.secret)
	return tokenString, err
}

func ParseToken(tokenStr string) map[string]interface{} {
	claims := &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtConf.secret, nil
	})

	if err != nil || !token.Valid {
		return nil
	}
	return claims.Data
}

func IsValid(hash, pass string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) != nil {
		return false
	}
	return true
}

func HashString(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	hashByte, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashByte), err
}
