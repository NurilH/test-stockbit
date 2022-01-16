package main

import (
	middlewares "test_stockbit/middleware"
	"test_stockbit/routes"
)

func main() {
	e := routes.New()
	// take notes http method activity
	middlewares.LogMiddlewares(e)
	// start on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
