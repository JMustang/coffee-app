package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JMustang/coffee-app/db"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	config Config
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("✅ API is listening on http://localhost:" + port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// TODO: add Handler
	}
	return srv.ListenAndServe()
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	app := &Application{
		config: cfg,
		// TODO: add models
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
