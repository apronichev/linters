package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
)

var GlobalVar string // gochecknoglobals: global var

func main() {
	// revive (formerly golint): don't use underscores in Go names
	user_input := "example"

	// staticcheck: this condition is always true
	if len(user_input) >= 0 {
		fmt.Println("Input has a length")
	}

	// errcheck: error return value not checked
	_, _ = ioutil.ReadFile("nonexistent.txt")

	// gosec: use of weak cryptographic primitive
	h := md5.New()
	fmt.Println(h)

	// bodyclose: response body must be closed
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// gosimple: should use strings.ToUpper
	upper := ""
	for _, r := range user_input {
		upper += string(r - 32)
	}

	// ineffassign: ineffectual assignment
	y := 10
	y = 20
	y = 10

	fmt.Println(upper, y)
	complexFunction(2)
}

func complexFunction(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if i%3 == 0 {
				result += i
			} else if i%5 == 0 {
				result += i * 2
			}
		} else {
			if i%7 == 0 {
				result += i * 3
			} else if i%11 == 0 {
				result += i * 4
			} else {
				result += i
			}
		}
	}
	return result
}
