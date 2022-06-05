SVC_NAME = demo01
PROTO_PATH = api/proto/$(SVC_NAME)/service/v1/
GOAPI_PATH = api/goapi/$(SVC_NAME)/service/v1/
CONFIG_PATH = internal/infrastructure/config/proto/
CONFIG_DATA = ./config/config.yaml
CONSTRUCTOR_PATH = ./internal/infrastructure/constructor/
CMD_PATH = ./cmd/$(SVC_NAME)/server/
ENV = DEV
MYSQL_PASSWORD = 000000

# build proto files to go files
protoc:
	protoc -I=third_party --proto_path=$(PROTO_PATH) --go_out=$(GOAPI_PATH) --go_opt=paths=source_relative --go-grpc_out=$(GOAPI_PATH) --go-grpc_opt=paths=source_relative $(PROTO_PATH)*.proto
	protoc --go_out=. --go_opt=paths=source_relative ${CONFIG_PATH}*.proto

# build application object
wire:
	wire $(CONSTRUCTOR_PATH)

# run in dev mode
run:
	export ENV=$(ENV) MYSQL_PASSWORD=$(MYSQL_PASSWORD) && go run $(CMD_PATH) --config=$(CONFIG_DATA)

# build in dev mode
build-dev:
	go build -o $(CMD_PATH)server $(CMD_PATH)

# build in prod mode
build-prod:
	go build -tags netgo -o $(CMD_PATH)server $(CMD_PATH)

# deployment in local k8s
deploy-dev:
	sh ./deployment/dev/command.sh

# deployment in test k8s
deploy-test:
	make build-prod
	mv ./cmd/$(SVC_NAME)/server/server ./deployment/test/

# deployment in test k8s
deploy-prod:
	make build-prod
	mv ./cmd/$(SVC_NAME)/server/server ./deployment/prod/
