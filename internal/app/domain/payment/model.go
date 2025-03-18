package payment

import "time"

// Payment represents a record of a subscription payment transaction
type Payment struct {
	ID             string    `db:"id"`
	SubscriptionID string    `db:"subscription_id"` // Foreign key referencing subscriptions.id
	Amount         float64   `db:"amount"`
	Currency       string    `db:"currency"`       // Default: "USD"
	PaymentMethod  string    `db:"payment_method"` // e.g., "credit_card", "paypal"
	Status         string    `db:"status"`         // Possible values: "pending", "successful", "failed"
	TransactionID  string    `db:"transaction_id"` // External payment processor transaction ID
	PaymentDate    time.Time `db:"payment_date"`   // Timestamp of payment
	FailureReason  string    `db:"failure_reason"` // If payment failed, reason will be stored
	CreatedAt      time.Time `db:"created_at"`     // Default: NOW()
	UpdatedAt      time.Time `db:"updated_at"`     // Default: NOW()
}
