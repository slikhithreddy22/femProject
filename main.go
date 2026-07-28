package main

import (
	"github.com/slikhithreddy22/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("we are running app")
}
