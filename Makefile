
# Directory where the GUI source code is located
GUI_DIR := gui

# Binary dir
DIST_DIR := dist

# Name of the binary file
BINARY_NAME := app

# scripts
SCRIPTS := scripts

# Install dependencies
install:
	go mod tidy
	cd $(GUI_DIR) && npm install

# prepare dev enviroment
prepare:
	@echo "Preparing dev environment"
	chmod +x ./git-hooks/*
	chmod +x ./${SCRIPTS}/*.sh
	
	./${SCRIPTS}/check_dependencies.sh
	./${SCRIPTS}/set_git_hooks.sh
	./${SCRIPTS}/create_fake_data.sh

	@make install

	cp config.template.toml config.toml

# Lint the GUI source code
lint_gui:
	cd $(GUI_DIR) && npm run lint

# Build the GUI source code
build_gui:
	@echo "Building GUI"
	cd $(GUI_DIR) && npm run build

# Generate a hash value for the binary file
hash:
	cd $(DIST_DIR) && sha256sum $(BINARY_NAME) > $(BINARY_NAME).sha256

# Include debugging information in the binary file
build_dev: main.go
	go build -o $(DIST_DIR)/debug/$(BINARY_NAME) main.go

# Build the application by running build_gui and hash targets
# Avoid debugging information in the binary file by using the -s and -w flags
build: main.go build_gui
	@echo "Building application"
	go build -ldflags="-s -w" -o $(DIST_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete"
