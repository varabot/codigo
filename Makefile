VSCODE_ARTIFACT_URL=$(shell cat assets/vscode_url.txt)

build:
	go build ./cmd/codigo

download:
	mkdir -p ~/.codigo
	curl -#Lo ~/.codigo/vscode-web.zip $(VSCODE_ARTIFACT_URL)
