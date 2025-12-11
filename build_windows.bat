set GOOS=windows
set GOARCH=amd64

go build -o build/adminServer.exe ./cmd/server/main.go
go build -o build/adminTimer.exe ./cmd/timerService/main.go

