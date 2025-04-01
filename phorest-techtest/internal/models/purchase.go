package models

type Purchase struct {
    ID             string  `gorm:"primaryKey" csv:"id"` // ID as the primary key
    AppointmentID  string  `gorm:"not null" csv:"appointment_id"`   // Foreign key to Appointment
    Name           string  `gorm:"not null" csv:name"`   // Name of the purchased item
    Price          float64 `gorm:"not null" csv:"price"`   // Price of the purchased item
    LoyaltyPoints  int     `gorm:"not null" csv:"loyalty_points"`   // Loyalty points associated with the purchase
}