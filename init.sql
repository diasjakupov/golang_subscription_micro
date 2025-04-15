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

