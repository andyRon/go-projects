package service

import (
	"context"
	"github.com/andyron/go-im/define"
	"github.com/andyron/go-im/helper"
	"github.com/andyron/go-im/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	user, err := models.GetUserBasicByAccountPassword(account, helper.GetMD5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(user.Identity, user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user_claims")
	uc := u.(*helper.UserClaims)
	user, err := models.GetUserBasicByIdentity(uc.Identity)
	if err != nil {
		log.Printf("[UserDetail ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": user,
	})
}

func UserQuery(c *gin.Context) {
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	user, err := models.GetUserBasicByAccount(account)
	if err != nil {
		log.Printf("[UserQuery ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	data := UserQueryResult{
		Nickname: user.Nickname,
		Sex:      user.Sex,
		Email:    user.Email,
		Avatar:   user.Avatar,
		IsFriend: false,
	}
	if models.JudgeUserIsFriend(user.Identity, uc.Identity) {
		data.IsFriend = true
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": data,
	})
}
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	cnt, err := models.GetUserBasicCountByEmail(email)
	if err != nil {
		log.Printf("[SendCode ERROR]:%v\n", err)
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱已被注册",
		})
		return
	}

	code := helper.GetCode()
	err = helper.SendCode(email, code)
	if err != nil {
		log.Printf("[SendCode ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	err = models.RDB.Set(context.Background(), define.RegisterPrefix+email, code, time.Second*time.Duration(define.ExpireTime)).Err()
	if err != nil {
		log.Printf("[SendCode ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

func Register(c *gin.Context) {
	code := c.PostForm("code")
	email := c.PostForm("email")
	account := c.PostForm("account")
	password := c.PostForm("password")
	if code == "" || email == "" || account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	// 判断账号是否唯一
	cnt, err := models.GetUserBasicCountByAccount(account)
	if err != nil {
		log.Printf("[Register ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号已被注册",
		})
		return
	}
	// 验证码验证
	r, err := models.RDB.Get(context.Background(), define.RegisterPrefix+email).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	if r != code {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	user := &models.UserBasic{
		Identity:  helper.GetUUID(),
		Account:   account,
		Password:  helper.GetMD5(password),
		Email:     email,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err = models.InsertOneUserBasic(user)
	if err != nil {
		log.Printf("[Register ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	token, err := helper.GenerateToken(user.Identity, user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func AddUser(c *gin.Context) {
	account := c.PostForm("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	user, err := models.GetUserBasicByAccount(account)
	if err != nil {
		log.Printf("[AddUser ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询错误",
		})
		return
	}
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	if models.JudgeUserIsFriend(user.Identity, uc.Identity) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "已经是好友",
		})
		return
	}
	// 保存房间记录
	rb := &models.RoomBasic{
		Identity:     helper.GetUUID(),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		UserIdentity: user.Identity,
	}
	err = models.InsertOneRoomBasic(rb)
	if err != nil {
		log.Printf("[AddUser ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 保存用户与房间的关联记录
	ur := &models.UserRoom{
		UserIdentity: uc.Identity,
		RoomIdentity: rb.Identity,
		RoomType:     1,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	if err = models.InsertOneUserRoom(ur); err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	ur = &models.UserRoom{
		UserIdentity: user.Identity,
		RoomIdentity: rb.Identity,
		RoomType:     1,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	if err = models.InsertOneUserRoom(ur); err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}
func UserDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	// 获取房间Identity
	roomIdentity := models.GetUserRoomIdentity(identity, uc.Identity)
	if roomIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "不为好友关系，无需删除",
		})
		return
	}
	// 删除user_room关联关系
	if err := models.DeleteUserRoom(roomIdentity); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 删除room_basic
	if err := models.DeleteRoomBasic(roomIdentity); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

type UserQueryResult struct {
	Nickname string `json:"nickname"`
	Sex      int    `bson:"sex"`
	Email    string `bson:"email"`
	Avatar   string `bson:"avatar"`
	IsFriend bool   `json:"is_friend"` // 是否是好友 【true-是，false-否】
}
