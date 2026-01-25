# Getting Started with Keptn

![Keptn Banner](https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=1200&h=400&fit=crop)

## Introduction

Keptn is an event-based control plane for continuous delivery and automated operations. It helps you orchestrate your deployments and makes sure everything runs smoothly.

## Installation

First, install the Keptn CLI:

```bash
curl -sL https://get.keptn.sh | bash
keptn install --platform=kubernetes
```

## Key Features

- **Automated Operations**: Self-healing and auto-remediation
- **Quality Gates**: Automated quality evaluation
- **Multi-Stage Delivery**: Progressive delivery across environments

## Quick Start

1. Create a new project:
```bash
keptn create project sockshop --shipyard=./shipyard.yaml
```

2. Add a service:
```bash
keptn add-resource --project=sockshop --service=carts --stage=dev
```

3. Deploy your first version:
```bash
keptn trigger delivery --project=sockshop --service=carts --image=docker.io/keptnexamples/carts:0.13.1
```

## Next Steps

- Configure your quality gates
- Set up monitoring integration
- Define your shipyard stages

**Happy shipping with Keptn! 🚀**
