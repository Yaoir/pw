pw: pw.go
	@go build pw.go

test:
	@go run pw.go

man: pw.1.ronn
	@ronn --roff --manual="User Commands" --organization="Jay Ts" --date="2018-12-25" pw.1.ronn > man1/pw.1
	@gzip -f pw.1
	@mv pw.1.gz man1
	@man -l man1/pw.1.gz

showman:
	@man -l man1/pw.1.gz

# local backup. Create the .bak directory first

backup back bup:
	@cp *.go Makefile push README.md .bak
