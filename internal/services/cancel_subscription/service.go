package cancelsubscription

import (
	"errors"
	repository "subscriptions/internal/repository/subscriptions"
	messagequeue "subscriptions/internal/services/message_queue"
	"time"
)

type CancelSubscriptionUseCase interface {
	Execute(userID string) error
}

type cancelSubscriptionUseCase struct {
	repo         repository.SubscriptionRepository
	messageQueue messagequeue.MessageQueue
}

func NewCancelSubscriptionUseCase(repo repository.SubscriptionRepository, mq messagequeue.MessageQueue) CancelSubscriptionUseCase {
	return &cancelSubscriptionUseCase{repo: repo, messageQueue: mq}
}

func (uc *cancelSubscriptionUseCase) Execute(userID string) error {
	sub, err := uc.repo.GetActiveSubscription(userID)
	if err != nil {
		return err
	}
	now := time.Now()
	sub.Status = "cancelled"
	sub.CancelledAt = &now
	sub.UpdatedAt = now
	if err := uc.repo.UpdateSubscription(sub); err != nil {
		return errors.New("failed to cancel subscription")
	}
	uc.messageQueue.Publish("SubscriptionCancelled", sub)
	return nil
}
