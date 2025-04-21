package start

import (
	"fmt"
	"log"
	"net/http"

	"subscriptions/internal/app/config"
	"subscriptions/internal/app/connections"
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

// HTTP initializes and starts the HTTP server.
// It sets up the dependencies including a database connection,
// repositories, services, and routes, then starts listening on the configured address.
func HTTP(conn *connections.Connections, cfg *config.HTTPServerConfig) {
	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	// Initialize repositories with the database connection.
	repo := subRepository.NewDBSubscriptionRepository(conn)
	planRepo := planRepository.NewDBPlanRepository(conn)

	// Initialize services
	paymentService := &payment.DummyPaymentService{}
	messageQueue := &messagequeue.DummyMessageQueue{}

	// Create use cases with dependency injection
	createUC := createsubscription.NewCreateSubscriptionUseCase(repo, paymentService, messageQueue, planRepo)
	checkUC := checksubscription.NewCheckSubscriptionUseCase(repo)
	cancelUC := cancelsubscription.NewCancelSubscriptionUseCase(repo, messageQueue)
	renewUC := renewsubscription.NewRenewSubscriptionUseCase(repo, planRepo, paymentService, messageQueue)
	getPlansUC := plans.NewGetSubscriptionPlansUseCase(planRepo)

	// Set up the HTTP handler with all use cases
	handler := delieveries.NewHandler(createUC, checkUC, cancelUC, renewUC, getPlansUC)

	// Create the router and start the HTTP server
	router := delieveries.NewRouter(handler)
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Starting HTTP server on %s", addr)

	router.Start(addr)
}
