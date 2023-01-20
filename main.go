package main

import (
	"e-commerce-api/config"
	"e-commerce-api/controller"
	"e-commerce-api/exception"
	"e-commerce-api/repository"
	"e-commerce-api/service"
	"e-commerce-api/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.ConnectDatabase()
	merchRespository := repository.NewMerchRepository()
	merchService := service.NewMerchService(merchRespository, db)
	merchController := controller.NewMerchController(merchService)

	router := httprouter.New()

	router.GET("/api/v1/merch", merchController.FindAll)
	router.GET("/api/v1/merch/:merchId", merchController.FindById)
	router.POST("/api/v1/merch", merchController.Create)
	router.PUT("/api/v1/merch/:merchId", merchController.Update)
	router.DELETE("/api/v1/merch/:merchId", merchController.Delete)
	//router.POST("/api/v1/merch/:merchId", merchController.Order)

	router.PanicHandler = exception.ErrorHandler

	err := http.ListenAndServe(":8080", router)
	utils.PanicIfError(err)
	log.Print("The is Server Running on localhost port 8080")
}