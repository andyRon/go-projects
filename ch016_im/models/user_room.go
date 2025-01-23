package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserRoom struct {
	RoomIdentity string `bson:"room_identity"`
	UserIdentity string `bson:"user_identity"`
	RoomType     int    `bson:"room_type"` // 房间类型，1 私聊，2 群聊
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByRoomIdentity(roomIdentity string) ([]*UserRoom, error) {
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		return nil, err
	}
	urs := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}

func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := Mongo.Collection(UserRoom{}.CollectionName()).FindOne(context.Background(),
		bson.D{{"user_identity", userIdentity}, {"room_identity", roomIdentity}}).Decode(ur)
	return ur, err
}

func JudgeUserIsFriend(ui1, ui2 string) bool {
	// ui1 单聊房间列表
	cur, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(),
		bson.D{{"user_identity", ui1}, {"room_type", 1}})
	roomIdentities := make([]string, 0)
	if err != nil {
		log.Printf("JudgeUserIsFriend error: %v\n", err)
		return false
	}
	for cur.Next(context.Background()) {
		ur := new(UserRoom)
		err := cur.Decode(ur)
		if err != nil {
			log.Printf("JudgeUserIsFriend error: %v\n", err)
			return false
		}
		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}

	// 获取关联ui2单聊房间的嗯数量
	cnt, err := Mongo.Collection(UserRoom{}.CollectionName()).CountDocuments(context.Background(),
		bson.M{"user_identity": ui2, "room_type": 1, "room_identity": bson.M{"$in": roomIdentities}})
	if err != nil {
		log.Printf("JudgeUserIsFriend error: %v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}
	return false
}

func InsertOneUserRoom(ur *UserRoom) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).InsertOne(context.Background(), ur)
	return err
}

func GetUserRoomIdentity(ui1, ui2 string) string {
	cur, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(),
		bson.D{{"user_identity", ui1}, {"room_type", 1}})
	roomIdentities := make([]string, 0)
	if err != nil {
		log.Printf("GetUserRoomByIdentity error: %v\n", err)
		return ""
	}
	if cur.Next(context.Background()) {
		ur := new(UserRoom)
		err := cur.Decode(ur)
		if err != nil {
			log.Printf("GetUserRoomByIdentity error: %v\n", err)
			return ""
		}
		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}

	ur := new(UserRoom)
	err = Mongo.Collection(UserRoom{}.CollectionName()).FindOne(context.Background(),
		bson.M{"user_identity": ui2, "room_room_typetype": 1, "room_identity": bson.M{"$in": roomIdentities}}).Decode(ur)
	if err != nil {
		log.Printf("GetUserRoomByIdentity error: %v\n", err)
		return ""
	}
	return ur.RoomIdentity
}

func DeleteUserRoom(roomIdentity string) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).DeleteOne(context.Background(),
		bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		log.Printf("DeleteUserRoom error: %v\n", err)
		return err
	}
	return nil
}
