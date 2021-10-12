package crypto

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
)

/*
   |--------------------------------------------------------------------------
   | RSA Public Key @Access
   |--------------------------------------------------------------------------
*/
func RsaAccessPublicKey() (*rsa.PublicKey, error) {
	publicKey := []byte(ReadFile(os.Getenv("PEM_FILE")))
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	return key, err
}
