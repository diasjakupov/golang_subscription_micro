// Планы подписок
Table subscription_plans {
  id varchar(36) [pk]
  name varchar(100) [not null]
  description text
  price decimal(10,2) [not null]
  billing_cycle varchar(20) [not null, note: 'monthly, quarterly, yearly']
  duration_days int [not null]
  is_active boolean [not null, default: true]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

// Подписки
Table subscriptions {
  id varchar(36) [pk]
  user_id varchar(36) [ref: > subscribers.id, not null]
  plan_id varchar(36) [ref: > subscription_plans.id, not null]
  start_date timestamp [not null]
  end_date timestamp [not null]
  status varchar(20) [not null, default: "active", note: 'active, cancelled, expired, pending']
  auto_renew boolean [not null, default: true]
  cancelled_at timestamp
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
  
  indexes {
    (subscriber_id, status)
  }
}

// Платежи
Table payments {
  id varchar(36) [pk]
  subscription_id varchar(36) [ref: > subscriptions.id, not null]
  amount decimal(10,2) [not null]
  currency varchar(3) [not null, default: "USD"]
  payment_method varchar(50)
  status varchar(20) [not null, note: 'pending, successful, failed']
  transaction_id varchar(100)
  payment_date timestamp
  failure_reason text
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
  
  indexes {
    (subscription_id, status)
  }
}

// События подписок
Table subscription_events {
  id varchar(36) [pk]
  subscription_id varchar(36) [ref: > subscriptions.id, not null]
  event_type varchar(50) [not null, note: 'created, renewed, cancelled, expired, failed_payment']
  event_data json
  created_at timestamp [not null, default: `now()`]
  
  indexes {
    (subscription_id, event_type)
    created_at
  }
}

// Счета
Table invoices {
  id varchar(36) [pk]
  subscription_id varchar(36) [ref: > subscriptions.id, not null]
  payment_id varchar(36) [ref: > payments.id]
  amount decimal(10,2) [not null]
  status varchar(20) [not null, default: "pending", note: 'draft, pending, paid, cancelled']
  due_date timestamp
  paid_date timestamp
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}
