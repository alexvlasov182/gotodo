package main

import (
	"github.com/alexvlasov182/gotodo"
	"log"
	"github.com/alexvlasov182/gotodo/pkg/handler"
	"github.com/alexvlasov182/gotodo/pkg/repository"
	"github.com/alexvlasov182/gotodo/pkg/service"
	"github.com/spf13/viper"
)

func main()  {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error  {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
