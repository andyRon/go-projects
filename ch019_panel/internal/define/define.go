package define

import "github.com/golang-jwt/jwt/v4"

type SystemConfig struct {
	Port  string `json:"port"`  // 端口
	Entry string `json:"entry"` // 入口地址
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserClaim struct {
	jwt.RegisteredClaims
}

var (
	Key            = []byte("panel")
	PID            int
	PageSize       = 20
	ShellDir       = "./shell"
	LogDir         = "./log"
	DefaultWebDir  = "/home/wwwroot/"
	NginxConfigDir = "/home/nginx/conf/"
)
