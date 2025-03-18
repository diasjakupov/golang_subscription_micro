package plan

import (
	"encoding/json"
	"net/http"
	"time"
)

type PlanHandlerHttp struct{}

func NewPlanHandler() PlanHandler {
	return &PlanHandlerHttp{}
}

func (handler *PlanHandlerHttp) GetAllPlans(w http.ResponseWriter, r *http.Request) {
	plans := []SubscriptionPlan{
		{
			ID:           "plan_001",
			Name:         "Basic Plan",
			Description:  "This plan offers basic features.",
			Price:        9.99,
			BillingCycle: "monthly",
			DurationDays: 30,
			IsActive:     true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           "plan_002",
			Name:         "Premium Plan",
			Description:  "This plan offers premium features.",
			Price:        19.99,
			BillingCycle: "monthly",
			DurationDays: 30,
			IsActive:     true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(plans); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
