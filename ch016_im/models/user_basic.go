package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	user := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).FindOne(context.Background(), bson.M{"account": account, "password": password}).Decode(user)
	return user, err
}

func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	user := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).FindOne(context.Background(), bson.M{"identity": identity}).Decode(user)
	return user, err
}

func GetUserBasicByAccount(account string) (*UserBasic, error) {
	user := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).FindOne(context.Background(), bson.M{"account": account}).Decode(user)
	return user, err
}

func GetUserBasicCountByEmail(email string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).CountDocuments(context.Background(), bson.M{"email": email})
}

func GetUserBasicCountByAccount(account string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).CountDocuments(context.Background(), bson.M{"account": account})
}

func InsertOneUserBasic(user *UserBasic) error {
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).InsertOne(context.Background(), user)
	return err
}
