// Copyright 2014 Seth Hoenig. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// bcrypt-tool is command line tool for common bcrypt functions
// including the ability to generate hashes, determine if a password
// matches a hash, and compute cost from a hash.
package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

const (
	helpText = `Usage: bcrypt-tool [action] argument ...
  ACTIONS
    hash  [password] <cost> Generate hash given password and optional cost (4-31)
    match [password] [hash] Print "yes" and return 0 if password is a match
                            for hash, or print "no" and return 1 otherwise 
    cost  [hash]            Print the cost of hash (4-31)
`
)

func main() {
	os.Args = os.Args[1:]

	if len(os.Args) < 2 {
		help()
	}
	switch os.Args[0] {
	case "cost":
		if len(os.Args) != 2 {
			help()
		}
		c := cost(os.Args[1])
		fmt.Println(fmt.Sprintf("%d", c))
	case "match":
		if len(os.Args) != 3 {
			help()
		}
		ok := match(os.Args[1], os.Args[2])
		if ok {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
			os.Exit(1)
		}
	case "hash":
		if len(os.Args) > 4 {
			help()
		}
		passwd := os.Args[1]
		cost := bcrypt.DefaultCost
		if len(os.Args) == 3 {
			c, e := strconv.Atoi(os.Args[2])
			if e != nil {
				help()
			}
			cost = c
		}
		hash := hash(passwd, cost)
		fmt.Println(hash)
	default:
		help()
	}
}

func help() {
	fmt.Fprintln(os.Stderr, helpText)
	os.Exit(2)
}

func cost(hash string) int {
	h := []byte(hash)
	c, e := bcrypt.Cost(h)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(2)
	}
	return c
}

func match(password, hash string) bool {
	p := []byte(password)
	h := []byte(hash)
	e := bcrypt.CompareHashAndPassword(h, p)
	if e != nil {
		return false
	}
	return true
}

func hash(password string, cost int) string {
	p := []byte(password)
	h, e := bcrypt.GenerateFromPassword(p, cost)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(2)
	}
	return string(h)
}
