package service

import (
	"encoding/json"
	"fmt"
	"github.com/andyron/panel/define"
	"github.com/andyron/panel/helper"
	"github.com/andyron/panel/models"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"net/http"
	"syscall"
	"time"
)

func GetSystemConfig() *define.SystemConfig {
	sc := new(define.SystemConfig)
	c := new(models.Config)
	dsc := getDefaultSystemConfig()
	dscByte, _ := json.Marshal(dsc)
	err := models.DB.Model(new(models.Config)).Where("key = 'system").
		Attrs(map[string]interface{}{"key": "system", "value": string(dscByte)}).FirstOrCreate(c).Error
	if err != nil {
		panic("[Init System_Config Error] : " + err.Error())
	}
	err = json.Unmarshal([]byte(c.Value), sc)
	if err != nil {
		panic("[Unmarshal Error] : " + err.Error())
	}
	return sc
}

func InitUserConfig() *define.User {
	dUser := getDefaultUser()
	dUserByte, _ := json.Marshal(dUser)
	u := new(define.User)
	c := new(models.Config)
	err := models.DB.Model(new(models.Config)).Where("key ='user").
		Attrs(map[string]interface{}{"key": "user", "value": string(dUserByte)}).FirstOrCreate(c).Error
	if err != nil {
		panic("[Init User Error] : " + err.Error())
	}
	err = json.Unmarshal([]byte(c.Value), u)
	if err != nil {
		panic("[Unmarshal Error] : " + err.Error())
	}
	return u
}

func getDefaultSystemConfig() *define.SystemConfig {
	return &define.SystemConfig{
		Port:  "1888",
		Entry: "/" + helper.RandomString(8),
	}
}

func getDefaultUser() *define.User {
	return &define.User{
		Name:     helper.RandomString(8),
		Password: helper.RandomString(8),
	}
}

func UpdateSystemConfig(c echo.Context) error {
	var (
		port   = c.FormValue("port")
		entry  = c.FormValue("entry")
		config = new(models.Config)
		sc     = new(define.SystemConfig)
	)
	err := models.DB.Where("key ='system").First(config).Error
	if err != nil {
		log.Printf("[DB Error] : %v\n", err)
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
	}
	json.Unmarshal([]byte(config.Value), sc)

	if port != "" {
		sc.Port = ":" + port
	}
	if entry != "" {
		sc.Entry = "/" + entry
	}
	scByte, _ := json.Marshal(sc)

	err = models.DB.Model(new(models.Config)).Where("key ='system").
		Update("value", string(scByte)).Error
	if err != nil {
		log.Printf("[DB Error] : %v\n", err)
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "更新成功",
	})

	// 重启服务
	syscall.Kill(define.PID, syscall.SIGINT)
	return nil
}

func SystemState(c echo.Context) error {
	var (
		cpuUsedPercent  float64
		memUsedPercent  float64
		diskUsed        uint64
		diskUsedPercent float64
	)
	// CPU
	cpuPercent, _ := cpu.Percent(time.Second, true)
	for _, percent := range cpuPercent {
		cpuUsedPercent += percent
	}
	cpuUsedPercent /= float64(len(cpuPercent))
	// 内存
	vms, _ := mem.VirtualMemory()
	memUsedPercent = vms.UsedPercent
	// 磁盘
	partitions, _ := disk.Partitions(true)
	for _, partition := range partitions {
		usage, _ := disk.Usage(partition.Mountpoint)
		diskUsed += usage.Used
	}
	allUsage, _ := disk.Usage("/")
	diskUsedPercent = float64(diskUsed) / float64(allUsage.Total) * 100

	return c.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": echo.Map{
			"cpu_used_percent":  fmt.Sprintf("%.2f", cpuUsedPercent),
			"mem_used_percent":  fmt.Sprintf("%.2f", memUsedPercent),
			"disk_used_percent": fmt.Sprintf("%.2f", diskUsedPercent),
		},
	})
}
