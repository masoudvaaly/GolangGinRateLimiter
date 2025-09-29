package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ratelimiter/config"
	"ratelimiter/models"
)

// Create a new item
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO items (name, description, price) VALUES ($1, $2, $3) RETURNING id"
	err := config.DB.QueryRow(query, item.Name, item.Description, item.Price).Scan(&item.ID)
	if err != nil {
		fmt.Printf("DB insert error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func GetItems(c *gin.Context) {
	query := "SELECT id, name, description, price FROM items"
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get items"})
		return
	}
	defer rows.Close()

	var items []models.Item // assuming you have a struct Item

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan item"})
			return
		}
		items = append(items, item)
	}

	// check for iteration errors
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over items"})
		return
	}

	c.JSON(http.StatusOK, items)
}
