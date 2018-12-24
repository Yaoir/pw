package main

/* Print a random password, using letters (upper- and lowercase) and digits. */

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	)

/* Minimum and Maximum Length Allowed */

const MINLEN = 2
const MAXLEN = 256

/* Default Length (if no argment provided) */

/* This is based on:
 * First,
 * current CPU speed per computer: about 10 billion per sec = 1e10
 * number of seconds in year: 32e6 -> 1e7
 * years in age of universe (appx) 1e9
 * all of those multiplied together: T = 1e26 clock cycles
 * Second,
 * a-e: 26 letters, A-E: 26 letters, 0-9: 10 digits
 * so we are working in base (26+26+10) = base 62
 * To create a password long enough we're sure it won't be broken,
 * except by chance,
 * length = log(base 26) 1e26 = 15 digits. 
 */

const DEFLEN = 15

/* password string */

var password [MAXLEN+1]rune

func usage(cmdname string) {
	fmt.Printf("usage: %s [length]\n",cmdname);
	os.Exit(1);
}

var randomgen *rand.Rand

/* init random number generator */

func srandom() {
//
        randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

/* replacement for C library random() function */

func random() int {
//
        return randomgen.Int()
}

var length int = DEFLEN

func isalnum(c rune) bool {
	if isdigit(c) { return true }
	if c >= 'a' && c <= 'z' { return true }
	if c >= 'A' && c <= 'Z' { return true }
	return false
}

const max_ascii = 128

func generate() {
//
	for i := 0; i < length; {
	//
		for {
		//
			n := rune(random()%max_ascii)
			if isalnum(n) {
			//
				password[i] = n
				i++
				break
			}
		}
	}
	password[length] = '\000'
}

func isdigit(c rune) bool {
	if c >= '0' && c <= '9' { return true }
	return false
}

/* returns true if there is a digit in the string */

func digitcheck(s []rune) bool {

	for i := 0; s[i] != '\000'; i++ { if isdigit(s[i]) { return true } }
	return false
}

func main() {

	if len(os.Args) > 2 { usage(os.Args[0]) }

	if len(os.Args) == 2 {
		length, _ = strconv.Atoi(os.Args[1])
		if length == 0 { usage(os.Args[0]) }
		/* quietly enforce minimum/maximum length */
		if length < MINLEN { length = MINLEN }
		if length > MAXLEN { length = MAXLEN }
	}

	srandom()

	/* generate passwords until one passes test */

	for {
		generate()
		if digitcheck(password[:]) { break }
	}

	fmt.Printf("%s\n",string(password[:length]))
}
