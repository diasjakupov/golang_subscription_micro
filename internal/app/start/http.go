package start

import (
	"fmt"
	"log"
	"net/http"
	"subscriptions/internal/app/config"
	"subscriptions/internal/delieveries"
	planRepository "subscriptions/internal/repository/plans"
	subRepository "subscriptions/internal/repository/subscriptions"

	cancelsubscription "subscriptions/internal/services/cancel_subscription"
	checksubscription "subscriptions/internal/services/check_subscription"
	createsubscription "subscriptions/internal/services/create_subscription"
	messagequeue "subscriptions/internal/services/message_queue"
	"subscriptions/internal/services/payment"
	"subscriptions/internal/services/plans"
	renewsubscription "subscriptions/internal/services/renew_subscription"
)

func HTTP(cfg config.HTTPServerConfig) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	repo := subRepository.NewMemorySubscriptionRepository()
	paymentService := &payment.DummyPaymentService{}
	messageQueue := &messagequeue.DummyMessageQueue{}

	createUC := createsubscription.NewCreateSubscriptionUseCase(repo, paymentService, messageQueue)
	checkUC := checksubscription.NewCheckSubscriptionUseCase(repo)
	cancelUC := cancelsubscription.NewCancelSubscriptionUseCase(repo, messageQueue)
	renewUC := renewsubscription.NewRenewSubscriptionUseCase(repo, paymentService, messageQueue)

	planRepo := planRepository.NewMemoryPlanRepository()
	getPlansUC := plans.NewGetSubscriptionPlansUseCase(planRepo)

	handler := delieveries.NewHandler(createUC, checkUC, cancelUC, renewUC, getPlansUC)

	router := delieveries.NewRouter(handler)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Starting HTTP server on %s", addr)

	router.Start(addr)
}
