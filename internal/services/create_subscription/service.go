package createsubscription

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"subscriptions/internal/data"
	plansRepository "subscriptions/internal/repository/plans"
	repository "subscriptions/internal/repository/subscriptions"
	messagequeue "subscriptions/internal/services/message_queue"
	"subscriptions/internal/services/payment"
)

// CreateSubscriptionUseCase defines the interface for creating a subscription.
type CreateSubscriptionUseCase interface {
	Execute(req data.CreateSubscriptionRequest) (*data.Subscription, error)
}

// createSubscriptionUseCase is the concrete implementation of CreateSubscriptionUseCase.
type createSubscriptionUseCase struct {
	repo           repository.SubscriptionRepository
	paymentService payment.PaymentService
	messageQueue   messagequeue.MessageQueue
	planRepo       plansRepository.PlanRepository
}

// NewCreateSubscriptionUseCase constructs a new CreateSubscriptionUseCase.
func NewCreateSubscriptionUseCase(
	repo repository.SubscriptionRepository,
	ps payment.PaymentService,
	mq messagequeue.MessageQueue,
	pr plansRepository.PlanRepository,
) CreateSubscriptionUseCase {
	return &createSubscriptionUseCase{
		repo:           repo,
		paymentService: ps,
		messageQueue:   mq,
		planRepo:       pr,
	}
}

// Execute creates a new subscription after verifying that the user is not already subscribed.
// It processes the payment, retrieves the plan details to set the subscription duration, saves the subscription,
// associates the payment with the subscription, and publishes an event.
func (uc *createSubscriptionUseCase) Execute(req data.CreateSubscriptionRequest) (*data.Subscription, error) {
	// Check if an active subscription already exists.
	if existingSub, err := uc.repo.GetActiveSubscription(req.UserID); err == nil && existingSub != nil {
		return nil, errors.New("already subscribed")
	}

	// Process the payment.
	paymentResult, err := uc.paymentService.ProcessPayment(req.PaymentDetails)
	if err != nil {
		return nil, err
	}

	// Retrieve plan details to determine subscription duration.
	plan, err := uc.planRepo.GetPlanByID(req.PlanID)
	if err != nil {
		return nil, errors.New("failed to fetch plan details: " + err.Error())
	}

	now := time.Now()
	// Create a new subscription with duration based on the plan's DurationDays using the embedded BaseModel.
	sub := &data.Subscription{
		BaseModel: data.BaseModel{
			ID:        uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
		},
		UserID:    req.UserID,
		PlanID:    req.PlanID,
		StartDate: now,
		EndDate:   now.AddDate(0, 0, plan.DurationDays),
		Status:    "active",
		AutoRenew: true,
	}

	if err := uc.repo.SaveSubscription(sub); err != nil {
		return nil, errors.New("failed to create subscription: " + err.Error())
	}

	// Associate the processed payment with the new subscription.
	paymentResult.SubscriptionID = sub.ID

	// Publish the subscription creation event.
	uc.messageQueue.Publish("SubscriptionCreated", sub)

	return sub, nil
}
