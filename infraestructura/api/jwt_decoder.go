package api

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
	logger "github.com/fravega/go-logger/v2"
)

type JwtDecoder struct {
	publicKey *rsa.PublicKey
	parser    *jwt.Parser
}

func CreateJwtDecoder(pemFile string) *JwtDecoder {
	pem, err := ioutil.ReadFile(pemFile)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while parsing the pem file")
		return nil
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	if err != nil {
		logger.GetDefaultLogger().Error("Invalid public pem")
		panic(err.Error())
	}
	return &JwtDecoder{publicKey: publicKey, parser: &jwt.Parser{ValidMethods: []string{"RS512"}}}
}

func (d *JwtDecoder) GetTokenAttributes(tokenString string) (string, string, []string, error) {
	token, err := d.decode(tokenString)
	if err != nil {
		return "", "", nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sellerId := claims["sellerId"].(string)
		permissions := getClaims(claims, "permissions")
		email := claims["email"].(string)
		return email, sellerId, permissions, nil
	} else {
		return "", "", nil, fmt.Errorf("Token invalido")
	}
}
func getClaims(claims jwt.MapClaims, key string) []string {
	buffer := []string{}
	if tokenClaims, ok := claims[key]; ok {
		buffer = toStrings(tokenClaims.([]interface{}))
	}
	return buffer
}
func toStrings(i []interface{}) []string {
	s := []string{}
	for _, e := range i {
		s = append(s, e.(string))
	}
	return s
}

func (d *JwtDecoder) decode(tokenString string) (*jwt.Token, error) {
	return d.parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return d.publicKey, nil
	})
}
