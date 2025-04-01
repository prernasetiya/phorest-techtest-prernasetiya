package models

type Appointment struct {
    ID         string     `gorm:"primaryKey" csv:"id"` // ID as the primary key
    ClientID   string     `gorm:"not null" csv:"client_id"`   // Foreign key to Client
    StartTime  string     `gorm:"not null" csv:"start_time"`   // Start time of the appointment
    EndTime    string     `gorm:"not null" csv:"end_time"`   // End time of the appointment
    Services   []Service  `gorm:"foreignKey:AppointmentID"` // One-to-many relationship with Service
    Purchases  []Purchase `gorm:"foreignKey:AppointmentID"` // One-to-many relationship with Purchase
}