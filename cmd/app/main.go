package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/zxcv9203/golang-be-sample/internal/config"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest"
	"log"
)

func main() {
	port := flag.String("p", "8080", "애플리케이션을 지정할 포트를 지정합니다.")
	flag.Parse()

	startFiber(*port)
}

func startFiber(port string) {
	app := fiber.New()

	handler := config.InitDependencies()
	rest.SetRoutes(app, handler)

	err := app.Listen(":" + port)
	if err != nil {
		log.Panicln(err)
	}
}
