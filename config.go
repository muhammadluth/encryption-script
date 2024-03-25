package main

import (
	"encryption-script/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func getServiceProperties() model.ServiceProperties {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	return getEnv()
}

func getEnv() model.ServiceProperties {
	fmt.Println("Starting Load Config " + time.Now().Format(time.RFC3339Nano))

	// SERVICE
	svcPort, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	svcPoolSizeConnection, _ := strconv.Atoi(os.Getenv("SERVICE_POOL_SIZE_CONNECTION"))
	svcTimezone, _ := time.LoadLocation(os.Getenv("SERVICE_TIMEZONE"))
	svcDebugMode, _ := strconv.ParseBool(os.Getenv("SERVICE_DEBUG_MODE"))

	svcProperties := model.ServiceProperties{
		ServiceName:               os.Getenv("SERVICE_NAME"),
		ServicePort:               svcPort,
		ServicePoolSizeConnection: svcPoolSizeConnection,
		ServiceTimezone:           svcTimezone,
		ServiceDebugMode:          svcDebugMode,
	}
	if err := validator.New().Struct(svcProperties); err != nil {
		panic(err)
	}
	fmt.Println("Finish Load Config " + time.Now().Format(time.RFC3339Nano))
	return svcProperties
}
