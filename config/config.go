package config

import (
	"context"
	"log"
	"project/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_user *mongo.Collection
var DB_school *mongo.Collection

func initDB() (DB *mongo.Database, ctx context.Context) {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://test:test68@cluster0.f1xd1.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("Cluster0")
	return DB, ctx
}

func GetUserDB(DB *mongo.Database, ctx context.Context) {
	var users []model.User

	DB_user = DB.Collection("userdb")
	cursor_user, err := DB_user.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor_user.All(ctx, &users); err != nil {
		panic(err)
	}
	defer cursor_user.Close(ctx)
}

func GetSchoolDB(DB *mongo.Database, ctx context.Context) {
	var schools []model.School
	DB_school = DB.Collection("schooldb")
	cursor_school, err := DB_school.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor_school.All(ctx, &schools); err != nil {
		panic(err)
	}
	defer cursor_school.Close(ctx)
}

func GetClient() {
	DB, _ := initDB()
	DB_user = DB.Collection("userdb")
	DB_school = DB.Collection("schooldb")
	//GetUserDB(DB, ctx)
	//GetSchoolDB(DB, ctx)
}
