package createsubscription

import (
	"errors"
	"subscriptions/internal/data"
	repository "subscriptions/internal/repository/subscriptions"
	messagequeue "subscriptions/internal/services/message_queue"
	"subscriptions/internal/services/payment"
	"time"

	"github.com/google/uuid"
)

type CreateSubscriptionUseCase interface {
	Execute(req data.CreateSubscriptionRequest) (*data.Subscription, error)
}

type createSubscriptionUseCase struct {
	repo           repository.SubscriptionRepository
	paymentService payment.PaymentService
	messageQueue   messagequeue.MessageQueue
}

func NewCreateSubscriptionUseCase(repo repository.SubscriptionRepository, ps payment.PaymentService, mq messagequeue.MessageQueue) CreateSubscriptionUseCase {
	return &createSubscriptionUseCase{
		repo:           repo,
		paymentService: ps,
		messageQueue:   mq,
	}
}

func (uc *createSubscriptionUseCase) Execute(req data.CreateSubscriptionRequest) (*data.Subscription, error) {
	// Check if an active subscription exists.
	if sub, err := uc.repo.GetActiveSubscription(req.UserID); err == nil && sub != nil {
		return nil, errors.New("already subscribed")
	}
	// Process payment.
	payment, err := uc.paymentService.ProcessPayment(req.PaymentDetails)
	if err != nil {
		return nil, err
	}
	// Create new subscription (using a 30-day duration for demo purposes).
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 30)
	sub := &data.Subscription{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		PlanID:    req.PlanID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    "active",
		AutoRenew: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := uc.repo.SaveSubscription(sub); err != nil {
		return nil, errors.New("failed to create subscription")
	}
	// Associate payment with the subscription.
	payment.SubscriptionID = sub.ID
	// Publish event.
	uc.messageQueue.Publish("SubscriptionCreated", sub)
	return sub, nil
}
