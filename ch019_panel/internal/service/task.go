package service

import (
	"bufio"
	"errors"
	"github.com/andyron/panel/define"
	"github.com/andyron/panel/helper"
	"github.com/andyron/panel/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"syscall"
)

func TaskDetail(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "参数不能为空",
		})
	}
	data := new(TaskDetailResponse)
	err := models.DB.Model(new(models.Task)).Select("id, name, spec, shell_path data").
		Where("id = ?", id).Find(data).Error
	if err != nil {
		log.Println("[DB ERROR]:", err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}
	b, err := os.ReadFile(data.Data)
	if err != nil {
		log.Println("[DB ERROR]:", err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}
	data.Data = string(b)
	return c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

func TaskList(c echo.Context) error {
	index, _ := strconv.Atoi(c.QueryParam("index"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	tasks := make([]*models.Task, 0)
	var cnt int64

	size = helper.If(size == 0, define.PageSize, size).(int)
	index = helper.If(index == 0, 1, index).(int)

	err := models.DB.Model(new(models.Task)).Count(&cnt).Offset((index - 1) * size).Limit(size).Find(&tasks).Error
	if err != nil {
		log.Println("[DB ERROR]" + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "加载成功",
		"data": echo.Map{
			"list":  tasks,
			"count": cnt,
		},
	})
}

func TaskAdd(c echo.Context) error {
	name := c.FormValue("name")
	spec := c.FormValue("spec")
	data := c.FormValue("data")
	if name == "" || spec == "" || data == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "参数不能为空",
		})
	}
	shellName := helper.UUID()
	task := &models.Task{
		Name:      name,
		Spec:      spec,
		ShellPath: define.ShellDir + "/" + shellName + ".sh",
		LogPath:   define.LogDir + "/" + shellName + ".log",
	}
	err := models.DB.Create(task).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	f, err := os.Create(task.ShellPath)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(task.ShellPath), 0777)
		f, err = os.Create(task.ShellPath)
		if err != nil {
			log.Println("[CREATE FILE ERROR] : " + err.Error())
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "系统异常 : " + err.Error(),
			})
		}
	}
	w := bufio.NewWriter(f)
	w.WriteString(data)
	w.Flush()

	c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "新增成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
	return nil
}

func TaskEdit(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	spec := c.FormValue("spec")
	data := c.FormValue("data")
	if id == "" || name == "" || spec == "" || data == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "参数不能为空",
		})
	}

	task := new(models.Task)
	err := models.DB.Where("id = ?", id).First(task).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	task.Name = name
	task.Spec = spec
	err = models.DB.Where("id = ?", id).Updates(task).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}

	f, err := os.Create(task.ShellPath)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(task.ShellPath), 0777)
		f, err = os.Create(task.ShellPath)
		if err != nil {
			log.Println("[CREATE FILE ERROR] : " + err.Error())
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "系统异常 : " + err.Error(),
			})
		}
	}
	w := bufio.NewWriter(f)
	w.WriteString(data)
	w.Flush()

	c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "编辑成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
	return nil
}

func TaskDelete(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "参数不能为空",
		})
	}
	err := models.DB.Where("id = ?", id).Delete(new(models.Task)).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "删除成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)

	return nil
}
