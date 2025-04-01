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


    // Update a client by ID
    r.PUT("/clients/:id", controllers.UpdateClient)  

    // Fetch a singel client by ID 
    r.GET("/clients/:id", controllers.GetClientByID)

    // Delete a client by ID
    r.DELETE("/clients/:id", controllers.DeleteClient) 

    return r
}