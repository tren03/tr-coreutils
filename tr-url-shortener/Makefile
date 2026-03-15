BINARY_NAME := url-shortener
BUILD_DIR := bin
MAIN_PACKAGE := ./cmd

.PHONY: build run clean

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)
