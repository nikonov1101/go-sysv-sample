package main

import (
	"github.com/labstack/echo"
	"log"
	"time"
)

func main() {
	go shitLogger()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		log.Println("Handle HTTP Request...")
		return c.JSON(200, map[string]interface{}{
			"status": "OK",
		})
	})
	e.Start("127.0.0.1:1489")
}

func shitLogger() {
	defer shitLogger()
	log.Println("Some log output.")
	time.Sleep(time.Second * 10)
}
