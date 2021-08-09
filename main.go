package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Books struct {
	Id     int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Title  string `gorm:"not null" form:"title" json:"title"`
	Author string `gorm:"not null" form:"author" json:"author"`
}

func InitDb() *gorm.DB {
	// Opening sqlite db file
	db, err := gorm.Open("sqlite3", "./data.db")

	// Display SQL queries
	db.LogMode(true)

	// When error occures
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Books{}) {
		db.CreateTable(&Books{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Books{})
	}

	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/books", PostBook)
		v1.GET("/books", GetBooks)
		v1.GET("/books/:id", GetBook)
		v1.PUT("/books/:id", UpdateBook)
		v1.DELETE("/books/:id", DeleteBook)
	}

	r.Run(":8080")
}

func PostBook(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var book Books
	c.Bind(&book)

	if book.Title != "" && book.Author != "" {
		// INSERT INTO "books" (name) VALUES (book.Name);
		db.Create(&book)
		// Display error
		c.JSON(201, gin.H{"success": book})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"title\": \"How to be rich\", \"author\": \"FatPotato\" }" http://localhost:8080/api/v1/books
}

func GetBooks(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var books []Books
	// SELECT * FROM books
	db.Find(&books)

	// Display JSON result
	c.JSON(200, books)

	// curl -i http://localhost:8080/api/v1/books
}

func GetBook(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var book Books
	// SELECT * FROM books WHERE id = 1;
	db.First(&book, id)

	if book.Id != 0 {
		// Display JSON result
		c.JSON(200, book)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Book not found"})
	}

	// curl -i http://localhost:8080/api/v1/books/1
}

func UpdateBook(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id book
	id := c.Params.ByName("id")
	var book Books
	// SELECT * FROM books WHERE id = 1;
	db.First(&book, id)

	if book.Title != "" && book.Author != "" {

		if book.Id != 0 {
			var newBook Books
			c.Bind(&newBook)

			result := Books{
				Id:     book.Id,
				Title:  newBook.Title,
				Author: newBook.Author,
			}

			// UPDATE books SET title='newBook.Title', author='newBook.Author' WHERE id = book.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Book not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"title\": \"Thea\", \"author\": \"Merlyn\" }" http://localhost:8080/api/v1/books/1
}

func DeleteBook(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id book
	id := c.Params.ByName("id")
	var book Books
	// SELECT * FROM books WHERE id = 1;
	db.First(&book, id)

	if book.Id != 0 {
		// DELETE FROM books WHERE id = book.Id
		db.Delete(&book)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Book #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Book not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/books/1
}

func OptionsBook(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
