package invoice

import "time"

// Invoice represents a billing invoice for a subscription
type Invoice struct {
	ID             string     `db:"id"`
	SubscriptionID string     `db:"subscription_id"` // Foreign key referencing subscriptions.id
	PaymentID      *string    `db:"payment_id"`      // Nullable, references payments.id
	Amount         float64    `db:"amount"`
	Status         string     `db:"status"`     // "draft", "pending", "paid", "cancelled"
	DueDate        time.Time  `db:"due_date"`   // When the invoice is due
	PaidDate       *time.Time `db:"paid_date"`  // Nullable, timestamp when paid
	CreatedAt      time.Time  `db:"created_at"` // Default: NOW()
	UpdatedAt      time.Time  `db:"updated_at"` // Default: NOW()
}
