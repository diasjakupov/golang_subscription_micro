package subscriptions

import "time"

// Subscription represents a user's subscription
type Subscription struct {
	ID          string     `db:"id"`
	UserID      string     `db:"user_id"` // Reference to the User Service
	PlanID      string     `db:"plan_id"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     time.Time  `db:"end_date"`
	Status      string     `db:"status"` // active, cancelled, expired, pending
	AutoRenew   bool       `db:"auto_renew"`
	CancelledAt *time.Time `db:"cancelled_at,omitempty"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}
