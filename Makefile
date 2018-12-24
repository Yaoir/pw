pw: pw.go
	@go build pw.go

test:
	@go run pw.go

backup back bup:
	@cp *.go Makefile push .bak

install:
	cp pw /home/jay/.bin
