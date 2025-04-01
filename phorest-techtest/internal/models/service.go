package models

type Service struct {
    ID             string  `gorm:"primaryKey" csv:"id"` // ID as the primary key
    AppointmentID  string  `gorm:"not null" csv:"appointment_id"`   // Foreign key to Appointment
    Name           string  `gorm:"not null" csv:"name"`   // Name of the service
    Price          float64 `gorm:"not null" csv:"price"`   // Price of the service
    LoyaltyPoints  int     `gorm:"not null" csv:"loyalty_points"`   // Loyalty points associated with the service
}