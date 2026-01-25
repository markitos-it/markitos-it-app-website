# Monitoring & Observability

![Monitoring](https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=1200&h=400&fit=crop)

## The Three Pillars

### 1. Metrics
Numerical measurements over time

### 2. Logs
Discrete events with context

### 3. Traces
Request paths through distributed systems

## Prometheus & Grafana

### Prometheus Configuration

```yaml
# prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'application'
    static_configs:
      - targets: ['app:3000']
    metrics_path: '/metrics'
```

### Application Metrics

```javascript
const client = require('prom-client');

// Create a Registry
const register = new client.Registry();

// Add default metrics
client.collectDefaultMetrics({ register });

// Custom counter
const httpRequestsTotal = new client.Counter({
  name: 'http_requests_total',
  help: 'Total number of HTTP requests',
  labelNames: ['method', 'route', 'status'],
  registers: [register]
});

// Custom histogram
const httpRequestDuration = new client.Histogram({
  name: 'http_request_duration_seconds',
  help: 'Duration of HTTP requests in seconds',
  labelNames: ['method', 'route', 'status'],
  registers: [register]
});

// Middleware
app.use((req, res, next) => {
  const end = httpRequestDuration.startTimer();
  
  res.on('finish', () => {
    httpRequestsTotal.inc({
      method: req.method,
      route: req.route?.path || req.path,
      status: res.statusCode
    });
    
    end({
      method: req.method,
      route: req.route?.path || req.path,
      status: res.statusCode
    });
  });
  
  next();
});

// Expose metrics
app.get('/metrics', async (req, res) => {
  res.set('Content-Type', register.contentType);
  res.end(await register.metrics());
});
```

## Structured Logging

### Winston Logger

```javascript
const winston = require('winston');

const logger = winston.createLogger({
  level: 'info',
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.errors({ stack: true }),
    winston.format.json()
  ),
  defaultMeta: { service: 'user-service' },
  transports: [
    new winston.transports.File({ filename: 'error.log', level: 'error' }),
    new winston.transports.File({ filename: 'combined.log' })
  ]
});

// Usage
logger.info('User logged in', {
  userId: '123',
  ip: '192.168.1.1',
  timestamp: new Date()
});

logger.error('Database connection failed', {
  error: error.message,
  stack: error.stack,
  database: 'users'
});
```

### ELK Stack Integration

```javascript
// Send logs to Elasticsearch
const { ElasticsearchTransport } = require('winston-elasticsearch');

logger.add(new ElasticsearchTransport({
  level: 'info',
  clientOpts: { node: 'http://elasticsearch:9200' },
  index: 'logs'
}));
```

## Distributed Tracing

### OpenTelemetry Setup

```javascript
const { NodeSDK } = require('@opentelemetry/sdk-node');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');

const sdk = new NodeSDK({
  traceExporter: new JaegerExporter({
    endpoint: 'http://jaeger:14268/api/traces'
  }),
  instrumentations: [getNodeAutoInstrumentations()]
});

sdk.start();

// Manual instrumentation
const { trace } = require('@opentelemetry/api');

const tracer = trace.getTracer('app-tracer');

async function processOrder(orderId) {
  return tracer.startActiveSpan('process-order', async (span) => {
    span.setAttribute('order.id', orderId);
    
    try {
      const order = await fetchOrder(orderId);
      span.addEvent('order-fetched', { size: order.items.length });
      
      await validateOrder(order);
      await chargePayment(order);
      
      span.setStatus({ code: SpanStatusCode.OK });
      return order;
    } catch (error) {
      span.recordException(error);
      span.setStatus({ code: SpanStatusCode.ERROR, message: error.message });
      throw error;
    } finally {
      span.end();
    }
  });
}
```

## Alert Rules

### Prometheus AlertManager

```yaml
# alerts.yml
groups:
  - name: application
    interval: 30s
    rules:
      - alert: HighErrorRate
        expr: |
          sum(rate(http_requests_total{status=~"5.."}[5m]))
          /
          sum(rate(http_requests_total[5m])) > 0.05
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High error rate detected"
          description: "Error rate is {{ $value | humanizePercentage }}"
      
      - alert: HighLatency
        expr: |
          histogram_quantile(0.95,
            rate(http_request_duration_seconds_bucket[5m])
          ) > 1
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High latency detected"
          description: "95th percentile latency is {{ $value }}s"
      
      - alert: ServiceDown
        expr: up == 0
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: "Service is down"
          description: "{{ $labels.instance }} is unreachable"
```

## Grafana Dashboards

### Sample PromQL Queries

```promql
# Request rate
rate(http_requests_total[5m])

# Error rate
sum(rate(http_requests_total{status=~"5.."}[5m]))
/ sum(rate(http_requests_total[5m]))

# 95th percentile latency
histogram_quantile(0.95,
  rate(http_request_duration_seconds_bucket[5m]))

# CPU usage
rate(process_cpu_seconds_total[5m]) * 100

# Memory usage
process_resident_memory_bytes / 1024 / 1024
```

## Health Checks

```javascript
// Kubernetes liveness & readiness
app.get('/health/live', (req, res) => {
  res.json({ status: 'ok' });
});

app.get('/health/ready', async (req, res) => {
  try {
    await database.ping();
    await redis.ping();
    res.json({ status: 'ready' });
  } catch (error) {
    res.status(503).json({ status: 'not ready', error: error.message });
  }
});
```

## SLOs & SLIs

### Service Level Indicators

```javascript
// Track SLI metrics
const sliMetrics = {
  availability: new client.Gauge({
    name: 'sli_availability',
    help: 'Percentage of successful requests'
  }),
  
  latency: new client.Histogram({
    name: 'sli_latency_seconds',
    help: 'Request latency for SLI',
    buckets: [0.1, 0.3, 0.5, 1, 2, 5]
  }),
  
  errorRate: new client.Gauge({
    name: 'sli_error_rate',
    help: 'Error rate for SLI'
  })
};

// Calculate SLI periodically
setInterval(async () => {
  const metrics = await calculateSLI();
  sliMetrics.availability.set(metrics.availability);
  sliMetrics.errorRate.set(metrics.errorRate);
}, 60000);
```

### Error Budgets

```
SLO: 99.9% availability
Error Budget: 0.1% = 43.2 minutes/month

If error budget is exhausted:
- Stop feature releases
- Focus on reliability
- Investigate incidents
```

## Best Practices

1. **Use structured logging** - JSON format
2. **Add correlation IDs** - Track requests across services
3. **Set meaningful alerts** - Avoid alert fatigue
4. **Create dashboards** - For different stakeholders
5. **Document runbooks** - For common incidents

**Observe, measure, improve! 📊**
