package services

import (
	"fmt"
	"time"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type BillingService struct {
	repo             *repositories.BillingRepository
	patientRepo      *repositories.PatientRepository
}

func NewBillingService(repo *repositories.BillingRepository, patientRepo *repositories.PatientRepository) *BillingService {
	return &BillingService{repo: repo, patientRepo: patientRepo}
}

func generateInvoiceNumber() string {
	now := time.Now()
	return fmt.Sprintf("INV-%s-%d", now.Format("20060102"), now.UnixMilli()%100000)
}

func (s *BillingService) Create(req *dto.CreateBillingRequest) (*models.Billing, error) {
	var totalAmount float64
	var items []models.BillingItem

	for _, item := range req.Items {
		totalPrice := float64(item.Quantity) * item.UnitPrice
		totalAmount += totalPrice
		items = append(items, models.BillingItem{
			ItemType:   item.ItemType,
			ItemName:   item.ItemName,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: totalPrice,
		})
	}

	billing := &models.Billing{
		InvoiceNumber: generateInvoiceNumber(),
		PatientID:     req.PatientID,
		AppointmentID: req.AppointmentID,
		InpatientID:   req.InpatientID,
		TotalAmount:   totalAmount,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "unpaid",
		PaidAmount:    0,
		BillingDate:   time.Now(),
		Notes:         req.Notes,
		Items:         items,
	}

	err := s.repo.Create(billing)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(billing.ID)
}

func (s *BillingService) GetByID(id uint) (*models.Billing, error) {
	return s.repo.FindByID(id)
}

func (s *BillingService) GetAll(page, limit int, status string) ([]models.Billing, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.FindAll(page, limit, status)
}

func (s *BillingService) GetByPatient(patientID uint) ([]models.Billing, error) {
	return s.repo.FindByPatientID(patientID)
}

func (s *BillingService) Pay(id uint, req *dto.PayBillingRequest) (*models.Billing, error) {
	billing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("billing not found")
	}

	newPaid := billing.PaidAmount + req.Amount
	if newPaid >= billing.TotalAmount {
		billing.PaymentStatus = "paid"
		billing.PaidAmount = billing.TotalAmount
	} else {
		billing.PaymentStatus = "partial"
		billing.PaidAmount = newPaid
	}

	if req.PaymentMethod != "" {
		billing.PaymentMethod = req.PaymentMethod
	}

	err = s.repo.Update(billing)
	return billing, err
}

func (s *BillingService) GetRevenue() (float64, error) {
	return s.repo.GetTotalRevenue()
}

func (s *BillingService) GetUnpaidTotal() (float64, error) {
	return s.repo.GetTotalUnpaid()
}
