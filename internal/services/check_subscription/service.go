package checksubscription

import (
	"context"
	"subscriptions/internal/data"
	repository "subscriptions/internal/repository/subscriptions"
)

type CheckSubscriptionUseCase interface {
	Execute(ctx context.Context, userID string) (*data.Subscription, error)
}

type checkSubscriptionUseCase struct {
	repo repository.SubscriptionRepository
}

func NewCheckSubscriptionUseCase(repo repository.SubscriptionRepository) CheckSubscriptionUseCase {
	return &checkSubscriptionUseCase{repo: repo}
}

func (uc *checkSubscriptionUseCase) Execute(ctx context.Context, userID string) (*data.Subscription, error) {
	return uc.repo.GetActiveSubscription(ctx, userID)
}
