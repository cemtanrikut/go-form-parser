APP_NAME=go-form-parser
OUTPUT_DIR=output
OUTPUT_PDF=$(OUTPUT_DIR)/form_output.pdf
DOCKER_IMAGE=go-form-parser

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf $(APP_NAME) $(OUTPUT_PDF) $(OUTPUT_DIR)

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run --rm -v $(PWD)/form.xml:/app/form.xml -v $(PWD)/output:/app/output $(DOCKER_IMAGE)

.PHONY: build run clean docker-build docker-run
