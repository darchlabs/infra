REMOTE=ec2-user@ec2-34-234-95-250.compute-1.amazonaws.com

WASM_VERSION=0.0.1

ONBOARDING_VERSION=0.0.1
ONBOARDING_PORT=5800

INFRA_VERSION=0.0.1
INFRA_PORT=5900
INFRA_DSN=

install:
	@brew install pulumi/tap/pulumi
	@pulumi new kubernetes-go  
	@go get github.com/pulumi/pulumi-kubernetes/sdk/v3

compile-wasm:
	@echo "[compile] Compiling wasm..."
	# @GOOS=js GOARCH=wasm go build -o bin/backend-v$(WASM_VERSION).wasm cmd/wasm/main.go
	@tinygo build -o bin/backend-v$(WASM_VERSION).wasm -target wasm cmd/wasm/main.go

compile-onboarding:
	@echo "[compile] Compiling onboarding service..."
	@go build -o bin/onboarding-v$(ONBOARDING_VERSION) cmd/onboarding/main.go

dev-onboarding:
	@echo "[dev] Running onboarding dev service..."
	@export $$(cat .env) && go run cmd/onboarding/main.go

compile-infra:
	@echo "[compile] Compiling infra service..."
	@go build -o bin/infra-v$(INFRA_VERSION) cmd/infra/main.go

compose-app-up:
	@echo "[composing app up]"
	@docker-compose -f manifest/$(env)/compose/docker-compose.yml up --build

compose-app-down:
	@echo "[composing app down]"
	@docker-compose -f manifest/$(env)/compose/docker-compose.yml down

dev-infra:
	@echo "[dev] Running infra dev service..."
	@PORT=$(INFRA_PORT) DSN=$(INFRA_DSN) go run cmd/onboarding/main.go

compile-remote:
	@echo "[compile-remote] Uploading golang file..."
	@scp backend/cmd/wasm/main.go $(REMOTE):~
	# @echo "[compile-remote] Compiling wasm..."
	# @ssh $(REMOTE) "GOOS=js GOARCH=wasm go build -o main.wasm main.go"
	# @echo "[compile-remote] Copy wasm file from remote..."
	# @scp $(REMOTE):~/main.wasm ./backend/bin
	# @echo "[compile-remote] Remove files from remote..."
	# @ssh $(REMOTE) "rm main.*"

apply-dev:
	@echo "[apply-dev] applying dev to cluster..."
	@kubectl apply -f manifest/dev/k8s/.

delete-dev:
	@echo "[apply-dev] applying dev to cluster..."
	@kubectl delete -f manifest/dev/k8s/.
