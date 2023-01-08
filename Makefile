BINARY_NAME = "algorithms-essentials"
BINARIES = "./bin"
GITHUB = "github.com/max-rodziyevsky/algorithms-essentials"
APP_DIR = "./cmd/algorithms-essentials"

init:
	@echo "::> Creating a module root..."
	@go mod init ${GITHUB}
	@go mod tidy
	@echo "::> Finished!"

build:
	@echo "::> Building..."
	@go build -o ${BINARIES}/${BINARY_NAME} ${APP_DIR}
	@echo "::> Finished!"

run:
	@go build -o ${BINARIES}/${BINARY_NAME} ${APP_DIR}
	@${BINARIES}/${BINARY_NAME}

clean:
	@echo "::> Cleaning..."
	@go clean
	@rm -rf ${BINARIES}
	@go mod tidy
	@echo "::> Finished"

.PNONY: init build run clean