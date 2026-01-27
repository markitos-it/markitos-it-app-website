.DEFAULT_GOAL := help

.PHONY: help app-start app-go-build app-docker-local-build app-docker-local-start app-clean app-deploy-tag app-delete-tag k8s-local-forward

help:
	@echo "📋 Available commands:"
	@echo ""
	@echo "  make app-start               	- Start app with Go (development)"
	@echo "  make app-clean                 - Remove dist/ and Docker :local image"
	@echo "  make app-deploy-tag <version>  - Create and push git tag (e.g., 1.2.3)"
	@echo "  make app-delete-tag <version>  - Delete git tag locally and remotely"
	@echo ""

app-start:
	bash bin/app/start.sh

app-clean:
	bash bin/app/clean.sh

app-deploy-tag:
	bash bin/app/deploy-tag.sh $(filter-out $@,$(MAKECMDGOALS))
app-delete-tag:
	bash bin/app/delete-tag.sh $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
