### Introduction

pw is a secure password generator. It creates passwords that are alphanumeric strings that contain at least one digit.

### Quick Start

Compile:

```
$ go build pw.go
```
or if you have GNU make installed:
```
$ make
```

Run:

```
$ pw
YVSBve2ec7izj4y
```

### Manual Page


```
PW(1)                            User Commands                           PW(1)



NAME
       pw - secure password generator

SYNOPSIS
       pw [length]

DESCRIPTION
       pw(1)  generates and prints a string composed of randomly-chosen upper-
       and lowercase letters and digits from the ASCII character set (in other
       words, Unicode UTF-8 with all code points less than 128). Specifically,
       letters A-F and a-f along with digits 0-9 are used.

       The string contains at least one digit.

OPTIONS
       The length of the generated password can be specified with an  optional
       argument, which must be an integer between 2 and 256.

       Short  passwords  are  not secure, but are allowed so that pw(1) can be
       used to generate random alphanumeric strings for other purposes.

       By default, a length of 15 is used.

EXAMPLES
       Print a 15-character password:

       pw

       Print a 10-character password:

       pw 10

       Use pw(1) to create a file with five random characters in its name:

       echo ´hello, world´ >"hello-$(pw 5)"

BUGS
       Punctuation characters are not included in passwords. There is no check
       to  make sure the password contains at least one lowercase and at least
       one uppercase letter.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                           December 2018                           PW(1)
```
