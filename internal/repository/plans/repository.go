package plans

import (
	"context"
	"subscriptions/internal/app/connections"
	"subscriptions/internal/data"

	"gorm.io/gorm"
)

// PlanRepository defines the interface for subscription plan data operations.
type PlanRepository interface {
	GetAllPlans(ctx context.Context) ([]data.SubscriptionPlan, error)
	GetPlanByID(ctx context.Context, d string) (*data.SubscriptionPlan, error)
}

// DBPlanRepository implements PlanRepository using a GORM database connection.
type DBPlanRepository struct {
	db *gorm.DB
}

// NewDBPlanRepository initializes a new repository using a real database connection from the connections package.
func NewDBPlanRepository(conn *connections.Connections) *DBPlanRepository {
	return &DBPlanRepository{
		db: conn.DB,
	}
}

// GetAllPlans retrieves all available subscription plans from the database.
func (repo *DBPlanRepository) GetAllPlans(ctx context.Context) ([]data.SubscriptionPlan, error) {
	var plans []data.SubscriptionPlan
	result := repo.db.WithContext(ctx).Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}
	return plans, nil
}

// GetPlanByID retrieves a subscription plan with the specified id from the database.
// It returns an error if no plan is found.
func (repo *DBPlanRepository) GetPlanByID(ctx context.Context, id string) (*data.SubscriptionPlan, error) {
	var plan data.SubscriptionPlan
	result := repo.db.WithContext(ctx).First(&plan, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &plan, nil
}
