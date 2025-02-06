package define

import "github.com/dgrijalva/jwt-go"

var MyKey = []byte("meeting")

type UserClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
