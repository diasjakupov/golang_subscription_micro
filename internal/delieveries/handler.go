package delieveries

import (
	"net/http"

	"subscriptions/internal/data"
	cancelsubscription "subscriptions/internal/services/cancel_subscription"
	checksubscription "subscriptions/internal/services/check_subscription"
	createsubscription "subscriptions/internal/services/create_subscription"
	"subscriptions/internal/services/plans"
	renewsubscription "subscriptions/internal/services/renew_subscription"

	"github.com/labstack/echo/v4"
)

// Handler holds all use cases for subscription operations.
type Handler struct {
	createUC   createsubscription.CreateSubscriptionUseCase
	checkUC    checksubscription.CheckSubscriptionUseCase
	cancelUC   cancelsubscription.CancelSubscriptionUseCase
	renewUC    renewsubscription.RenewSubscriptionUseCase
	getPlansUC plans.GetSubscriptionPlansUseCase
}

// NewHandler constructs a new Handler with all required use cases.
func NewHandler(
	createUC createsubscription.CreateSubscriptionUseCase,
	checkUC checksubscription.CheckSubscriptionUseCase,
	cancelUC cancelsubscription.CancelSubscriptionUseCase,
	renewUC renewsubscription.RenewSubscriptionUseCase,
	getPlansUC plans.GetSubscriptionPlansUseCase,
) *Handler {
	return &Handler{
		createUC:   createUC,
		checkUC:    checkUC,
		cancelUC:   cancelUC,
		renewUC:    renewUC,
		getPlansUC: getPlansUC,
	}
}

// sendErrorResponse is a helper function to standardize error responses.
func sendErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]string{"error": message})
}

// CreateSubscriptionHandler handles the creation of a new subscription.
func (h *Handler) CreateSubscriptionHandler(c echo.Context) error {
	var req data.CreateSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "invalid request payload")
	}

	sub, err := h.createUC.Execute(req)
	if err != nil {
		if err.Error() == "already subscribed" {
			return sendErrorResponse(c, http.StatusConflict, err.Error())
		}
		return sendErrorResponse(c, http.StatusBadRequest, "subscription creation failed: "+err.Error())
	}

	return c.JSON(http.StatusCreated, sub)
}

// CheckSubscriptionHandler retrieves the active subscription for a user.
func (h *Handler) CheckSubscriptionHandler(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "user_id is required")
	}

	sub, err := h.checkUC.Execute(userID)
	if err != nil {
		return sendErrorResponse(c, http.StatusNotFound, "no active subscription found")
	}

	return c.JSON(http.StatusOK, sub)
}

// CancelSubscriptionHandler cancels an existing subscription.
func (h *Handler) CancelSubscriptionHandler(c echo.Context) error {
	var req data.CancelSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "invalid request payload")
	}

	if err := h.cancelUC.Execute(req.UserID); err != nil {
		return sendErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "subscription cancelled successfully"})
}

// RenewSubscriptionHandler processes the renewal of a subscription.
func (h *Handler) RenewSubscriptionHandler(c echo.Context) error {
	var req data.RenewSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "invalid request payload")
	}

	if err := h.renewUC.Execute(req); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "renewal failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "subscription renewed successfully"})
}

// GetPlansHandler retrieves all subscription plans.
func (h *Handler) GetPlansHandler(c echo.Context) error {
	plans, err := h.getPlansUC.Execute()
	if err != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to retrieve plans: "+err.Error())
	}
	return c.JSON(http.StatusOK, plans)
}
