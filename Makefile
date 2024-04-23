
GUI_DIR := gui
DIST_DIR := dist
BINARY_NAME := app

lint_gui:
	cd $(GUI_DIR) && npm run lint

build_gui:
	cd $(GUI_DIR) && npm run build

hash:
	cd $(DIST_DIR) && sha256sum $(BINARY_NAME) > $(BINARY_NAME).sha256

build: lint_gui build_gui
	go build -o $(DIST_DIR)/$(BINARY_NAME) main.go
	@make hash