set GOROOT=c:\go
set GOOS=windows

cd ..

set GOARCH=amd64
go build -ldflags "-s -w" -i -v -o bin/release/64/pass.exe
go build -race -i -v -o bin/debug/64/pass.exe

set GOARCH=386
go build -ldflags "-s -w" -i -v -o bin/release/32/pass.exe
go build -i -v -o bin/debug/32/pass.exe