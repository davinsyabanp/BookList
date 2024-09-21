package controllers

import (
	"books/database"
	"books/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pubDate, err := time.Parse("2006-01-02", book.PublicationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	result, err := database.DB.Exec("INSERT INTO books (title, author, publication_date, publisher, number_of_pages, category_id) VALUES (?, ?, ?, ?, ?, ?)",
		book.Title, book.Author, pubDate, book.Publisher, book.NumberOfPages, book.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)
	c.JSON(http.StatusCreated, book)
}

func ListBooks(c *gin.Context) {
	query := "SELECT * FROM books"
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		var pubDate []uint8
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &pubDate, &book.Publisher, &book.NumberOfPages, &book.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		book.PublicationDate = string(pubDate)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	var pubDate []uint8
	err := database.DB.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &pubDate, &book.Publisher, &book.NumberOfPages, &book.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	book.PublicationDate = string(pubDate)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	pubDate, err := time.Parse("2006-01-02", book.PublicationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	_, err = database.DB.Exec("UPDATE books SET title=?, author=?, publication_date=?, publisher=?, number_of_pages=?, category_id=? WHERE id=?",
		book.Title, book.Author, pubDate, book.Publisher, book.NumberOfPages, book.CategoryID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
