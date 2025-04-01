package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "phorest-techtest/internal/database"
    "phorest-techtest/internal/models"
)



func UpdateClient(c *gin.Context) {
    id := c.Param("id") // Get the client ID from the URL

    var client models.Client
    // Check if the client exists
    if err := database.DB.First(&client, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    // Create a struct for the update payload
    var updatePayload struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
        Email     string `json:"email"`
        Phone     string `json:"phone"`
        Gender    string `json:"gender"`
        Banned    bool   `json:"banned"`
    }

    // Bind the JSON body to the update payload
    if err := c.ShouldBindJSON(&updatePayload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Update the client fields
    client.FirstName = updatePayload.FirstName
    client.LastName = updatePayload.LastName
    client.Email = updatePayload.Email
    client.Phone = updatePayload.Phone
    client.Gender = updatePayload.Gender
    client.Banned = updatePayload.Banned

    // Save the updated client
    if err := database.DB.Save(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func GetClientByID(c *gin.Context) {
    id := c.Param("id") // Get the client ID from the URL

    var client models.Client
    // Fetch the client from the database
    if err := database.DB.First(&client, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    c.JSON(http.StatusOK, client)
}




func DeleteClient(c *gin.Context) {
    id := c.Param("id") // Get the client ID from the URL

    // Check if the client exists
    var client models.Client
    if err := database.DB.First(&client, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    // Attempt to delete the client
    if err := database.DB.Delete(&client).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

