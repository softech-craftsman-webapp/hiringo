package crypto

import "github.com/golang-jwt/jwt"

/*
   |--------------------------------------------------------------------------
   | RSA
   |--------------------------------------------------------------------------
*/
var SigningMethodName = jwt.SigningMethodRS512.Name
var SigningMethod = jwt.SigningMethodRS512
var AccessPublicKey = RsaAccessPublicKey

/*
   |--------------------------------------------------------------------------
   | ECDSA
   |--------------------------------------------------------------------------
*/
// var SigningMethodName = jwt.SigningMethodES512.Name
// var SigningMethod = jwt.SigningMethodES512
// var AccessPublicKey = EcdsaAccessPublicKey
