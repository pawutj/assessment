package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/controllers"
	"github.com/pawutj/assessment/pkg/db"
	"github.com/pawutj/assessment/pkg/middleware"
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
	e.Use(middleware.AuthMiddleware)
	e.GET("/expenses", expenseController.GetExpensesController)
	e.GET("/expenses/:id", expenseController.GetExpenseController)
	e.POST("/expenses", expenseController.CreateExpenseController)
	e.PUT("/expenses/:id", expenseController.UpdateExpenseController)

	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Printf("after shutdown")
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
