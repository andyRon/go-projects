package service

import (
	"github.com/andyron/panel/define"
	"github.com/andyron/panel/helper"
	"github.com/andyron/panel/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func SoftList(c echo.Context) error {
	var (
		index, _ = strconv.Atoi(c.QueryParam("index"))
		size, _  = strconv.Atoi(c.QueryParam("size"))
		s        = make([]*models.Soft, 0)
		cnt      int64
	)

	size = helper.If(size == 0, define.PageSize, size).(int)
	index = helper.If(index == 0, 1, index).(int)

	err := models.DB.Model(new(models.Soft)).Count(&cnt).Offset((1 - index) * size).Limit(size).Find(&s).Error
	if err != nil {
		log.Println("[DB ERROR]:", err)
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "success",
		"data": echo.Map{
			"list":  s,
			"count": cnt,
		},
	})
}

func SoftOperation(c echo.Context) error {
	var (
		op        = c.FormValue("op")
		id        = c.FormValue("id")
		s         = new(models.Soft)
		shellPath string
		logPath   = define.LogDir + "/" + helper.UUID() + ".log"
	)
	if op == "" || id == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "参数不能为空",
		})
	}
	err := models.DB.Model(new(models.Soft)).Where("id = ?", id).First(s).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "未找到该软件",
			})
		}
		log.Println("[DB ERROR]:", err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
	}

	switch op {
	case "start":
		shellPath = s.ShellStart
	case "stop":
		shellPath = s.ShellStop
	case "restart":
		shellPath = s.ShellRestart
	case "install":
		shellPath = s.ShellInstall
	case "uninstall":
		shellPath = s.ShellUninstall
	default:
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "操作类型错误",
		})
	}

	go func() {
		eq := &models.ExecuteQueue{
			SoftId:    int(s.ID),
			ShellPath: shellPath,
			LogPath:   logPath,
			State:     1,
		}
		err = models.DB.Create(eq).Error
		if err != nil {
			log.Fatalln("[DB ERROR]:", err.Error())
		}
		helper.RunShell(shellPath, logPath)
		err = models.DB.Model(new(models.ExecuteQueue)).Where("id = ?", eq.ID).Update("state", 2).Error
		if err != nil {
			log.Println("[DB ERROR]:", err.Error()) // TODO
		}
	}()

	return c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "success",
	})
}
