package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// HealthChecker 检测网站是否能访问

func main() {
	app := &cli.App{
		Name:  "HealthChecker",
		Usage: "A tiny tool that checks whether a website is runnig or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/*
$ go run . --domain pexels.com
[UP] pexels.com is reachable,
From 192.168.18.218:54779
 To: 104.18.66.220:80
$ go run . --domain baidu.com
[UP] baidu.com is reachable,
From 192.168.18.218:54825
 To: 110.242.68.66:80
$ go run . --d amazon.com -p 8080
[DOWN] amazon.com is unreachable,
 Error: dial tcp 52.94.236.248:8080: connect: connection refused



*/
