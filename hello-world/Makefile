build-HelloWorldFunction:
	powershell -Command "$$env:GOOS='linux'; $$env:GOARCH='amd64'; $$env:CGO_ENABLED='0'; go build -o bootstrap main.go"
	copy bootstrap "$(ARTIFACTS_DIR)\bootstrap"