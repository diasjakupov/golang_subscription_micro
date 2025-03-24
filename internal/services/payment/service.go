package payment

import (
	"errors"
	"math/rand"
	"subscriptions/internal/data"
	"time"

	"github.com/google/uuid"
)

type PaymentService interface {
	ProcessPayment(details data.PaymentDetails) (data.Payment, error)
	ProcessRenewalPayment(details data.PaymentDetails) (data.Payment, error)
}

type DummyPaymentService struct{}

func (d *DummyPaymentService) ProcessPayment(details data.PaymentDetails) (data.Payment, error) {
	if details.Amount <= 0 || rand.Intn(10) < 2 {
		return data.Payment{}, errors.New("payment failed: insufficient funds or invalid amount")
	}
	payment := data.Payment{
		ID:            uuid.New().String(),
		Amount:        details.Amount,
		Currency:      "USD",
		PaymentMethod: details.PaymentMethod,
		Status:        "successful",
		TransactionID: uuid.New().String(),
		PaymentDate:   time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	return payment, nil
}

func (d *DummyPaymentService) ProcessRenewalPayment(details data.PaymentDetails) (data.Payment, error) {
	return d.ProcessPayment(details)
}
