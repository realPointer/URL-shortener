package main

import (
	"github.com/realPointer/url-shortener/internal/app"
)

// @title url-shortener service
// @version 1.0.0
// @description Just a link shortener

// @host localhost:8080
// @BasePath /v1

// @contact.name Andrew
// @contact.url https://t.me/realPointer

func main() {
	// Run application
	app.Run()
}
