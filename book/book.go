package book

import (
	"fibertest/database"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"titile"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	fmt.Println("in get books function")
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}
func GetBook(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "a"
	book.Author = "b"
	book.Rating = 5
	db.Create(book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("a books")
}
