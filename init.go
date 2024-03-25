package main

import (
	"encryption-script/app/server"
	"encryption-script/src/handler"
	"encryption-script/src/handler/router"
	"encryption-script/src/usecase"
	"fmt"
	"strings"
)

func RunApplication() {
	fmt.Println("Init Configuration")
	svcProperties := getServiceProperties()
	fmt.Printf("%s SERVICE\n", strings.ToUpper(svcProperties.ServiceName))

	iSetupServer := server.NewSetupServer(svcProperties)
	_, fiberRouter := iSetupServer.InitServerConfiguration()

	iRSAEncryptionUsecase := usecase.NewRSAEncryptionUsecase()
	iRSAEncryptionRouter := router.NewRSAEncryptionRouter(iRSAEncryptionUsecase)
	iEncryptionHttpHandler := handler.NewEncryptionHttpHandler(fiberRouter, iRSAEncryptionRouter)
	iEncryptionHttpHandler.Routers()

	// setup server
	iSetupServer.InitServer()
}
