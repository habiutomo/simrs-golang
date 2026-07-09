package models

import "time"

type Billing struct {
	ID                uint          `gorm:"primaryKey" json:"id"`
	InvoiceNumber     string        `gorm:"uniqueIndex;size:30;not null" json:"invoice_number"`
	PatientID         uint          `gorm:"not null;index" json:"patient_id"`
	Patient           Patient       `gorm:"foreignKey:PatientID" json:"patient"`
	AppointmentID     *uint         `gorm:"index" json:"appointment_id"`
	Appointment       *Appointment  `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	InpatientID       *uint         `gorm:"index" json:"inpatient_id"`
	Inpatient         *Inpatient    `gorm:"foreignKey:InpatientID" json:"inpatient,omitempty"`
	TotalAmount       float64       `gorm:"not null;default:0" json:"total_amount"`
	PaymentMethod     string        `gorm:"size:30" json:"payment_method"`
	PaymentStatus     string        `gorm:"size:20;not null;default:'unpaid'" json:"payment_status"` // unpaid, paid, partial
	PaidAmount        float64       `gorm:"default:0" json:"paid_amount"`
	BillingDate       time.Time     `gorm:"not null" json:"billing_date"`
	Notes             string        `gorm:"size:255" json:"notes"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Items             []BillingItem `gorm:"foreignKey:BillingID" json:"items,omitempty"`
}

func (Billing) TableName() string {
	return "billings"
}

type BillingItem struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	BillingID   uint       `gorm:"not null;index" json:"billing_id"`
	ItemType    string     `gorm:"size:30;not null" json:"item_type"` // consultation, medication, room, procedure
	ItemName    string     `gorm:"size:150;not null" json:"item_name"`
	Quantity    int        `gorm:"not null;default:1" json:"quantity"`
	UnitPrice   float64    `gorm:"not null" json:"unit_price"`
	TotalPrice  float64    `gorm:"not null" json:"total_price"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (BillingItem) TableName() string {
	return "billing_items"
}
