-- Create table for Invoice
CREATE TABLE invoices (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    subscription_id VARCHAR(36) NOT NULL,
    payment_id VARCHAR(36),
    amount DOUBLE PRECISION NOT NULL,
    status VARCHAR(20) NOT NULL, -- possible values: "draft", "pending", "paid", "cancelled"
    due_date TIMESTAMP NOT NULL,
    paid_date TIMESTAMP
);

-- Create table for Payment
CREATE TABLE payments (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    subscription_id VARCHAR(36) NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'USD',
    payment_method VARCHAR(50) NOT NULL,  -- e.g., "credit_card", "paypal"
    status VARCHAR(20) NOT NULL,          -- possible values: "pending", "successful", "failed"
    transaction_id VARCHAR(50) NOT NULL,  -- external payment processor transaction ID
    payment_date TIMESTAMP NOT NULL,
    failure_reason TEXT                   -- if payment failed, reason will be stored
);

-- Create table for SubscriptionPlan
CREATE TABLE subscription_plans (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    billing_cycle VARCHAR(20) NOT NULL, -- e.g., "monthly", "quarterly", "yearly"
    duration_days INT NOT NULL,
    is_active BOOLEAN NOT NULL
);

-- Create table for Subscription
CREATE TABLE subscriptions (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id VARCHAR(36) NOT NULL,         -- reference to the User Service
    plan_id VARCHAR(36) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,          -- values: "active", "cancelled", "expired", "pending"
    auto_renew BOOLEAN NOT NULL,
    cancelled_at TIMESTAMP
);

-- Create table for SubscriptionEvent
CREATE TABLE subscription_events (
    id VARCHAR(36) PRIMARY KEY,
    subscription_id VARCHAR(36) NOT NULL,
    event_type VARCHAR(50) NOT NULL,      -- e.g., "created", "renewed", "cancelled", "expired"
    event_data TEXT NOT NULL,             -- JSON data
    created_at TIMESTAMP NOT NULL
);

-- seed data
INSERT INTO subscription_plans (
  id,
  created_at,
  updated_at,
  name,
  description,
  price,
  billing_cycle,
  duration_days,
  is_active
)
VALUES (
  'plan_monthly',
  NOW(),
  NOW(),
  'Monthly Plan',
  'Access to basic features on a monthly basis.',
  20.00,
  'monthly',
  30,
  TRUE
);