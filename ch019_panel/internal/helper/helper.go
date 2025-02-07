package helper

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/andyron/panel/define"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"time"
)

func RandomString(n int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	rand.Seed(time.Now().UnixNano()) // TODO
	ans := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, str[rand.Intn(len(str))])
	}
	return string(ans)
}

func GenerateToken() (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, &define.UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24 * 30),
			},
		},
	})
	token, err := tokenStruct.SignedString(define.Key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) error {
	claims, err := jwt.ParseWithClaims(token, &define.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return define.Key, nil
	})
	if err != nil {
		return err
	}
	if !claims.Valid {
		return errors.New("error Token")
	}
	return nil
}

func If(bo bool, a, b interface{}) interface{} {
	if bo {
		return a
	}
	return b
}

func UUID() string {
	return uuid.NewV4().String()
}

func RunShell(shellPath, logPath string) {
	cmdChmod := exec.Command("sh", "-c", "chmod +x "+shellPath)
	var errChmod bytes.Buffer
	cmdChmod.Stderr = &errChmod
	if err := cmdChmod.Run(); err != nil {
		log.Println("[CHMOD ERROR] : " + err.Error())
	}
	//
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND, 0666)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(logPath), 0777)
		file, err = os.Create(logPath)
		if err != nil {
			log.Fatalln("[CREATE ERROR] : " + err.Error())
		}
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\n")
	writer.Flush()

	//
	cmdShell := exec.Command("sh", "-c", "nohup "+shellPath+" >> "+logPath+" 2>&1 &") // TODO
	var outShell, errShell bytes.Buffer
	cmdShell.Stdout = &outShell
	cmdShell.Stderr = &errShell
	if err := cmdShell.Run(); err != nil {
		log.Println("[SHELL ERROR] : "+err.Error()+" ErrShell : ", errShell.String())
	}
	log.Println(outShell.String())
}
