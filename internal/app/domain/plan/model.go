package plan

import "time"

// SubscriptionPlan represents different subscription tiers
type SubscriptionPlan struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	Price        float64   `db:"price"`
	BillingCycle string    `db:"billing_cycle"` // monthly, quarterly, yearly
	DurationDays int       `db:"duration_days"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
