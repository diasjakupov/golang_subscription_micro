package renewsubscription

import (
	"context"
	"errors"
	"subscriptions/internal/data"
	plansRepository "subscriptions/internal/repository/plans"
	repository "subscriptions/internal/repository/subscriptions"
	messagequeue "subscriptions/internal/services/message_queue"
	"subscriptions/internal/services/payment"
	"time"
)

type RenewSubscriptionUseCase interface {
	Execute(ctx context.Context, req data.RenewSubscriptionRequest) error
}

type renewSubscriptionUseCase struct {
	repo           repository.SubscriptionRepository
	paymentService payment.PaymentService
	messageQueue   messagequeue.MessageQueue
	planRepo       plansRepository.PlanRepository
}

func NewRenewSubscriptionUseCase(repo repository.SubscriptionRepository, pr plansRepository.PlanRepository,
	ps payment.PaymentService, mq messagequeue.MessageQueue) RenewSubscriptionUseCase {
	return &renewSubscriptionUseCase{
		repo:           repo,
		paymentService: ps,
		messageQueue:   mq,
		planRepo:       pr,
	}
}

func (uc *renewSubscriptionUseCase) Execute(ctx context.Context, req data.RenewSubscriptionRequest) error {
	sub, err := uc.repo.GetActiveSubscription(ctx, req.UserID)
	if err != nil {
		return err
	}
	payment, paymentErr := uc.paymentService.ProcessRenewalPayment(req.PaymentDetails)
	if paymentErr != nil {
		return paymentErr
	}

	plan, planErr := uc.planRepo.GetPlanByID(ctx, sub.PlanID)
	if planErr != nil {
		return planErr
	}

	sub.EndDate = sub.EndDate.AddDate(0, 0, plan.DurationDays)
	sub.UpdatedAt = time.Now()
	if err := uc.repo.UpdateSubscription(ctx, sub); err != nil {
		return errors.New("failed to renew subscription")
	}
	uc.messageQueue.Publish("SubscriptionRenewed", map[string]interface{}{
		"subscription": sub,
		"payment":      payment,
	})
	return nil
}
