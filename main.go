package main

import (
	"fmt"
	"github.com/Fabricio2210/gofiber/elastic"
	"github.com/Fabricio2210/gofiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	elastic.ConnectElastic()
	fmt.Println("Running!!!!!")
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${latency}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))
	router.DefaultRouter(app, "DSP")
	router.DefaultRouter(app, "DDM")
	router.DefaultRouter(app, "RAW")
	router.DefaultRouter(app, "POP")
	router.DefaultRouter(app, "SHINKO")
	router.DefaultRouter(app, "reacts")
	router.DefaultRouter(app, "AQUA")
	router.DefaultRouter(app, "BEAM")
	router.DefaultRouter(app, "DECEPTICRON")
	router.DefaultRouter(app, "DOODY")
	router.DefaultRouter(app, "MEERKAT")
	router.DefaultRouter(app, "PROPER")
	router.DefaultRouter(app, "TBS")
	router.DefaultRouter(app, "WPIG")
	app.Listen(":3000")
}
