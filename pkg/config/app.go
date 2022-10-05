package config

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var client *mongo.Client
var ctx = context.TODO()
func ViperEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
  
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
  
	value, ok := viper.Get(key).(string)
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
  
	return value
  }
func Connect() *mongo.Client{
	 db := ViperEnvVariable("MONGO_URI")
	
	clientOptions := options.Client().ApplyURI(db)
	connectDB, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
	fmt.Println(err)
	}

	client=connectDB
	fmt.Println("Database connected")
	return client
}
func GetDB() *mongo.Client{
	if client==nil{
		return Connect()
	}
	return client
}
func CloseClientDB() {
    if client == nil {
        return
    }

    err := client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connection to MongoDB closed.")
}