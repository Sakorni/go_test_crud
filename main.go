package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gomasters/repository"
	"gomasters/router"
	"gomasters/service"
	"os"
)

func main() {
	if err := readConfig(); err != nil{
		logrus.Fatalf("error occured during reading config: %v", err)
	}
	if err := godotenv.Load(); err != nil{
		logrus.Fatalf("error occured during parsing .env: %v", err)
	}
	var server router.Server
	config := repository.Config{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	db, err := repository.InitDB(config)
	if err != nil{
		logrus.Fatalf("error occured during connecting to DB: %v", err)
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := router.NewHandler(service)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil{
		logrus.Fatalf("an error occured during server's work: %s",err)
	}
}

func readConfig() error{
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs/")   // path to look for the config file in
	return viper.ReadInConfig()

}