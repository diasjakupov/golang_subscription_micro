package plans

import (
	"subscriptions/internal/app/connections"
	"subscriptions/internal/data"

	"gorm.io/gorm"
)

// PlanRepository defines the interface for subscription plan data operations.
type PlanRepository interface {
	GetAllPlans() ([]data.SubscriptionPlan, error)
	GetPlanByID(id string) (*data.SubscriptionPlan, error)
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
func (repo *DBPlanRepository) GetAllPlans() ([]data.SubscriptionPlan, error) {
	var plans []data.SubscriptionPlan
	result := repo.db.Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}
	return plans, nil
}

// GetPlanByID retrieves a subscription plan with the specified id from the database.
// It returns an error if no plan is found.
func (repo *DBPlanRepository) GetPlanByID(id string) (*data.SubscriptionPlan, error) {
	var plan data.SubscriptionPlan
	result := repo.db.First(&plan, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &plan, nil
}
