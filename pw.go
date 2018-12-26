package main

// Print a random password, using ASCII letters and digits.

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	)

// Minimum and Maximum Length Allowed
// The minimum is two because a digit is required in the password.
// If the minimum were set to 1, then the result would always be simply a single digit.

const MINLEN = 2
const MAXLEN = 256

/* The default password length is based on:
 *
 * First,
 * Current CPU speed per computer: about 10 billion clock cycles per sec = 1e10
 * Number of seconds in year: 32e6 (round up to 1e7)
 * Age of universe in years: 1e10 (rounded down a bit)
 * All of those multiplied together: T = 1e27 clock cycles
 *
 * Second,
 * a-e: 26 letters, A-E: 26 letters, 0-9: 10 digits
 * so we are working in base (26+26+10) = base 62
 *
 * To create a password long enough we're sure it won't
 * be broken (except by extremely wild chance, of course):
 *
 * length = log(base 62) 1e27 = 15 digits. 
 *
 */

// Default password length

const DEFLEN = 15

var length int = DEFLEN	// length of password
var password []rune	// password string

func usage(cmdname string) {
	fmt.Fprintf(os.Stderr,"usage: %s [length]\n",cmdname);
	os.Exit(1);
}

var randomgen *rand.Rand

// init random number generator

func init_random() {
//
        randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))
// TODO: check for errors in above (?)
}

// return a random integer

func random() int {
//
        return randomgen.Int()
}

// returns true if the rune is an ASCII digit

func isdigit(c rune) bool {
//
	if c >= '0' && c <= '9' { return true }
	return false
}

// TODO:
// func isupper(c rune) bool {
// func islower(c rune) bool {
// ( and use those in isalnum() )
// Switch to byte, rather than rune

// returns true if the rune is an ASCII alphanumeric

func isalnum(c rune) bool {
	if isdigit(c) { return true }
	if c >= 'a' && c <= 'z' { return true }
	if c >= 'A' && c <= 'Z' { return true }
	return false
}

// returns true if there is a digit in the string

func digitcheck(s []rune) bool {
//
	for i := 0; i < len(s); i++ { if isdigit(s[i]) { return true } }
	return false
}

// TODO:
// uppercheck()
// lowercheck()

// Highest value of rune that is within range of ASCII characters

const max_ascii = 128

func generate() {
//
	for len(password) < length {
	//
		for {
		//
			n := rune(random()%max_ascii)
			if isalnum(n) {
			//
				password = append(password,n)
				break
			}
		}
	}
}

func main() {

	if len(os.Args) > 2 { usage(os.Args[0]) }

	if len(os.Args) == 2 {
		length, _ = strconv.Atoi(os.Args[1])
// TODO: usage() if error!
		if length == 0 { usage(os.Args[0]) }
		// quietly enforce minimum/maximum length
// TODO: exit if length too small or too large
		if length < MINLEN { length = MINLEN }
		if length > MAXLEN { length = MAXLEN }
	}

	init_random()

	// generate passwords until one passes test

	for {
		password = []rune{}
		generate()
		// make sure the password has at least one digit
		if digitcheck(password) { break }
	}

// TODO: check length
	fmt.Printf("%s\n",string(password))
}
