package repository

import (
	"errors"
	"subscriptions/internal/data"
	"sync"
)

type SubscriptionRepository interface {
	GetActiveSubscription(userID string) (*data.Subscription, error)
	SaveSubscription(sub *data.Subscription) error
	UpdateSubscription(sub *data.Subscription) error
}

type MemorySubscriptionRepository struct {
	mu            sync.RWMutex
	subscriptions map[string]*data.Subscription // keyed by userID
}

func NewMemorySubscriptionRepository() *MemorySubscriptionRepository {
	return &MemorySubscriptionRepository{
		subscriptions: make(map[string]*data.Subscription),
	}
}

func (repo *MemorySubscriptionRepository) GetActiveSubscription(userID string) (*data.Subscription, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	sub, exists := repo.subscriptions[userID]
	if !exists || sub.Status != "active" {
		return nil, errors.New("active subscription not found")
	}
	return sub, nil
}

func (repo *MemorySubscriptionRepository) SaveSubscription(sub *data.Subscription) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.subscriptions[sub.UserID] = sub
	return nil
}

func (repo *MemorySubscriptionRepository) UpdateSubscription(sub *data.Subscription) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.subscriptions[sub.UserID] = sub
	return nil
}
