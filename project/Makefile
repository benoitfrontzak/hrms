FRONT_END_BINARY=frontEndApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
EMPLOYEE_BINARY=employeeApp
CLAIM_BINARY=claimApp
LEAVE_BINARY=leaveApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_employee build_claim build_leave
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_employee: builds the mail binary as a linux executable
build_employee:
	@echo "Building employee binary..."
	cd ../employee-service && env GOOS=linux CGO_ENABLED=0 go build -o ${EMPLOYEE_BINARY} ./cmd/api
	@echo "Done!"

## build_claim: builds the mail binary as a linux executable
build_claim:
	@echo "Building claim binary..."
	cd ../claim-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CLAIM_BINARY} ./cmd/api
	@echo "Done!"

## build_leave: builds the mail binary as a linux executable
build_leave:
	@echo "Building leave binary..."
	cd ../leave-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LEAVE_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
## @-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo "Stopped front end!"