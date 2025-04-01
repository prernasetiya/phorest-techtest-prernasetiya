package controllers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "phorest-techtest/internal/database"
)

func GetTopClients(c *gin.Context) {
    // Parse query parameters
    limitParam := c.Query("limit") // Top X
    dateParam := c.Query("since") // Since date Y

    // Validate limit parameter
    limit, err := strconv.Atoi(limitParam)
    if err != nil || limit <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter. Limit must be a positive integer"})
        return
    }

    // Validate date parameter
    sinceDate, err := time.Parse("2006-01-02", dateParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
        return
    }

    // Query the database for top clients
    var results []struct {
        ID               string `json:"id"`
        FirstName        string `json:"first_name"`
        LastName         string `json:"last_name"`
        Email            string `json:"email"`
        TotalLoyaltyPoints int   `json:"total_loyalty_points"`
    }

    err = database.DB.Raw(`
        SELECT c.id, c.first_name, c.last_name, c.email,
               SUM(s.loyalty_points + COALESCE(p.loyalty_points, 0)) AS total_loyalty_points
        FROM clients c
        JOIN appointments a ON c.id = a.client_id
        LEFT JOIN services s ON a.id = s.appointment_id
        LEFT JOIN purchases p ON a.id = p.appointment_id
        WHERE c.banned = false AND a.start_time >= ?
        GROUP BY c.id
        ORDER BY total_loyalty_points DESC
        LIMIT ?
    `, sinceDate, limit).Scan(&results).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top clients"})
        return
    }

    
    if len(results) == 0 {
        c.JSON(http.StatusOK, []struct{}{})
        return
    }

    // Return the result
    c.JSON(http.StatusOK, results)
}