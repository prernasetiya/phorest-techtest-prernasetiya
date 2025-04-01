package controllers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "phorest-techtest/internal/database"
    "phorest-techtest/internal/models"
    "phorest-techtest/internal/routes"
)

func setupTestDb() {
    // Initialize the test database
    database.ConnectDatabase()

    // Clear existing data
    database.DB.Exec("DELETE FROM clients")

    // Seed test data
    clients := []models.Client{
        {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Phone: "1234567890", Gender: "male", Banned: false},
        {ID: "2", FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Phone: "0987654321", Gender: "female", Banned: false},
    }

    database.DB.Create(&clients)
}

func TestExtraController(t *testing.T) {
    // Set up Gin router
    router := routes.SetupRouter()

    // Set up the test database
    setupTestDb()

    t.Run("UpdateClient - Valid Request", func(t *testing.T) {
        // Create a test HTTP request
        clientUpdate := map[string]interface{}{
            "first_name": "Updated Name",
            "last_name":  "Updated Last Name",
            "email":      "updated@example.com",
            "phone":      "1112223333",
            "gender":     "male",
            "banned":     false,
        }
        body, _ := json.Marshal(clientUpdate)
        req, _ := http.NewRequest("PUT", "/clients/1", bytes.NewBuffer(body))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusOK, w.Code)

        // Assert the response body
        var updatedClient models.Client
        json.Unmarshal(w.Body.Bytes(), &updatedClient)
        assert.Equal(t, "Updated Name", updatedClient.FirstName)
        assert.Equal(t, "Updated Last Name", updatedClient.LastName)
        assert.Equal(t, "updated@example.com", updatedClient.Email)
        assert.Equal(t, "1112223333", updatedClient.Phone)
    })

    t.Run("GetClientByID - Valid Request", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("GET", "/clients/2", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusOK, w.Code)

        // Assert the response body
        var client models.Client
        json.Unmarshal(w.Body.Bytes(), &client)
        assert.Equal(t, "Jane", client.FirstName)
        assert.Equal(t, "Smith", client.LastName)
        assert.Equal(t, "jane@example.com", client.Email)
    })

    t.Run("GetClientByID - Client Not Found", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("GET", "/clients/999", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusNotFound, w.Code)

        // Assert the response body
        expected := `{"error":"Client not found"}`
        assert.JSONEq(t, expected, w.Body.String())
    })

    t.Run("DeleteClient - Valid Request", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("DELETE", "/clients/1", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusOK, w.Code)

        // Assert the response body
        expected := `{"message":"Client deleted successfully"}`
        assert.JSONEq(t, expected, w.Body.String())

        // Verify the client is deleted
        var client models.Client
        err := database.DB.First(&client, "id = ?", "1").Error
        assert.Error(t, err) // Should return an error because the client is deleted
    })

    t.Run("DeleteClient - Client Not Found", func(t *testing.T) {
        // Create a test HTTP request
        req, _ := http.NewRequest("DELETE", "/clients/999", nil)
        w := httptest.NewRecorder()

        // Perform the request
        router.ServeHTTP(w, req)

        // Assert the response code
        assert.Equal(t, http.StatusNotFound, w.Code)

        // Assert the response body
        expected := `{"error":"Client not found"}`
        assert.JSONEq(t, expected, w.Body.String())
    })
}