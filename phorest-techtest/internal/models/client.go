package models

type Client struct {
    ID          string        `gorm:"primaryKey" csv:"id"` // ID as the primary key
    FirstName   string        `gorm:"not null" csv:"first_name"`
    LastName    string        `gorm:"not null" csv:"last_name"`
    Email       string        `gorm:"not null;unique" csv:"email"`
    Phone       string        `gorm:"not null" csv:"phone"`
    Gender      string        `gorm:"not null" csv:"gender"`
    Banned      bool          `gorm:"not null" csv:"banned"`
    Appointments []Appointment `gorm:"foreignKey:ClientID"` // One-to-many relationship
}