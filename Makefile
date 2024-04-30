
# Directory where the GUI source code is located
GUI_DIR := gui

# Directory where the built GUI files will be placed
DIST_DIR := dist

# Name of the binary file
BINARY_NAME := app

# prepare dev enviroment
prepare:
	./GitHooksSetup.sh

# Lint the GUI source code
lint_gui:
	cd $(GUI_DIR) && npm run lint

# Build the GUI source code
build_gui:
	echo "Building GUI"
	cd $(GUI_DIR) && npm run build

# Generate a hash value for the binary file
hash:
	cd $(DIST_DIR) && sha256sum $(BINARY_NAME) > $(BINARY_NAME).sha256

# Include debugging information in the binary file
build_dev:
	go build -o $(DIST_DIR)/debug/$(BINARY_NAME) main.go

# Build the application by running lint_gui, build_gui, and hash targets
# Avoid debugging information in the binary file by using the -s and -w flags
build: lint_gui build_gui
	go build -ldflags="-s -w" -o $(DIST_DIR)/$(BINARY_NAME) main.go
	@make hash
	echo "Build complete"
