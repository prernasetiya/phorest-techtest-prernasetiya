package routes

import (
    "github.com/gin-gonic/gin"
    "phorest-techtest/internal/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // CSV import route : Imports , parses and saves CSV data to the database
    r.POST("/import", controllers.ImportCSV)

    return r
}