cd ..

go test -coverprofile cover.out
go tool cover -html=cover.out -o cover.html

start cover.html