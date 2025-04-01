package controllers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"
    "mime/multipart"
	"io"
    "github.com/stretchr/testify/assert"
    "phorest-techtest/internal/database"
    "phorest-techtest/internal/models"
    "phorest-techtest/internal/routes"
)

func setupTestDB() {
    // Initialize the test database
    database.ConnectDatabase()

    // Clear existing data
    database.DB.Exec("DELETE FROM purchases")
    database.DB.Exec("DELETE FROM services")
    database.DB.Exec("DELETE FROM appointments")
    database.DB.Exec("DELETE FROM clients")
}

func TestImportCSV(t *testing.T) {
    // Set up Gin router
    router := routes.SetupRouter()

    // Set up the test database
    setupTestDB()

    // Define test cases
    tests := []struct {
        fileName     string
        filePath     string
        expectedCode int
        model        interface{}
    }{
        {
            fileName:     "clients.csv",
            filePath:     "../../data/clients.csv",
            expectedCode: http.StatusOK,
            model:        &models.Client{},
        },
        {
            fileName:     "appointments.csv",
            filePath:     "../../data/appointments.csv",
            expectedCode: http.StatusOK,
            model:        &models.Appointment{},
        },
        {
            fileName:     "services.csv",
            filePath:     "../../data/services.csv",
            expectedCode: http.StatusOK,
            model:        &models.Service{},
        },
        {
            fileName:     "purchases.csv",
            filePath:     "../../data/purchases.csv",
            expectedCode: http.StatusOK,
            model:        &models.Purchase{},
        },
    }

    for _, test := range tests {
        t.Run(test.fileName, func(t *testing.T) {
            // Open the test file
            file, err := os.Open(test.filePath)
            assert.NoError(t, err)
            defer file.Close()

            // Create a multipart form file
            body := &bytes.Buffer{}
            writer := multipart.NewWriter(body)
            part, err := writer.CreateFormFile("file", filepath.Base(test.filePath))
            assert.NoError(t, err)

            _, err = file.Seek(0, 0)
            assert.NoError(t, err)

            _, err = io.Copy(part, file)
            assert.NoError(t, err)
            writer.Close()

            // Create a test HTTP request
            req := httptest.NewRequest("POST", "/import", body)
            req.Header.Set("Content-Type", writer.FormDataContentType())

            // Perform the request
            w := httptest.NewRecorder()
            router.ServeHTTP(w, req)

            // Assert the response code
            assert.Equal(t, test.expectedCode, w.Code)

            // Assert that data was imported into the database
            var count int64
            database.DB.Model(test.model).Count(&count)
            assert.Greater(t, count, int64(0), "Expected data to be imported into the database")
        })
    }
}