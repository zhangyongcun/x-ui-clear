CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s'  -o x-ui-clear
upx x-ui-clear