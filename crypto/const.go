package crypto

import "github.com/golang-jwt/jwt"

/*
   |--------------------------------------------------------------------------
   | RSA
   |--------------------------------------------------------------------------
*/
// var SigningMethodName = jwt.SigningMethodRS512.Name
// var SigningMethod = jwt.SigningMethodRS512
// var AccessPublicKey = RsaAccessPublicKey

/*
   |--------------------------------------------------------------------------
   | ECDSA
   |--------------------------------------------------------------------------
*/
var SigningMethodName = jwt.SigningMethodES256.Name
var SigningMethod = jwt.SigningMethodES256
var AccessPublicKey = EcdsaAccessPublicKey
