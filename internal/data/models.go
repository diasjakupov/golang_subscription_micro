package data

import "time"

// BaseModel provides common fields for most data models.
type BaseModel struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Invoice represents a billing invoice for a subscription.
type Invoice struct {
	BaseModel
	SubscriptionID string     `db:"subscription_id"` // Foreign key referencing subscriptions.id
	PaymentID      *string    `db:"payment_id"`      // Nullable, references payments.id
	Amount         float64    `db:"amount"`
	Status         string     `db:"status"`    // "draft", "pending", "paid", "cancelled"
	DueDate        time.Time  `db:"due_date"`  // When the invoice is due
	PaidDate       *time.Time `db:"paid_date"` // Nullable, timestamp when paid
}

// Payment represents a record of a subscription payment transaction.
type Payment struct {
	BaseModel
	SubscriptionID string    `db:"subscription_id"` // Foreign key referencing subscriptions.id
	Amount         float64   `db:"amount"`
	Currency       string    `db:"currency"`       // Default: "USD"
	PaymentMethod  string    `db:"payment_method"` // e.g., "credit_card", "paypal"
	Status         string    `db:"status"`         // Possible values: "pending", "successful", "failed"
	TransactionID  string    `db:"transaction_id"` // External payment processor transaction ID
	PaymentDate    time.Time `db:"payment_date"`   // Timestamp of payment
	FailureReason  string    `db:"failure_reason"` // If payment failed, reason will be stored
}

// SubscriptionPlan represents different subscription tiers.
type SubscriptionPlan struct {
	BaseModel
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Price        float64 `db:"price"`
	BillingCycle string  `db:"billing_cycle"` // monthly, quarterly, yearly
	DurationDays int     `db:"duration_days"`
	IsActive     bool    `db:"is_active"`
}

// Subscription represents a user's subscription.
type Subscription struct {
	BaseModel
	UserID      string     `db:"user_id"` // Reference to the User Service
	PlanID      string     `db:"plan_id"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     time.Time  `db:"end_date"`
	Status      string     `db:"status"` // active, cancelled, expired, pending
	AutoRenew   bool       `db:"auto_renew"`
	CancelledAt *time.Time `db:"cancelled_at,omitempty"`
}

// SubscriptionEvent represents events related to subscriptions.
// It is kept separate as it only records the creation time of the event.
type SubscriptionEvent struct {
	ID             string    `db:"id"`
	SubscriptionID string    `db:"subscription_id"`
	EventType      string    `db:"event_type"` // created, renewed, cancelled, expired
	EventData      string    `db:"event_data"` // JSON data
	CreatedAt      time.Time `db:"created_at"`
}
