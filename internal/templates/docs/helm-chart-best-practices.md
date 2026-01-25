# Helm Chart Best Practices

![Helm Charts](https://images.unsplash.com/photo-1605745341075-1a6e8b9e7b8e?w=1200&h=400&fit=crop)

## Introduction

Helm is the package manager for Kubernetes. Learn how to create production-ready Helm charts that are maintainable and secure.

## Chart Structure

```
mychart/
  Chart.yaml
  values.yaml
  templates/
    deployment.yaml
    service.yaml
    ingress.yaml
    _helpers.tpl
```

## Chart.yaml Best Practices

```yaml
apiVersion: v2
name: myapp
description: A Helm chart for my awesome app
type: application
version: 1.0.0
appVersion: "2.1.0"
keywords:
  - app
  - kubernetes
maintainers:
  - name: Your Name
    email: you@example.com
```

## Values Organization

Keep your `values.yaml` organized and well-documented:

```yaml
# Application configuration
app:
  name: myapp
  replicas: 3
  
# Image configuration
image:
  repository: myregistry/myapp
  tag: "latest"
  pullPolicy: IfNotPresent

# Resource limits
resources:
  limits:
    cpu: 500m
    memory: 512Mi
  requests:
    cpu: 250m
    memory: 256Mi
```

## Template Functions

Use Helm's built-in functions:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "mychart.fullname" . }}
  labels:
    {{- include "mychart.labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.app.replicas }}
  selector:
    matchLabels:
      {{- include "mychart.selectorLabels" . | nindent 6 }}
```

## Helper Templates

Create reusable templates in `_helpers.tpl`:

```yaml
{{- define "mychart.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "mychart.labels" -}}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
app.kubernetes.io/name: {{ include "mychart.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
```

## Security Practices

1. **Never hardcode secrets** - use Kubernetes secrets
2. **Set security contexts**
3. **Use resource limits**
4. **Enable network policies**

## Testing Your Charts

```bash
# Lint your chart
helm lint ./mychart

# Dry run installation
helm install myapp ./mychart --dry-run --debug

# Template rendering
helm template myapp ./mychart
```

## Versioning Strategy

- Bump `version` for chart changes
- Bump `appVersion` for application updates
- Follow [SemVer](https://semver.org/)

**Remember: A well-crafted chart is a joy to maintain! ⚓**
