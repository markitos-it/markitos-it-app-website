# Kubernetes Networking Deep Dive

![Kubernetes Network](https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1200&h=400&fit=crop)

## Network Model

Kubernetes networking follows these principles:
- Every Pod gets its own IP address
- Pods can communicate without NAT
- Nodes can communicate with Pods without NAT

## Service Types

### ClusterIP (Default)

Internal-only access:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: ClusterIP
  selector:
    app: myapp
  ports:
    - port: 80
      targetPort: 8080
```

### NodePort

Exposes on each Node's IP:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-nodeport
spec:
  type: NodePort
  selector:
    app: myapp
  ports:
    - port: 80
      targetPort: 8080
      nodePort: 30080
```

### LoadBalancer

Cloud provider integration:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-loadbalancer
spec:
  type: LoadBalancer
  selector:
    app: myapp
  ports:
    - port: 80
      targetPort: 8080
```

## Ingress

Route external HTTP(S) traffic:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: myapp.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: my-service
            port:
              number: 80
```

## Network Policies

Control traffic between Pods:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-frontend
spec:
  podSelector:
    matchLabels:
      app: backend
  policyTypes:
  - Ingress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
    ports:
    - protocol: TCP
      port: 8080
```

## DNS in Kubernetes

Services are accessible via DNS:
- `service-name.namespace.svc.cluster.local`
- `service-name.namespace`
- `service-name` (within same namespace)

## CNI Plugins

Popular choices:
- **Calico** - Feature-rich, network policies
- **Cilium** - eBPF-based, high performance
- **Flannel** - Simple overlay network
- **Weave** - Encrypted mesh network

## Troubleshooting

```bash
# Check service endpoints
kubectl get endpoints my-service

# Test connectivity from a Pod
kubectl run tmp-shell --rm -i --tty --image nicolaka/netshoot

# Inside the pod:
curl http://my-service
nslookup my-service
```

## Performance Tips

1. Use **headless services** for direct Pod access
2. Enable **topology-aware routing** for local traffic
3. Configure **resource limits** for network plugins
4. Use **host networking** for high-performance workloads

**Network is the backbone of your cluster! 🌐**
