package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gomasters/repository"
	"gomasters/router"
	"gomasters/service"
	"log"
)

func main() {
	if err := readConfig(); err != nil{
		log.Fatalf("error occured during reading config: %v", err)
	}
	var server router.Server
	config := repository.Config{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	}
	db, err := repository.InitDB(config)
	if err != nil{
		log.Fatalf("error occured during connecting to DB: %v", err)
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := router.NewHandler(service)
	if err := server.Run("8080", handler.InitRoutes()); err != nil{
		log.Fatalf("an error occured during server's work: %s",err)
	}
}

func readConfig() error{
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs/")   // path to look for the config file in
	return viper.ReadInConfig()

}