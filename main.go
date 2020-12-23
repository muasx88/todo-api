package main

import (
	"github.com/joho/godotenv"

	"github.com/muasx/todo_api/db"
	"github.com/muasx/todo_api/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("error load env")
	}

	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":8181"))
}
