package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/controllers"
	"github.com/pawutj/assessment/pkg/db"
	"github.com/pawutj/assessment/pkg/repositories"
	"github.com/pawutj/assessment/pkg/services"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:" + os.Getenv("PORT"))

	expenseController := controllers.ExpenseController{
		ExpenseService: services.ExpenseService{
			ExpenseRepository: repositories.ExpenseRepository{DB: db.ConnectDB()},
		},
	}

	e := echo.New()
	e.GET("/expenses/:id", expenseController.GetExpenseController)
	e.POST("/expenses", expenseController.CreateExpenseController)

	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}
