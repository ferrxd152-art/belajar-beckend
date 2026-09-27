package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database")
	}

	db.AutoMigrate(&Task{})

	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []Task
		db.Find(&tasks)

		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var task Task

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "JSON tidak valid",
			})
			return
		}

		db.Create(&task)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Task berhasil dibuat",
			"task":    task,
		})
	})

	r.Run(":8080")
}