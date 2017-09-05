package mod1

import "testing"

func Test1(t *testing.T) {
	u := &User{Name: "Leto"}
	if u.Name != "Leto" {
		t.Errorf("basic error")
	}

	Modify(u)
	if u.Name != "Leto" {
		t.Errorf("u.Name != \"Leto\"")
	}

	Modify2(u)
	if u.Name != "Paul" {
		t.Errorf("u.Name != \"Paul\"")
	}
}
