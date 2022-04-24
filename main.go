package main

import (
	handlers "book_store/pkg/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//middleware

	//database

	r := gin.New()           //routing
	book := r.Group("/book") //pathPrefix
	{
		book.GET("/", handlers.GetAllBook)    //fetches all the book
		book.GET("/:id", handlers.Searchbook) //fetches single book using path paramter
		book.GET("/author", handlers.BookbyAuthor)

		book.POST("/add", handlers.AddBook)

		book.PATCH("/:id", handlers.UpdateBook) //upadates a given book

		//book.GET("/purchase/:id", handlers.PurchaseBook)
		//extras

		book.GET("/purchase/inventory", handlers.Inventory)
	}
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	
}
