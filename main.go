package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	docs "app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var students []Student

const (
	PORT int = 4444
	TAG string = "/api/v1"
)

// Ping godoc
// @Summary Ping example
// @Description Pings the server
// @Tags example
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
    c.JSON(http.StatusOK, "pong")
}


func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger-ui/index.html")
	})

	router.GET("/ping", Ping)

	router.Run(fmt.Sprintf(":%d", PORT))
}

func handleStudents(w http.ResponseWriter, r *http.Request){
	switch(r.Method){
		case "GET":
			getData(w, r)
		case "POST":
			addData(w, r)
		case "PUT":
			fmt.Print(r.Method)
		case "DELETE":
			fmt.Print(r.Method)
		default:
			fmt.Print(r.Method)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func addData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newStudent Student

	json.NewDecoder(r.Body).Decode(&newStudent)
	newStudent.Id = len(students) + 1
	students = append(students, newStudent)

	json.NewEncoder(w).Encode(newStudent)
}

