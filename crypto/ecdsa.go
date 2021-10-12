package crypto

import (
	"crypto/ecdsa"
	"os"

	"github.com/golang-jwt/jwt"
)

/*
   |--------------------------------------------------------------------------
   | ECDSA Public Key @Access
   |--------------------------------------------------------------------------
*/
func EcdsaAccessPublicKey() (*ecdsa.PublicKey, error) {
	publicKey := []byte(ReadFile(os.Getenv("PEM_FILE")))
	key, err := jwt.ParseECPublicKeyFromPEM(publicKey)

	return key, err
}
