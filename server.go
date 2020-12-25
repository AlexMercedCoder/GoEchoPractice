package main
import (
	// "net/http"
	"github.com/labstack/echo/v4"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// "fmt"
)

func main() {

	// //OPEN DATABASE CONNECTION
	// DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	// //Print Error if DB Doesn't Connect
	// if err != nil {
	// 	panic("failed to connect database")
	//   } else {
	// 	fmt.Println("Database is Connected")
	//   }

	//Connect to Database
	Connect()

	//MIGRATE PERSON
	// DB.AutoMigrate(&Person{})

	//Create new App
	app := echo.New()

	//Create
	// DB.Create(&Person{Name: "Tony", Age: 32})

	// test2 := func (c echo.Context) error {
	
	// 	var people []Person
	
	// 	result := DB.Find(&people)
		
	// 	return c.JSON(http.StatusOK, result)
	
	// }

	//Root Route
	app.GET("/", Test)

	//Start Server
	app.Logger.Fatal(app.Start(":4000"))

}

