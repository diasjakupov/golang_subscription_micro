package start

import (
	"fmt"
	"log"
	"net/http"
	"subscriptions/internal/app/config"
	"subscriptions/internal/app/domain/plan"
)

func HTTP(cfg config.HTTPServerConfig) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	planHandler := plan.NewPlanHandler()
	http.HandleFunc("/plans", planHandler.GetAllPlans)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Starting HTTP server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
