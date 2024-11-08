package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON data"})
		}

		db.Create(&todo)

		c.JSON(200, todo)
	})

	// Route to get all Todos
	router.GET("/todos", func(c *gin.Context) {
		var todos []Todo

		db.Find(&todos)

		c.JSON(200, todos)
	})

	// Route to get a specific Todo_task by id
	router.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		todoID := c.Param("id")

		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		c.JSON(200, todo)
	})

	router.Run(":8080")
}
