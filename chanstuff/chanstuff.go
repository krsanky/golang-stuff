package chanstuff

import (
	"fmt"
	"time"
)

func LongJob1(c chan int, nums ...int) {
	time.Sleep(100 * time.Millisecond)
	time.Sleep(2 * time.Second)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	//return sum
	c <- sum
}

func Chan1() {
	ch := make(chan int)
	go LongJob1(ch, 1, 2, 3, 4, 5)

	fmt.Printf("do other stuff...\n")
	fmt.Printf("do other stuff...\n")
	fmt.Printf("do other stuff...\n")
	fmt.Printf("do other stuff...\n")
	fmt.Printf("do other stuff...\n")

	s := <-ch
	fmt.Printf("chan sum:%d\n", s)
}
