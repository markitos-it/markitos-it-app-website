# Microservices Design Patterns

![Microservices](https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1200&h=400&fit=crop)

## What are Microservices?

An architectural style where an application is composed of small, independent services that communicate over well-defined APIs.

## Core Patterns

### 1. API Gateway Pattern

Single entry point for all clients:

```javascript
// API Gateway with Express
const express = require('express');
const httpProxy = require('http-proxy');

const app = express();
const proxy = httpProxy.createProxyServer();

// Route to user service
app.all('/api/users/*', (req, res) => {
  proxy.web(req, res, { target: 'http://user-service:3001' });
});

// Route to order service
app.all('/api/orders/*', (req, res) => {
  proxy.web(req, res, { target: 'http://order-service:3002' });
});

// Route to product service
app.all('/api/products/*', (req, res) => {
  proxy.web(req, res, { target: 'http://product-service:3003' });
});

app.listen(3000);
```

### 2. Service Discovery

Dynamic service location:

```javascript
// Using Consul for service discovery
const Consul = require('consul');
const consul = new Consul();

// Register service
async function registerService() {
  await consul.agent.service.register({
    name: 'user-service',
    address: '10.0.1.5',
    port: 3001,
    check: {
      http: 'http://10.0.1.5:3001/health',
      interval: '10s'
    }
  });
}

// Discover service
async function discoverService(serviceName) {
  const result = await consul.health.service({
    service: serviceName,
    passing: true
  });
  return result[0].Service;
}
```

### 3. Circuit Breaker

Prevent cascading failures:

```javascript
const CircuitBreaker = require('opossum');

// Create circuit breaker
const options = {
  timeout: 3000,
  errorThresholdPercentage: 50,
  resetTimeout: 30000
};

async function callExternalService() {
  const response = await fetch('http://external-service/api');
  return response.json();
}

const breaker = new CircuitBreaker(callExternalService, options);

// Use it
breaker.fire()
  .then(data => console.log(data))
  .catch(err => console.error('Service unavailable'));

// Monitor circuit breaker
breaker.on('open', () => console.log('Circuit opened!'));
breaker.on('halfOpen', () => console.log('Circuit half-open, testing...'));
breaker.on('close', () => console.log('Circuit closed, back to normal'));
```

### 4. Saga Pattern

Distributed transactions:

```javascript
// Order saga orchestrator
class OrderSaga {
  async createOrder(orderData) {
    const sagaId = generateId();
    
    try {
      // Step 1: Create order
      const order = await orderService.create(orderData);
      
      // Step 2: Reserve inventory
      await inventoryService.reserve(order.items, sagaId);
      
      // Step 3: Process payment
      await paymentService.charge(order.total, sagaId);
      
      // Step 4: Send confirmation
      await notificationService.send(order.userId, sagaId);
      
      return { success: true, order };
      
    } catch (error) {
      // Compensating transactions
      await this.compensate(sagaId, error);
      return { success: false, error };
    }
  }
  
  async compensate(sagaId, error) {
    // Rollback in reverse order
    await notificationService.cancel(sagaId);
    await paymentService.refund(sagaId);
    await inventoryService.release(sagaId);
    await orderService.cancel(sagaId);
  }
}
```

### 5. Event Sourcing

Store events instead of state:

```javascript
// Event store
class EventStore {
  constructor() {
    this.events = [];
  }
  
  append(event) {
    event.timestamp = Date.now();
    this.events.push(event);
  }
  
  getEvents(aggregateId) {
    return this.events.filter(e => e.aggregateId === aggregateId);
  }
  
  replay(aggregateId) {
    const events = this.getEvents(aggregateId);
    return events.reduce((state, event) => {
      return this.apply(state, event);
    }, {});
  }
  
  apply(state, event) {
    switch(event.type) {
      case 'OrderCreated':
        return { ...state, status: 'pending', items: event.items };
      case 'OrderPaid':
        return { ...state, status: 'paid' };
      case 'OrderShipped':
        return { ...state, status: 'shipped', tracking: event.tracking };
      default:
        return state;
    }
  }
}
```

## Communication Patterns

### Synchronous (REST)

```javascript
// REST API call
const response = await fetch('http://user-service/api/users/123');
const user = await response.json();
```

### Asynchronous (Message Queue)

```javascript
// RabbitMQ publisher
const amqp = require('amqplib');

async function publishEvent(event) {
  const conn = await amqp.connect('amqp://localhost');
  const channel = await conn.createChannel();
  
  await channel.assertQueue('events');
  channel.sendToQueue('events', Buffer.from(JSON.stringify(event)));
  
  await channel.close();
  await conn.close();
}

// Consumer
async function consumeEvents() {
  const conn = await amqp.connect('amqp://localhost');
  const channel = await conn.createChannel();
  
  await channel.assertQueue('events');
  
  channel.consume('events', (msg) => {
    const event = JSON.parse(msg.content.toString());
    console.log('Received:', event);
    channel.ack(msg);
  });
}
```

## Data Management

### Database per Service

Each service owns its database:

```
User Service → PostgreSQL (users)
Order Service → MongoDB (orders)
Inventory Service → Redis (inventory)
Analytics Service → ClickHouse (events)
```

### CQRS (Command Query Responsibility Segregation)

Separate read and write models:

```javascript
// Write model (Commands)
class OrderCommandHandler {
  async createOrder(command) {
    const order = new Order(command);
    await orderRepository.save(order);
    await eventBus.publish(new OrderCreatedEvent(order));
  }
}

// Read model (Queries)
class OrderQueryHandler {
  async getOrder(orderId) {
    return await orderReadRepository.findById(orderId);
  }
  
  async getOrdersByUser(userId) {
    return await orderReadRepository.findByUserId(userId);
  }
}
```

## Observability

### Distributed Tracing

```javascript
// OpenTelemetry
const { trace } = require('@opentelemetry/api');

const tracer = trace.getTracer('order-service');

async function processOrder(orderId) {
  const span = tracer.startSpan('process-order');
  span.setAttribute('order.id', orderId);
  
  try {
    await validateOrder(orderId);
    await chargePayment(orderId);
    await updateInventory(orderId);
    
    span.setStatus({ code: SpanStatusCode.OK });
  } catch (error) {
    span.recordException(error);
    span.setStatus({ code: SpanStatusCode.ERROR });
  } finally {
    span.end();
  }
}
```

## Best Practices

1. **Single Responsibility**: One service, one purpose
2. **Loose Coupling**: Minimize dependencies
3. **High Cohesion**: Related functionality together
4. **Fault Tolerance**: Design for failure
5. **Observability**: Logs, metrics, traces

**Build resilient distributed systems! 🏗️**
