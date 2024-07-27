BINARY_NAME=micro

build:
	go build -o ${BINARY_NAME} cmd/main.go

run:
	./${BINARY_NAME}