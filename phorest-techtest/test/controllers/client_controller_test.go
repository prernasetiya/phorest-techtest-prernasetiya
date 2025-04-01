package controllers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "phorest-techtest/internal/database"
    "phorest-techtest/internal/models"
    "phorest-techtest/internal/routes"
)

func setupTestDatabase() {
    // Initialize the test database
    database.ConnectDatabase()

    // Clear existing data
    database.DB.Exec("DELETE FROM purchases")
    database.DB.Exec("DELETE FROM services")
    database.DB.Exec("DELETE FROM appointments")
    database.DB.Exec("DELETE FROM clients")

    // Seed test data
    clients := []models.Client{
        {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Phone: "0123", Gender: "Male", Banned: false},
        {ID: "2", FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Phone: "12345", Gender: "Female", Banned: false},
        {ID: "3", FirstName: "Banned", LastName: "Client", Email: "banned@example.com", Phone: "34567", Gender: "Male", Banned: true},
    }

    appointments := []models.Appointment{
        {ID: "a1", ClientID: "1", StartTime: "2023-01-01T10:00:00Z", EndTime: "2023-01-01T11:00:00Z"},
        {ID: "a2", ClientID: "2", StartTime: "2023-01-01T12:00:00Z", EndTime: "2023-01-01T13:00:00Z"},
    }

    services := []models.Service{
        {ID: "s1", AppointmentID: "a1", Name: "Service 1", Price: 50, LoyaltyPoints: 50},
        {ID: "s2", AppointmentID: "a2", Name: "Service 2", Price: 40, LoyaltyPoints: 30},
    }

    purchases := []models.Purchase{
        {ID: "p1", AppointmentID: "a1", Name: "Purchase 1", Price: 20, LoyaltyPoints: 20},
        {ID: "p2", AppointmentID: "a2", Name: "Purchase 2", Price: 50, LoyaltyPoints: 10},
    }

    database.DB.Create(&clients)
    database.DB.Create(&appointments)
    database.DB.Create(&services)
    database.DB.Create(&purchases)
}

func TestGetTopClients(t *testing.T) {
    // Set up Gin router
    router := routes.SetupRouter()

    // Set up the test database
    setupTestDatabase()

    t.Run("Valid request with top clients", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("GET", "/top-clients?limit=2&since=2023-01-01", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusOK, w.Code)

        // Assert the response body
        expected := `[{"id":"1","first_name":"John","last_name":"Doe","email":"john@example.com","total_loyalty_points":70},{"id":"2","first_name":"Jane","last_name":"Smith","email":"jane@example.com","total_loyalty_points":40}]`
        assert.JSONEq(t, expected, w.Body.String())
    })

    t.Run("Request with invalid limit parameter", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("GET", "/top-clients?limit=-1&since=2023-01-01", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusBadRequest, w.Code)

        // Assert the response body
        expected := `{"error":"Invalid limit parameter. Limit must be a positive integer"}`
        assert.JSONEq(t, expected, w.Body.String())
    })

    t.Run("Request with invalid date parameter", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("GET", "/top-clients?limit=2&since=invalid-date", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusBadRequest, w.Code)

        // Assert the response body
        expected := `{"error":"Invalid date format. Use YYYY-MM-DD"}`
        assert.JSONEq(t, expected, w.Body.String())
    })

    t.Run("Request with no clients found", func(t *testing.T) {
        // Create a test HTTP request with a future date
        req, _ := http.NewRequest("GET", "/top-clients?limit=2&since=2030-01-01", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusOK, w.Code)

        // Assert the response body
        expected := `[]`
        assert.JSONEq(t, expected, w.Body.String())
    })
}