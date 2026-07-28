package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/slikhithreddy22/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("we are running app")

	http.HandleFunc("/healthcheck", HealthCheck)
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
