package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	// 设置连接URI
	clientOptions := options.Client().ApplyURI("ch023_mongodb://localhost:27017")
	// 连接到mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	defer client.Disconnect(context.TODO())

	// 选择数据库和集合
	collection := client.Database("testdb").Collection("users")

	// 设置查询条件
	filter := bson.M{"age": bson.M{"$gt": 20}}

	// 执行查询
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// 遍历查询结果
	var users []User
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// 检查遍历过程中是否发生错误
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// 打印查询结果
	for _, user := range users {
		log.Printf("User: %+v\n", user)
	}

	/* 高级查询技巧 */
	// 使用投影
	opts := options.Find().SetProjection(bson.M{"name": 1, "email": 1, "_id": 0})
	cursor, err = collection.Find(context.TODO(), bson.M{}, opts)

	// 使用排序
	opts = options.Find().SetSort(bson.M{"age": 1})
	cursor, err = collection.Find(context.TODO(), filter, opts)

	// 使用分页
	page := 1
	limit := 10
	skip := (page - 1) * limit
	opts = options.Find().SetLimit(int64(limit)).SetSkip(int64(skip))
	cursor, err = collection.Find(context.TODO(), filter, opts)
}

// 使用结构体来定义文档的格式
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Age   int                `bson:"age"`
	Email int                `bson:"email"`
}
