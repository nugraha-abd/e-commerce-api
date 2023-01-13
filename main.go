package main

import (
	"e-commerce-api/config"
	"e-commerce-api/controller"
	"e-commerce-api/repository"
	"e-commerce-api/service"
	"e-commerce-api/utils"
	"log"
	"net/http"
)

func main() {
	db := config.ConnectDatabase()
	merchRespository := repository.NewMerchRepository()
	merchService := service.NewMerchService(merchRespository, db)
	merchController := controller.NewMerchController(merchService)

	router := http.NewServeMux()
	router.HandleFunc("/api/v1/merch", merchController.FindAll)
	router.HandleFunc("/api/v1/merch/:merchId", merchController.FindById)
	router.HandleFunc("/api/v1/merch", merchController.Create)
	router.HandleFunc("/api/v1/merch/:merchId", merchController.Update)
	router.HandleFunc("/api/v1/merch/:merchId", merchController.Delete)
	//r.HandleFunc("/api/v1/merch/:merchId", merchController.Order)

	err := http.ListenAndServe(":8080", router)
	utils.PanicIfError(err)
	log.Print("The is Server Running on localhost port 8080")
}