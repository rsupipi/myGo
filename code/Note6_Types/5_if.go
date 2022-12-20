package main

type User struct {
	id   int
	name string
}

func main() {
	u1 := User{
		id:   1,
		name: "pipi",
	}

	u2 := User{
		id:   2,
		name: "ruchi",
	}

	if u1 == u2 {
		println("u1 = u2")
	} else if u1.name != u2.name {
		println("u1.name != u2.name")
	} else {
		println("else what?")
	}
	//u1.name != u2.name
}
