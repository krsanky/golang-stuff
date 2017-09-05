package gstuff

import (
	"fmt"
	"testing"
)

func TestFnOne(t *testing.T) {
	str := "fart"
	ret := FnOne(str)
	fmt.Printf("%s\n", ret)
	//t.Fail()
}

func TestFnTwo(t *testing.T) {
	//t.Errorf("fail fail fail")
}

func TestFn3(t *testing.T) {
	//t.Fatalf("failed and stop all")
}
