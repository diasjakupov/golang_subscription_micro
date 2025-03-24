package plans

import (
	"subscriptions/internal/data"
	"time"
)

type PlanRepository interface {
	GetAllPlans() ([]data.SubscriptionPlan, error)
}

// MemoryPlanRepository is an in-memory implementation with mock data.
type MemoryPlanRepository struct {
	plans []data.SubscriptionPlan
}

func NewMemoryPlanRepository() *MemoryPlanRepository {
	now := time.Now()
	return &MemoryPlanRepository{
		plans: []data.SubscriptionPlan{
			{
				ID:           "plan_monthly",
				Name:         "Monthly Plan",
				Description:  "Access to basic features on a monthly basis.",
				Price:        9.99,
				BillingCycle: "monthly",
				DurationDays: 30,
				IsActive:     true,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			{
				ID:           "plan_quarterly",
				Name:         "Quarterly Plan",
				Description:  "Access to basic features on a quarterly basis.",
				Price:        27.99,
				BillingCycle: "quarterly",
				DurationDays: 90,
				IsActive:     true,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			{
				ID:           "plan_yearly",
				Name:         "Yearly Plan",
				Description:  "Access to basic features on a yearly basis.",
				Price:        99.99,
				BillingCycle: "yearly",
				DurationDays: 365,
				IsActive:     true,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		},
	}
}

func (repo *MemoryPlanRepository) GetAllPlans() ([]data.SubscriptionPlan, error) {
	return repo.plans, nil
}
