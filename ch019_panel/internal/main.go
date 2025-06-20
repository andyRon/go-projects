package main

import (
	"context"
	"fmt"
	"github.com/andyron/panel/define"
	"github.com/andyron/panel/models"
	"github.com/andyron/panel/router"
	"github.com/andyron/panel/service"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	define.PID = syscall.Getpid()
	fmt.Println("PID : " + strconv.Itoa(define.PID))
	models.NewDB()
	sc := service.GetSystemConfig()
	user := service.InitUserConfig()
	fmt.Println("Address : http://localhost" + sc.Port + sc.Entry)
	fmt.Println("Username : " + user.Name)
	fmt.Println("Password : " + user.Password)

	cron := make(chan struct{})
	go service.Cron(cron)

	e := echo.New()
	v1 := e.Group(sc.Entry)
	router.Router(v1)

	run := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT)
		select {
		case <-ch:
			timeout := 10 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			e.Shutdown(ctx)
			cron <- struct{}{}
			go main()
		}
	}()
	e.Logger.Print(e.Start(sc.Port))
	<-run
}
