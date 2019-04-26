package main

import (
	"fmt"
	"os"

	"github.com/krsanky/golang-stuff/chanstuff"
	"github.com/krsanky/golang-stuff/gstuff"
	"github.com/krsanky/golang-stuff/receiver"
)

func main() {
	//	for i, a := range os.Args[1:] {
	//		fmt.Printf("%d:%s ", i, a)
	//	}
	if len(os.Args) >= 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "chan":
			fmt.Printf("this: %s\n", gstuff.FnOne("hey"))
			chanstuff.Chan1()
		case "inter":
			receiver.Test1()
		default:
			usage()
		}
	} else {
		usage()
	}

}

func usage() {
	fmt.Println()
	fmt.Printf("gstuff [chan|inter]\n")
	fmt.Println()
}
