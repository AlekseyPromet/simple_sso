package models

type Application struct {
	ApplicationId string `bson:"applicationId"`
	Secret        string `bson:"secret"`
}
