package subscriptionevent

import "time"

// SubscriptionEvent represents events related to subscriptions
type SubscriptionEvent struct {
	ID             string    `db:"id"`
	SubscriptionID string    `db:"subscription_id"`
	EventType      string    `db:"event_type"` // created, renewed, cancelled, expired
	EventData      string    `db:"event_data"` // JSON data
	CreatedAt      time.Time `db:"created_at"`
}
