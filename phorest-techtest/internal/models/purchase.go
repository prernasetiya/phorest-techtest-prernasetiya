package models

type Purchase struct {
    ID             string  `gorm:"primaryKey"` // ID as the primary key
    AppointmentID  string  `gorm:"not null"`   // Foreign key to Appointment
    Name           string  `gorm:"not null"`   // Name of the purchased item
    Price          float64 `gorm:"not null"`   // Price of the purchased item
    LoyaltyPoints  int     `gorm:"not null"`   // Loyalty points associated with the purchase
}