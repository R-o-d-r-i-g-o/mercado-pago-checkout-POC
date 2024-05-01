.PHONY: dev-destroy
.PHONY: clean-cache

install:
	@echo Validating dependecies...
	@go mod tidy
	@echo Generating vendor...
	@go mod vendor

clean-cache:
	@echo Cleaning golang cache...
	@go clean -modcache
	@echo Deleting vendor folder if exists...
	@if exists .\vendor rmdir /s /q .\vendor

format:
	@gofmt -w .

run:
	@go run ./cmd/main.go

dev-up:
	@cd ./infra/local && docker-compose -p "code-space-api" up -d

dev-stop:
	@cd ./infra/local && docker-compose -p "code-space-api" stop

dev-destroy:
	@cd ./infra/local && if exist .\.docker rmdir /s /q .\.docker

docker:
	@docker build -f ./infra/deploy/Dockerfile -t codespace/api:latest .

generate-tree:
	@tree -I 'vendor' -o direc_tree.txt

test:
	@go test ./... -short -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

coverage:
	@echo "${COLOR_YELLOW}Running project coverage...${COLOR_WHITE}\n"
	@go test ./... -v -coverprofile=cover.out
	@go tool cover -html=cover.out -o cover.html
	@echo "${COLOR_GREEN}Coverage completed successfully.${COLOR_WHITE}"