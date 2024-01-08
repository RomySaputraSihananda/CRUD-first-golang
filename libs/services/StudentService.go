package libs_services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var students []Student

func init (){
	fmt.Println("ok")
}

var Create = func(ctx *gin.Context) {
	var requestBody Student

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	requestBody.Id = len(students)
	students = append(students, requestBody)

	ctx.JSON(http.StatusOK, requestBody)
}

var Read = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, students)
}