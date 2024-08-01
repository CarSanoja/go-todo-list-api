package handlers

import (
	"encoding/csv"
	"net/http"
	"os"

	"go-todo-list-api/config"
	"go-todo-list-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var todos []models.ToDo

func LoadToDos() error {
	file, err := os.Open(config.GetConfig().CSVFile)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	todos = []models.ToDo{}
	for _, record := range records[1:] {
		todo := models.ToDo{
			ID:        record[0],
			Title:     record[1],
			Completed: record[2] == "true",
		}
		todos = append(todos, todo)
	}
	return nil
}

func SaveToDos() error {
	file, err := os.Create(config.GetConfig().CSVFile)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write([]string{"ID", "Title", "Completed"})
	for _, todo := range todos {
		w.Write([]string{todo.ID, todo.Title, boolToStr(todo.Completed)})
	}
	return nil
}

func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func GetToDos(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, todos)
}

func GetToDoByID(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
}

func CreateToDo(c *gin.Context) {
	var todo models.ToDo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = uuid.New().String()
	todos = append(todos, todo)
	if err := SaveToDos(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, todo)
}

func UpdateToDo(c *gin.Context) {
	id := c.Param("id")
	var updatedToDo models.ToDo
	if err := c.ShouldBindJSON(&updatedToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			updatedToDo.ID = id
			todos[i] = updatedToDo
			if err := SaveToDos(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, updatedToDo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
}

func DeleteToDo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			if err := SaveToDos(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "ToDo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
}
