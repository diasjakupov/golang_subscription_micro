@startuml
title Subscription Request Flow

actor HTTP
participant API as "API Gateway"
participant Payment as "PaymentService"
participant Subscription as "SubscriptionService"
participant Queue as "Message Queue"

HTTP->>API: POST /api/v1/subscriptions
note over API: Validate request & authenticate

API->>Subscription: Get/Create Subscriber(userID)
Subscription-->>API: Subscriber Ready

API->>Payment: Process Payment(paymentDetails)
alt Payment failed
    Payment-->>API: Payment failed
    API-->>HTTP: 400 Bad Request (payment issue)
    note over HTTP, API: Early termination
end

Payment-->>Subscription: Create Subscription(userID, planID)
Subscription-->>API: Subscription Created

Subscription->>Queue: Publish SubscriptionCreatedEvent

API-->>HTTP: 201 Created (subscription details)

Queue-->>API: Send welcome email  

@enduml



/*************/
@startuml
title Subscribe Request Flow (Already Subscribed)

actor HTTP
participant API as "API Gateway"
participant Subscription as "SubscriptionService"

HTTP->>API: Send Subscribe Request
note over API: Validate request & authenticate

API->>Subscription: Check if Subscriber Exists
Subscription-->>API: Exists

API->>Subscription: Check for Active Subscription
Subscription-->>API: Active Subscription Exists

API-->>HTTP: Return "Already Subscribed" Response

@enduml



/*************/
@startuml
title Check Request Flow

actor HTTP
participant API as "API Gateway"
participant Subscription as "SubscriptionService"

HTTP->>API: Send Request (Check)
note over API: Determine request type

API->>Subscription: Retrieve Subscription Data
Subscription-->>API: Active Subscription Info

API-->>HTTP: Return Active Subscription Info

@enduml


/*************/
@startuml
title Cancel Request Flow

actor HTTP
participant API as "API Gateway"
participant Subscription as "SubscriptionService"
participant Queue as "Message Queue"

HTTP->>API: Send Cancel Request
note over API: Validate request & authenticate

API->>Subscription: Retrieve Subscription Record
Subscription-->>API: Active Subscription Exists

API->>Subscription: Update Subscription Status to "Cancelled"
Subscription-->>API: Subscription Cancelled

Subscription->>Queue: Publish "Subscription Cancelled" Event

API-->>HTTP: Return Cancellation Success

@enduml


/*************/
@startuml
title Renew Request Flow

actor HTTP
participant API as "API Gateway"
participant Subscription as "SubscriptionService"
participant Payment as "PaymentService"
participant Queue as "Message Queue"

HTTP->>API: Send Renew Request
note over API: Validate request & authenticate

API->>Subscription: Retrieve Subscription Record
Subscription-->>API: Active Subscription Exists

API->>Payment: Process Renewal Payment
Payment-->>API: Payment Success

API->>Subscription: Update Subscription End Date
Subscription-->>API: Subscription Updated

Subscription->>Queue: Publish "Subscription Renewed" Event

API-->>HTTP: Return Renewal Success

Queue-->>API: Send renew email

@enduml

