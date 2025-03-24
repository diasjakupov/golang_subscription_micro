package plans

import (
	"subscriptions/internal/data"
	repository "subscriptions/internal/repository/plans"
)

type GetSubscriptionPlansUseCase interface {
	Execute() ([]data.SubscriptionPlan, error)
}

type getSubscriptionPlansUseCase struct {
	planRepo repository.PlanRepository
}

func NewGetSubscriptionPlansUseCase(planRepo repository.PlanRepository) GetSubscriptionPlansUseCase {
	return &getSubscriptionPlansUseCase{
		planRepo: planRepo,
	}
}

func (uc *getSubscriptionPlansUseCase) Execute() ([]data.SubscriptionPlan, error) {
	return uc.planRepo.GetAllPlans()
}
