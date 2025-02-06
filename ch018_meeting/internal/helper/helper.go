package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	. "github.com/andyron/meeting/internal/define"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func UUID() string {
	return uuid.NewV4().String()
}

func GenerateToken(id uint, name string) (string, error) {
	UserClaim := &UserClaims{
		Id:             id,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(MyKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

func Encode(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func Decode(str string, obj interface{}) {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, obj)
	if err != nil {
		panic(err)
	}
}
