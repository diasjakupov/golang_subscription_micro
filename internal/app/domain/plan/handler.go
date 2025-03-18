package plan

import (
	"net/http"
)

type PlanHandler interface {
	GetAllPlans(w http.ResponseWriter, r *http.Request)
}
