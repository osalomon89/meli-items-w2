package main

import (
	"log"

	"github.com/osalomon89/meli-items-w2/internal/adapters/controller"
	"github.com/osalomon89/meli-items-w2/internal/adapters/repository"
	"github.com/osalomon89/meli-items-w2/internal/infrastructure/mysql"
	"github.com/osalomon89/meli-items-w2/internal/infrastructure/web"
	"github.com/osalomon89/meli-items-w2/internal/usecase"
)

func main() {

	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatal(err)
	}

	// itemRepository := repository.NewItemRepository()
	itemRepository := repository.NewMysqlItemRepository(conn)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemController := controller.NewItemController(itemUsecase)

	err = web.NewHTTPServer(itemController)
	if err != nil {
		log.Fatal(err)
	}

}
