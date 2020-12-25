package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"fmt"
)

func Test (c echo.Context) error {
	
	fmt.Println("1")
	var people []Person
	fmt.Println("2")
	DB.Find(&people)
	fmt.Println("3")
	fmt.Println(people)
	return c.JSON(http.StatusOK, people)

}