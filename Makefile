
GUI_DIR := gui

lint_gui:
	cd $(GUI_DIR) && npm run lint

build_gui:
	cd $(GUI_DIR) && npm run build

hash:
	cd dist && sha256sum app > app.sha256

build: lint_gui build_gui
	go build -o dist/app main.go
	@make hash