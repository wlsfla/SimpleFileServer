
@echo off

echo Windows 32bit binary complie
SET GOOS=windows& SET GOARCH=386&go build -o bin/SimpleFileServer-x86.exe -ldflags="-s -w" src/main.go

echo Windows 64bit binary complie
SET GOOS=windows& SET GOARCH=amd64&go build -o bin/SimpleFileServer-x64.exe -ldflags="-s -w" src/main.go