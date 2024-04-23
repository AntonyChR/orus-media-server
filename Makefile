
# Directory where the GUI source code is located
GUI_DIR := gui

# Directory where the built GUI files will be placed
DIST_DIR := dist

# Name of the binary file
BINARY_NAME := app

# Lint the GUI source code
lint_gui:
	cd $(GUI_DIR) && npm run lint

# Build the GUI source code
build_gui:
	cd $(GUI_DIR) && npm run build

# Generate a hash value for the binary file
hash:
	cd $(DIST_DIR) && sha256sum $(BINARY_NAME) > $(BINARY_NAME).sha256

# Build the application by running lint_gui, build_gui, and hash targets
build: lint_gui build_gui
	go build -o $(DIST_DIR)/$(BINARY_NAME) main.go
	@make hash