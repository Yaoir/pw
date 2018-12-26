# Makefile for pw
#
# Written for GNU make, but will probably work
# with any Unix/POSIX-compatible make program
#
# Requirements:
#	Compile (build): You must have Go installed.
#	To make and view the manual page,
#		ronn, gzip, man

# Modify these two settings to fit your needs:

# 1. Where to install the pw program:
BINDIR=/home/jay/.bin
# Or perhaps one of:
#BINDIR=/usr/local/bin
#BINDIR=/usr/bin

# 2. Where to install the manual page:
MANDIR=/usr/local/man/man1
# or maybe one of:
#MANDIR=/usr/local/share/man/man1
#MANDIR=/usr/share/man/man1

# Compile (build) the program

pw: pw.go
	@go build pw.go

test:
	@go run pw.go

# Make the manual page

man: pw.1.ronn
	@ronn --roff --manual="User Commands" --organization="Jay Ts" --date="2018-12-25" pw.1.ronn > man1/pw.1
	@gzip -f pw.1
	@mv pw.1.gz man1
	@man -l man1/pw.1.gz

# Display the manual page

showman:
	@man -l man1/pw.1.gz

# Install the pw program and its manual page

install:
	@cp pw $BINDIR
	@cp man1/pw.1 $MANDIR


# Local backup. Create the .bak directory first,
# and modify the file list as necessary.

backup back bup:
	@cp .gitignore *.go Makefile push README.md pw.1.ronn TODO .bak
