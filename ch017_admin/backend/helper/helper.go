package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/andyron/mini-admin/define"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

// MD5 生成md5
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

var myKey = []byte("im")

// GenerateToken 生成token
func GenerateToken(id uint, identity, name, roleIdentity string, expireAt int64) (string, error) {
	uc := UserClaim{
		Id:           id,
		Identity:     identity,
		IsAdmin:      false,
		Name:         name,
		RoleIdentity: roleIdentity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenStr, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil

}

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

func UUID() string {
	return uuid.NewV4().String()
}

func httpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error
	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 处理 header
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)
		if err != nil {
			return nil, err
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header)
}

func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header)
}

func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header)
}

func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header)
}

// RFC3339ToNormalTime RFC3339 日期格式标准化
func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}
