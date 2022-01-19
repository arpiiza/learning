package main

import "fmt"

type customErr struct {
	help string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.help)
}

func foo(e error) {
	fmt.Println(e.(customErr).help)
}

func main() {
	c1 := customErr{
		help: "this went wrong",
	}
	foo(c1)
}
