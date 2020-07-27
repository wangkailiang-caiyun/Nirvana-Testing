package user

import (
	"context"

	"github.com/caicloud/nirvana/log"
	"github.com/wangkailiang-caiyun/Nirvana-Testing/pkg/mgo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User 用户基础信息
type User struct {
	ID         string   `json:"id" bson:"_id"`
	UserName   string   `json:"user_name" bson:"user_name"`
	Password   string   `json:"password" bson:"password"`
	CreateAt   int64    `json:"create_at" bson:"create_at"`
	LastModify int64    `json:"last_modify" bson:"last_modify"`
	Address    []string `json:"address" bson:"address"`
}

//AddressRequest push address request
type AddressRequest struct {
	UserID  string `json:"user_id" `
	Address string `json:"address"`
}

// GetUserList 获取用户列表
func GetUserList(ctx context.Context, pageSize, pageNo int) ([]User, error) {
	var offset int64 = int64(pageNo-1) * int64(pageSize)

	limit := int64(pageSize)

	userTable := mgo.Mongo.Database("testing").Collection("user")
	c := context.Background()
	r, err := userTable.Find(c, bson.D{}, &options.FindOptions{Limit: &limit, Skip: &offset, Sort: bson.M{"create_at": 1}})

	if err != nil {
		return nil, err
	}
	defer r.Close(c)

	result := make([]User, 0, pageSize)
	for r.Next(c) {
		var u User
		r.Decode(&u)
		result = append(result, u)
	}
	return result, nil
}

// FetchUser 获取用户详细信息
func FetchUser(ctx context.Context, userID string) (User, error) {
	userTable := mgo.Mongo.Database("testing").Collection("user")
	objID, _ := primitive.ObjectIDFromHex(userID)
	result := userTable.FindOne(context.Background(), bson.M{"_id": objID})
	user := User{}
	err := result.Decode(&user)
	return user, err
}

//CreateUser 创建用户
func CreateUser(ctx context.Context, user User) (*User, error) {
	userTable := mgo.Mongo.Database("testing").Collection("user")
	b := bson.M{"user_name": user.UserName, "password": user.Password, "create_at": user.CreateAt, "last_modify": user.LastModify}
	r, err := userTable.InsertOne(context.Background(), b)
	if err != nil {
		return nil, err
	}
	if v, ok := r.InsertedID.(primitive.ObjectID); ok {
		user.ID = v.Hex()
	}
	return &user, nil
}

//UpdateUser 更新用户信息
func UpdateUser(ctx context.Context, userID string, user User) (bool, error) {
	userTable := mgo.Mongo.Database("testing").Collection("user")
	ID, _ := primitive.ObjectIDFromHex(userID)
	b := bson.M{"$set": bson.M{"user_name": user.UserName, "password": user.Password, "create_at": user.CreateAt, "last_modify": user.LastModify}}
	result, err := userTable.UpdateOne(context.Background(), bson.M{"_id": ID}, b)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}

//DeleteUser 删除用户
func DeleteUser(ctx context.Context, userID string) (bool, error) {
	userTable := mgo.Mongo.Database("testing").Collection("user")
	objID, _ := primitive.ObjectIDFromHex(userID)
	result, err := userTable.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}

//AddUserAddress 添加用户地址
func AddUserAddress(ctx context.Context, addRequest AddressRequest) (bool, error) {
	userTable := mgo.Mongo.Database("testing").Collection("user")
	objID, _ := primitive.ObjectIDFromHex(addRequest.UserID)

	b := bson.M{
		"$push": bson.M{"address": addRequest.Address},
	}
	result, err := userTable.UpdateOne(context.Background(), bson.M{"_id": objID}, b)
	if err != nil {
		return false, err
	}
	log.Infof("userID : %s Address : %s \n", addRequest.UserID, addRequest.Address)
	return result.ModifiedCount == 1, nil
}

//RemoveUserAddress 删除用户地址
func RemoveUserAddress(ctx context.Context, addRequest AddressRequest) (bool, error) {
	userTable := mgo.Mongo.Database("testing").Collection("user")
	objID, _ := primitive.ObjectIDFromHex(addRequest.UserID)

	b := bson.M{
		"$pull": bson.M{"address": addRequest.Address},
	}
	result, err := userTable.UpdateOne(context.Background(), bson.M{"_id": objID}, b)
	if err != nil {
		return false, err
	}
	log.Infof("userID : %s Address : %s \n", addRequest.UserID, addRequest.Address)
	return result.ModifiedCount == 1, nil
}
