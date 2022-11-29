BINARY_NAME=main.out

all: vet staticcheck test build

build:
	go build -o ${BINARY_NAME} main.go

test:
	go test ./...

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

dep:
	go mod download

vet:
	go vet

staticcheck:
	staticcheck ./...
