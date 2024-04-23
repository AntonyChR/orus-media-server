
GUI_DIR := gui
DIST_DIR := dist

lint_gui:
	cd $(GUI_DIR) && npm run lint

build_gui:
	cd $(GUI_DIR) && npm run build

hash:
	cd $(DIST_DIR) && sha256sum app > app.sha256

build: lint_gui build_gui
	go build -o $(DIST_DIR)/app main.go
	@make hash