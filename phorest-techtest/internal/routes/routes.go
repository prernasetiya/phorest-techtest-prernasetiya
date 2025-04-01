package routes

import (
    "github.com/gin-gonic/gin"
    "phorest-techtest/internal/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // CSV import route : Imports , parses and saves CSV data to the database
    r.POST("/import", controllers.ImportCSV)

	// Get top X non-banned clients with most loyalty points since date Y
	r.GET("/top-clients", controllers.GetTopClients)

    return r
}