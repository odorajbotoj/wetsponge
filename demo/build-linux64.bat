@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o wetsponge-v2.0.0_dev3-linux64 main.go
pause