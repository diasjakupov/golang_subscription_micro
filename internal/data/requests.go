package data

type CreateSubscriptionRequest struct {
	UserID         string         `json:"user_id"`
	PlanID         string         `json:"plan_id"`
	PaymentDetails PaymentDetails `json:"payment_details"`
}

type PaymentDetails struct {
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type RenewSubscriptionRequest struct {
	UserID         string         `json:"user_id"`
	PaymentDetails PaymentDetails `json:"payment_details"`
}

type CancelSubscriptionRequest struct {
	UserID string `json:"user_id"`
}
