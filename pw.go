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

const DEFLEN = 15	// Default password length
var length int = DEFLEN	// Actual password length

func usage(cmdname string) {
	fmt.Fprintf(os.Stderr,"usage: %s [length]\n",cmdname);
	os.Exit(1);
}

var randomgen *rand.Rand

// Initialize random number generator

func init_random() {
//
	t := time.Now()
	if t.IsZero() {
		fmt.Fprintf(os.Stderr,"%s: Cannot initialize random number generator\n",os.Args[0])
		os.Exit(2)
	}
	randomgen = rand.New(rand.NewSource(t.UnixNano()))
}

// Return a random integer

func random() int {
//
        return randomgen.Int()
}

// Returns true if the rune is an ASCII digit

func isdigit(c rune) bool {
//
	if c >= '0' && c <= '9' { return true }
	return false
}

// Returns true if the rune is an uppercase letter

func isupper(c rune) bool {
//
	if c >= 'A' && c <= 'Z' { return true }
	return false
}

// Returns true if the rune is an lowercase letter

func islower(c rune) bool {
//
	if c >= 'a' && c <= 'z' { return true }
	return false
}

// Returns true if the rune is an ASCII alphanumeric

func isalphanumeric(c rune) bool {
	if isdigit(c) { return true }
	if isupper(c) { return true }
	if islower(c) { return true }
	return false
}

// Returns true if there is a digit in the string

func digitcheck(s []rune) bool {
//
	for i := 0; i < len(s); i++ { if isdigit(s[i]) { return true } }
	return false
}

// Returns true if there is an uppercase letter in the string

func uppercheck(s []rune) bool {
//
	for i := 0; i < len(s); i++ { if isupper(s[i]) { return true } }
	return false
}

// Returns true if there is an lowercase letter in the string

func lowercheck(s []rune) bool {
//
	for i := 0; i < len(s); i++ { if islower(s[i]) { return true } }
	return false
}

// Highest value of rune that is within range of ASCII characters

const max_ascii = 128

// Create a random string and return it

func generate() []rune {
//
	randstr := []rune{}

	for len(randstr) < length {
	//
		for {
		//
			n := rune(random()%max_ascii)
			if isalphanumeric(n) {
			//
				randstr = append(randstr,n)
				break
			}
		}
	}

	return randstr
}

func main() {
	var s string
	var err error
	var pw []rune

	if len(os.Args) > 2 { usage(os.Args[0]) }

	if len(os.Args) == 2 {
		length, err = strconv.Atoi(os.Args[1])
		if err != nil { usage(os.Args[0]) }

		if length < MINLEN || length > MAXLEN {
			if length < MINLEN { s = "small" } else { s = "large" }
			fmt.Fprintf(os.Stderr,"%s: length is too %s\n",os.Args[0],s)
			os.Exit(1)
		}
	}

	init_random()

	// generate passwords until one passes tests

	for {
		pw = generate()
		// make sure the password has at least one digit, one uppercase letter, and one lowercase letter
		if digitcheck(pw) && uppercheck(pw) && lowercheck(pw) { break }
	}

	fmt.Printf("%s\n",string(pw))
}
