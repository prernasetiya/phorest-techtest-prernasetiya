package models

type Client struct {
    ID          string        `gorm:"primaryKey"` // ID as the primary key
    FirstName   string        `gorm:"not null"`
    LastName    string        `gorm:"not null"`
    Email       string        `gorm:"not null;unique"`
    Phone       string        `gorm:"not null"`
    Gender      string        `gorm:"not null"`
    Banned      bool          `gorm:"not null"`
    Appointments []Appointment `gorm:"foreignKey:ClientID"` // One-to-many relationship
}