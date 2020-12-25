init:
	sudo npm install --global light-server
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" html

build:
	GOOS=js GOARCH=wasm go build -o output/html5/app.wasm app.go

run:
	http-server -p 8001 output/html5

build-run:
	GOOS=js GOARCH=wasm go build -o output/html5/app.wasm app.go
	# http-server -p 8001 output/html5
	light-server -s output/html5 -p 8001
