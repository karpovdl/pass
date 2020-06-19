:: Script for creating an application cover

cd ../..

:: Del old data
IF EXIST cover.out DEL cover.out /Q /S
IF EXIST cover.html DEL cover.html /Q /S

:: Code cover
go test -coverprofile cover.out
go tool cover -html=cover.out -o cover.html

:: Run report
start cover.html