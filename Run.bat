:: this bat file is supposed to build the application from source code, build the documents and
:: run the psychopass app after all of these are done.
:: created by ALiwoto (woto@kaizoku.cyou)
@echo off
TITLE Building wotoplatform binary file
go mod tidy
go build -o WotoPlatform.exe && .\WotoPlatform.exe
