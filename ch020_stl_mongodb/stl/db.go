package stl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// @param 参数名 数据类型 "参数说明（约束条件、默认值等）"
// @return 返回值类型 "返回值描述"

// 操作数据库MongoDB的函数集

// SaveSTLMongo 将stl文件中的数据存入MongoDB
// @param modelSTL ModelSTL stl三角网格信息
// @param user string 数据库的用户名
// @param pwd string 数据库的密码
// @param ip string 数据库的ip地址
// @param port int 数据库的端口号
// @param database string 数据库名称
// @param collection string 数据库集合名称
// @param timeout int64 超时秒数
// @return err error 错误信息
func SaveSTLMongo(modelSTL ModelSTL, user, pwd, ip string, port int, database, collection string, timeout int64) (err error) {
	// 设置连接MongoDB数据库的用户名、密码、ip地址、端口号
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, ip, port))
	// 连接MongoDB需要传入一个上下文context对象，这里使用context.WithTimeout,超时秒数值为timeout
	// 之所以需要传入context对象，是因为当顶级goroutine退出时，可以方便的终结所有子goroutine
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// 连接到MongoDB，得到数据库客户端对象client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("stl/db.go/SaveSTLMongo, connect to mongodb error:", err)
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("stl/db.go/SaveSTLMongo, ping mongodb fatal error:", err)
		return
	}
	log.Println("stl/db.go/SaveSTLMongo, ping mongodb successfully!")
	// 函数退出时，关闭客户端client
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("stl/db.go/SaveSTLMongo, disconnected to MongoDB error:", err)
			return
		}
	}()

	// 根据输入的database和collection名称,获取存储stl数据的MongoDB集合
	coll := client.Database(database).Collection(collection)
	// 查询条件filterM，用来检索是否已有名为modelSTL.Name的记录即document，若有则删除
	filterM := bson.M{"name": modelSTL.Name}
	docCount, err := coll.CountDocuments(ctx, filterM)
	// 删除全部的"name"字段为modelSTL.Name的数据库记录
	if docCount > 0 {
		_, err = coll.DeleteMany(ctx, filterM)
		if err != nil {
			log.Println("stl/db.go/SaveSTLMongo, delete documents error:", err)
			return
		}
	}
	// 插入modelSTL结构体数据，会根据ModelSTL数据类型的json tag值来设置相应document的字段
	_, err = coll.InsertOne(ctx, modelSTL)
	if err != nil {
		log.Println("stl/db.go/SaveSTLMongo, insert document error:", err)
		return
	}
	return
}

// QuerySTLMongo 查询
// @param stlName string stl文件的名称
// @param user string 数据库的用户名
// @param pwd string 数据库的密码
func QuerySTLMongo(stlName, user, pwd, ip string, port int, database, collection string, timeout int64) (modelSTL ModelSTL, err error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, ip, port))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("stl/db.go/QuerySTLMongo, connect to mongodb error:", err)
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("stl/db.go/QuerySTLMongo, ping mongodb fatal error:", err)
		return
	}
	log.Println("stl/db.go/QuerySTLMongo, ping mongodb successfully!")
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("stl/db.go/QuerySTLMongo, disconnected to MongoDB error:", err)
			return
		}
	}()

	coll := client.Database(database).Collection(collection)
	filterM := bson.M{"name": stlName}
	err = coll.FindOne(ctx, filterM).Decode(&modelSTL)
	if err != nil {
		log.Printf("stl/db.go/QuerySTLMongo, find and decode modelSTL(name=%s) error:%s\n", stlName, err)
		return
	}
	return
}
