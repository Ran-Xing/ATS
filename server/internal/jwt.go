package internal

import (
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var (
	jwtSecret        = []byte(viper.GetString("jwtSecret"))
	signedString     string
	decryptionString *jwt.Token
)

type JWTClaims struct {
	UserName string `json:"username,omitempty"`
	jwt.StandardClaims
}

func JwtEncryption(userName string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":      time.Now().Unix() - 5,
		"exp":      time.Now().Unix() + 60*60*48,
		"iss":      "nft-manager",
		"username": userName,
	})
	// Sign and get the complete encoded token as a string using the secret
	if signedString, err = t.SignedString(jwtSecret); err != nil {
		log.Errorf("JWT signed string error: %v", err)
		return "", err
	}
	return signedString, nil
}

func JwtDecryption(signedString string) (string, bool) {
	if decryptionString, err = jwt.ParseWithClaims(signedString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}); err != nil {
		log.Errorf("JWT decryption error: %v", err)
		return "", false
	}
	if !decryptionString.Valid {
		return "", false
	}
	return decryptionString.Claims.(*JWTClaims).UserName, true
}
