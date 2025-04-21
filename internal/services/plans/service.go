package plans

import (
	"context"
	"subscriptions/internal/data"
	repository "subscriptions/internal/repository/plans"
)

type GetSubscriptionPlansUseCase interface {
	Execute(ctx context.Context) ([]data.SubscriptionPlan, error)
}

type getSubscriptionPlansUseCase struct {
	planRepo repository.PlanRepository
}

func NewGetSubscriptionPlansUseCase(planRepo repository.PlanRepository) GetSubscriptionPlansUseCase {
	return &getSubscriptionPlansUseCase{
		planRepo: planRepo,
	}
}

func (uc *getSubscriptionPlansUseCase) Execute(ctx context.Context) ([]data.SubscriptionPlan, error) {
	return uc.planRepo.GetAllPlans(ctx)
}
