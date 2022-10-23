package main

import (
	"github.com/alexvlasov182/gotodo"
	"log"
	"github.com/alexvlasov182/gotodo/pkg/handler"
	"github.com/alexvlasov182/gotodo/pkg/repository"
	"github.com/alexvlasov182/gotodo/pkg/service"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
