package renewsubscription

import (
	"errors"
	"subscriptions/internal/data"
	repository "subscriptions/internal/repository/subscriptions"
	messagequeue "subscriptions/internal/services/message_queue"
	"subscriptions/internal/services/payment"
	"time"
)

type RenewSubscriptionUseCase interface {
	Execute(req data.RenewSubscriptionRequest) error
}

type renewSubscriptionUseCase struct {
	repo           repository.SubscriptionRepository
	paymentService payment.PaymentService
	messageQueue   messagequeue.MessageQueue
}

func NewRenewSubscriptionUseCase(repo repository.SubscriptionRepository, ps payment.PaymentService, mq messagequeue.MessageQueue) RenewSubscriptionUseCase {
	return &renewSubscriptionUseCase{
		repo:           repo,
		paymentService: ps,
		messageQueue:   mq,
	}
}

func (uc *renewSubscriptionUseCase) Execute(req data.RenewSubscriptionRequest) error {
	sub, err := uc.repo.GetActiveSubscription(req.UserID)
	if err != nil {
		return err
	}
	payment, err := uc.paymentService.ProcessRenewalPayment(req.PaymentDetails)
	if err != nil {
		return err
	}
	sub.EndDate = sub.EndDate.AddDate(0, 0, 30)
	sub.UpdatedAt = time.Now()
	if err := uc.repo.UpdateSubscription(sub); err != nil {
		return errors.New("failed to renew subscription")
	}
	uc.messageQueue.Publish("SubscriptionRenewed", map[string]interface{}{
		"subscription": sub,
		"payment":      payment,
	})
	return nil
}
