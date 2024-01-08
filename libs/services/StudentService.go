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

var Create = func(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Student{12, "oeir"})
}

var Read = func(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Student{12, "oeir"})
}