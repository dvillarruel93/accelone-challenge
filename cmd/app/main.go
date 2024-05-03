package main

import (
	"accelone-challenge/delivery/rest"
	"accelone-challenge/repository"
	"accelone-challenge/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	// initialize memory DB
	memoryDB := make(map[string]interface{})

	// initialize layers
	contactRepo := repository.NewContactRepository(memoryDB)
	contactUsecase := usecase.NewContactUsecase(contactRepo)

	// initialize echo web framework
	e := echo.New()

	// Debug is a flag to show more details in console for each request
	e.Debug = true
	// HideBanner is a flag to hide a banner once you start the app
	e.HideBanner = true

	rest.NewContactHandler(e, contactUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
