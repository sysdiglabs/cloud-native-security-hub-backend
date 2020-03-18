.PHONY: test build push

test:
	GO15VENDOREXPERIMENT=0 ginkgo -r .

dev:
	RESOURCES_PATH=test/fixtures/resources VENDOR_PATH=test/fixtures/vendors go run cmd/server/main.go

watch:
	ag -l | entr -c go test -v ./...

build:
	docker build -f Dockerfile.server -t gcr.io/mateo-burillo-ns/promhub-backend .
	docker build -f Dockerfile.dbimport -t gcr.io/mateo-burillo-ns/promhub-dbimport .

push: build
	docker push gcr.io/mateo-burillo-ns/securityhub-backend
	docker push gcr.io/mateo-burillo-ns/securityhub-dbimport

deploy: deploy-backend deploy-frontend

deploy-backend:
	kubectl -n securityhub patch deployment backend -p "{\"spec\": {\"template\": {\"metadata\": { \"labels\": {  \"redeploy\": \"$(shell date +%s)\"}}}}}"

deploy-frontend:
	kubectl -n securityhub patch deployment frontend -p "{\"spec\": {\"template\": {\"metadata\": { \"labels\": {  \"redeploy\": \"$(shell date +%s)\"}}}}}"
