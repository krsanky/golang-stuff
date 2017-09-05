package mod1

type User struct {
	Name string
}

func Modify(u *User) {
	u = &User{Name: "Paul"}
}

func Modify2(u *User) {
	u.Name = "Paul"
}

func F1() (str string) {
	str = "str..."
	return
}
