.DEFAULT_GOAL := help

.PHONY: help app-go-start app-go-build app-docker-local-build app-docker-local-start app-clean app-deploy-tag app-delete-tag k8s-local-forward

help:
	@echo "📋 Available commands:"
	@echo ""
	@echo "  make app-go-start              - Start app with Go (development)"
	@echo "  make app-go-build              - Build Go binary to dist/app"
	@echo "  make app-docker-local-build    - Build Docker image with :local tag"
	@echo "  make app-docker-local-start    - Start Docker container locally"
	@echo "  make app-clean                 - Remove dist/ and Docker :local image"
	@echo "  make app-deploy-tag <version>  - Create and push git tag (e.g., 1.2.3)"
	@echo "  make app-delete-tag <version>  - Delete git tag locally and remotely"
	@echo "  make k8s-local-forward         - Forward K8s service to localhost:8080"
	@echo ""

app-go-start:
	bash bin/app/go-start.sh

app-go-build:
	bash bin/app/go-build.sh

app-docker-local-build:
	bash bin/app/docker-local-build.sh

app-docker-local-start:
	bash bin/app/docker-local-start.sh

app-clean:
	bash bin/app/clean.sh

app-deploy-tag:
	bash bin/app/deploy-tag.sh $(filter-out $@,$(MAKECMDGOALS))
app-delete-tag:
	bash bin/app/delete-tag.sh $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
