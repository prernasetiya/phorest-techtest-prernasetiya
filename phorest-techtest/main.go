package main

import (
    "phorest-techtest/internal/database"
    "phorest-techtest/internal/routes"
)

func main() {
    // Connect to the database
    database.ConnectDatabase()

    // Set up routes for Rest API
    r := routes.SetupRouter()

    // Start the server
    r.Run(":8080") // Run on port 8080
}