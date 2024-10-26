@echo off

go build -o json_sys\json_sys.exe json_sys\json_sys.go
go build server.go