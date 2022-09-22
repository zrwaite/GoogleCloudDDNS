package auth

import (
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zrwaite/google-cloud-ddns/models"
)

type TokenStruct struct {
	ISS   string `json:"iss"`
	Scope string `json:"scope"`
	Aud   string `json:"aud"`
	jwt.StandardClaims
}

func EncodeToken(params *models.Params) (tokenString string, success bool) {
	var err error
	jwtKeyString := params.JWT_KEY
	if jwtKeyString == "" {
		fmt.Println("JWT_KEY not found")
		return "", false
	}
	var jwtKey *rsa.PrivateKey
	jwtKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtKeyString))
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	expirationTime := time.Now().Add(time.Hour)
	tokenBody := TokenStruct{
		ISS:   params.ISS,
		Scope: "https://www.googleapis.com/auth/ndev.clouddns.readwrite",
		Aud:   "https://oauth2.googleapis.com/token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenBody)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Failed to create jwt - " + err.Error())
		return "", false
	}
	return tokenString, true
}
