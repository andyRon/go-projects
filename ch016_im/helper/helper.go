package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/andyron/go-im/define"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GetMD5 生成md5
func GetMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

var myKey = []byte("im")

// GenerateToken 生成token
func GenerateToken(identity, email string) (string, error) {
	uc := UserClaims{
		Identity:         identity,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// AnalyzeToken 解析token
func AnalyzeToken(tokenStr string) (*UserClaims, error) {
	uc := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenStr, uc, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyze Token Error: %v", err)
	}
	return uc, nil
}

// SendCode 像用户邮箱发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <andyron@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("<h1>您的验证码是：" + code + "</h1></b>")
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "andyron@qq.com", define.MailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

func GetUUID() string {
	return uuid.NewV4().String()
}

// GetCode 生成随机验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < 6; i++ {
		str += strconv.Itoa(rand.Intn(10))
	}
	return str
}
