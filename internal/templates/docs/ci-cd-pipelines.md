# Modern CI/CD Pipelines

![CI/CD Pipeline](https://images.unsplash.com/photo-1667372393119-3d4c48d07fc9?w=1200&h=400&fit=crop)

## What is CI/CD?

**Continuous Integration (CI)**: Automatically test code changes
**Continuous Deployment (CD)**: Automatically deploy to production

## GitHub Actions Example

```yaml
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '20'
          
      - name: Install dependencies
        run: npm ci
        
      - name: Run tests
        run: npm test
        
      - name: Run linter
        run: npm run lint

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Build Docker image
        run: docker build -t myapp:${{ github.sha }} .
        
      - name: Push to registry
        run: |
          echo ${{ secrets.DOCKER_PASSWORD }} | docker login -u ${{ secrets.DOCKER_USERNAME }} --password-stdin
          docker push myapp:${{ github.sha }}

  deploy:
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Deploy to Kubernetes
        run: |
          kubectl set image deployment/myapp myapp=myapp:${{ github.sha }}
```

## GitLab CI Example

```yaml
stages:
  - test
  - build
  - deploy

variables:
  DOCKER_IMAGE: $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA

test:
  stage: test
  image: node:20
  script:
    - npm ci
    - npm test
    - npm run lint

build:
  stage: build
  image: docker:24
  services:
    - docker:24-dind
  script:
    - docker build -t $DOCKER_IMAGE .
    - docker push $DOCKER_IMAGE
  only:
    - main

deploy:production:
  stage: deploy
  image: bitnami/kubectl:latest
  script:
    - kubectl set image deployment/myapp myapp=$DOCKER_IMAGE
  only:
    - main
  environment:
    name: production
```

## Jenkins Pipeline

```groovy
pipeline {
    agent any
    
    stages {
        stage('Test') {
            steps {
                sh 'npm ci'
                sh 'npm test'
            }
        }
        
        stage('Build') {
            steps {
                sh 'docker build -t myapp:${BUILD_NUMBER} .'
            }
        }
        
        stage('Deploy') {
            when {
                branch 'main'
            }
            steps {
                sh 'kubectl apply -f k8s/'
                sh 'kubectl set image deployment/myapp myapp=myapp:${BUILD_NUMBER}'
            }
        }
    }
    
    post {
        always {
            junit 'test-results/*.xml'
        }
    }
}
```

## Best Practices

### 1. Keep Pipelines Fast
- Run tests in parallel
- Cache dependencies
- Use smaller Docker images

### 2. Security
- Store secrets securely
- Scan for vulnerabilities
- Sign your artifacts

### 3. Observability
- Send notifications on failures
- Track deployment metrics
- Generate deployment reports

### 4. Testing Strategy
```yaml
Unit Tests → Integration Tests → E2E Tests → Deploy
```

## Deployment Strategies

### Blue-Green Deployment
Deploy new version alongside old, switch traffic instantly.

### Canary Deployment
Gradually roll out to small percentage of users.

### Rolling Update
Replace instances one by one.

## Monitoring Your Pipeline

Track these metrics:
- **Build time**
- **Test coverage**
- **Deployment frequency**
- **Mean time to recovery (MTTR)**

**Automate everything, deploy with confidence! 🚀**
