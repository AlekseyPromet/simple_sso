package models

type User struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type LoginPass struct {
	AppID    string `bson:"appId"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
