package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/controllers"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	e := echo.New()
	e.GET("/expenses/:id", controllers.GetExpenseController)

	log.Fatal(e.Start(":2565"))
}
