package database

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "phorest-techtest/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("phorest.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = DB.AutoMigrate(&models.Client{}, &models.Appointment{}, &models.Service{}, &models.Purchase{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}