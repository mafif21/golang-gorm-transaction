package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-transaction-experiment/config"
	"golang-transaction-experiment/controller"
	"golang-transaction-experiment/http"
	"golang-transaction-experiment/repositories"
	"golang-transaction-experiment/service"
)

func main() {
	app := fiber.New()
	viperConf := config.NewViper()

	dbConf := &DBConfig{}
	dbConf.Read(viperConf)

	db, err := OpenConnection(dbConf)
	if err != nil {
		panic(err)
	}

	//txProvider := config.NewGenericTxProvider(db, func(tx *gorm.DB) *repositories.Adapters {
	//	return &repositories.Adapters{
	//		OrderRepository:   repositories.NewOrderRepositoryImpl(tx),
	//		UserRepository:    repositories.NewUserRepositoryImpl(tx),
	//		ProductRepository: repositories.NewProductRepositoryImpl(tx),
	//	}
	//})

	orderRepo := repositories.NewOrderRepositoryImpl(db)
	//orderService := service.NewOrderServiceImpl(txProvider)
	orderService := service.NewOrderServiceImpl2(orderRepo)
	orderController := controller.NewOrderControllerImpl(orderService)

	router := &http.RouteConfig{
		App:             app,
		OrderController: orderController,
	}

	router.Setup()

	router.App.Listen(":3000")
}
