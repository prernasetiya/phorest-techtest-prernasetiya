package controllers

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/gocarina/gocsv"
    "phorest-techtest/internal/database"
    "phorest-techtest/internal/models"
)

func ImportCSV(c *gin.Context) {
    // Parse the uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
        return
    }

    // Save the file temporarily
    tempFile := "./temp/uploaded.csv"
    if err := c.SaveUploadedFile(file, tempFile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }
    defer os.Remove(tempFile) // Clean up the temp file

    // Open the file
    f, err := os.Open(tempFile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
        return
    }
    defer f.Close()

    // Determine the type of data based on the file name
    switch file.Filename {
    case "clients.csv":
        var clients []models.Client
        if err := gocsv.UnmarshalFile(f, &clients); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse clients.csv"})
            return
        }
        for _, client := range clients {
            database.DB.Create(&client)
        }
    case "appointments.csv":
        var appointments []models.Appointment
        if err := gocsv.UnmarshalFile(f, &appointments); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse appointments.csv"})
            return
        }
        for _, appointment := range appointments {
            database.DB.Create(&appointment)
        }
    case "services.csv":
        var services []models.Service
        if err := gocsv.UnmarshalFile(f, &services); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse services.csv"})
            return
        }
        for _, service := range services {
            database.DB.Create(&service)
        }
    case "purchases.csv":
        var purchases []models.Purchase
        if err := gocsv.UnmarshalFile(f, &purchases); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse purchases.csv"})
            return
        }
        for _, purchase := range purchases {
            database.DB.Create(&purchase)
        }
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
}