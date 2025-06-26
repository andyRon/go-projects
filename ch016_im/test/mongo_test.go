package test

import (
	"context"
	"fmt"
	"github.com/andyron/go-im/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	// 1. 配置连接参数
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2. 创建客户端
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("连接失败: ", err)
	}
	defer client.Disconnect(ctx) // 程序退出时关闭连接

	// 3. 健康检查
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("心跳检测失败: ", err)
	}
	log.Println("成功连接 MongoDB！")
}

func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("im")
	user := new(models.UserBasic)
	err = db.Collection("user").FindOne(context.Background(), bson.D{}).Decode(user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("user ==> ", user)
}

func TestFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "",
		Password: "",
	}).ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("im")
	cursor, err := db.Collection("user_room").Find(context.Background(), bson.D{})
	urs := make([]*models.UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(models.UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			t.Fatal(err)
		}
		urs = append(urs, ur)
	}
	for _, v := range urs {
		fmt.Println("userRoom ==> ", v)
	}
}
