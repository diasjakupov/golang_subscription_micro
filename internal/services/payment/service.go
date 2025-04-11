package payment

import (
	"errors"
	"math/rand"
	"subscriptions/internal/data"
	"time"

	"github.com/google/uuid"
)

// PaymentService defines operations for processing payments.
type PaymentService interface {
	ProcessPayment(details data.PaymentDetails) (data.Payment, error)
	ProcessRenewalPayment(details data.PaymentDetails) (data.Payment, error)
}

// DummyPaymentService is a dummy implementation of PaymentService.
type DummyPaymentService struct{}

// ProcessPayment attempts to process a payment with the provided details.
// It creates a Payment model using the embedded BaseModel for IDs and timestamps.
func (d *DummyPaymentService) ProcessPayment(details data.PaymentDetails) (data.Payment, error) {
	// Simulate potential payment failure based on amount and random chance.
	if details.Amount <= 0 || rand.Intn(10) < 2 {
		return data.Payment{}, errors.New("payment failed: insufficient funds or invalid amount")
	}

	now := time.Now()
	payment := data.Payment{
		BaseModel: data.BaseModel{
			ID:        uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Amount:        details.Amount,
		Currency:      "USD",
		PaymentMethod: details.PaymentMethod,
		Status:        "successful",
		TransactionID: uuid.New().String(),
		PaymentDate:   now,
	}

	return payment, nil
}

// ProcessRenewalPayment processes a renewal payment using the same logic as ProcessPayment.
func (d *DummyPaymentService) ProcessRenewalPayment(details data.PaymentDetails) (data.Payment, error) {
	return d.ProcessPayment(details)
}
