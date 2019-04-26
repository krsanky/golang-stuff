package receiver

import "fmt"

type S1 struct {
	Name string
}

func (s1 *S1) Change1() {
	s1.Name = s1.Name + "1"
}

func (s1 *S1) Replace(s2 *S1) {
	s1 = s2
}

func Test1() {
	fmt.Printf("Test1\n")
	s1 := S1{}	
	fmt.Printf("s1.Name:%s\n", s1.Name)
	s1.Change1()
	fmt.Printf("s1.Name:%s\n", s1.Name)

	
	fmt.Printf("swap whole receiver struct...\n")
	fmt.Printf("s1.Name:%s\n", s1.Name)
	s2 := S1{}
	s2.Name = "s2-name"
	s1.Replace(&s2)
	fmt.Printf("s1.Name:%s\n", s1.Name)
}
