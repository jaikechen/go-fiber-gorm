package main

import (
	"fibertest/book"
	"fibertest/database"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("aaa")
	}
	fmt.Println("database opened")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("migrated")

}

func helloWorld(c *fiber.Ctx) {

	c.Send("hello, ")
}

func setupRoutes(app *fiber.App) {

	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Delete("/api/book", book.DeleteBook)
}

func main() {
	fmt.Println("auto resetart")
	fmt.Println("auto resetart----fffff")
	initDatabase()
	defer database.DBConn.Close()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	fmt.Println("listen on 3000")

}
