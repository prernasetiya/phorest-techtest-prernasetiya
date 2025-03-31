package models

type Appointment struct {
    ID         string     `gorm:"primaryKey"` // ID as the primary key
    ClientID   string     `gorm:"not null"`   // Foreign key to Client
    StartTime  string     `gorm:"not null"`   // Start time of the appointment
    EndTime    string     `gorm:"not null"`   // End time of the appointment
    Services   []Service  `gorm:"foreignKey:AppointmentID"` // One-to-many relationship with Service
    Purchases  []Purchase `gorm:"foreignKey:AppointmentID"` // One-to-many relationship with Purchase
}