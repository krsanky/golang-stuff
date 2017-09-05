package main

import (
	"fmt"

	"github.com/krsanky/golang-stuff/chanstuff"
	"github.com/krsanky/golang-stuff/gstuff"
)

func main() {
	fmt.Printf("this: %s\n", gstuff.FnOne("hey"))
	chanstuff.Chan1()
}
