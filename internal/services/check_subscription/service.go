package checksubscription

import (
	"subscriptions/internal/data"
	repository "subscriptions/internal/repository/subscriptions"
)

type CheckSubscriptionUseCase interface {
	Execute(userID string) (*data.Subscription, error)
}

type checkSubscriptionUseCase struct {
	repo repository.SubscriptionRepository
}

func NewCheckSubscriptionUseCase(repo repository.SubscriptionRepository) CheckSubscriptionUseCase {
	return &checkSubscriptionUseCase{repo: repo}
}

func (uc *checkSubscriptionUseCase) Execute(userID string) (*data.Subscription, error) {
	return uc.repo.GetActiveSubscription(userID)
}
