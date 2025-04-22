package repository

import (
	"context"
	"subscriptions/internal/data"

	"gorm.io/gorm"
)

// SubscriptionRepository defines the interface for subscription data operations.
type SubscriptionRepository interface {
	// GetActiveSubscription retrieves the active subscription for a given user.
	GetActiveSubscription(ctx context.Context, userID string) (*data.Subscription, error)
	// SaveSubscription creates a new subscription in the database.
	// This operation is wrapped in a transaction.
	SaveSubscription(ctx context.Context, sub *data.Subscription) error
	// UpdateSubscription updates an existing subscription in the database.
	// This operation is wrapped in a transaction.
	UpdateSubscription(ctx context.Context, sub *data.Subscription) error
}

// DBSubscriptionRepository implements SubscriptionRepository using GORM.
type DBSubscriptionRepository struct {
	db *gorm.DB
}

// NewDBSubscriptionRepository initializes a new repository using the provided database connection.
func NewDBSubscriptionRepository(conn *gorm.DB) *DBSubscriptionRepository {
	return &DBSubscriptionRepository{
		db: conn,
	}
}

// GetActiveSubscription retrieves a user's active subscription by filtering on user_id and "active" status.
func (repo *DBSubscriptionRepository) GetActiveSubscription(ctx context.Context, userID string) (*data.Subscription, error) {
	var sub data.Subscription
	// Perform a query to find the subscription that is active for the specified user.
	err := repo.db.WithContext(ctx).Where("user_id = ? AND status = ?", userID, "active").First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// SaveSubscription creates a new subscription record in the database within a transaction.
func (repo *DBSubscriptionRepository) SaveSubscription(ctx context.Context, sub *data.Subscription) error {
	return repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create the new subscription record.
		if err := tx.Create(sub).Error; err != nil {
			return err
		}
		// Additional steps could be added here if necessary.
		return nil
	})
}

// UpdateSubscription updates an existing subscription record in the database within a transaction.
func (repo *DBSubscriptionRepository) UpdateSubscription(ctx context.Context, sub *data.Subscription) error {
	return repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Save will perform an update if the record already exists.
		if err := tx.Save(sub).Error; err != nil {
			return err
		}
		return nil
	})
}
