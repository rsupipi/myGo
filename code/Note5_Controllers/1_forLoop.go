package main

/* define type underlying a basic type */

func main() {
	/* for loop */
	var a1 int // assigning the default value
	for a1 < 5 {
		a1++
		print(a1, " : ")
	}
	println("introduced for loop") //1 : 2 : 3 : 4 : 5 : introduced for loop

	/* break */
	var a2 int
	for a2 < 5 {
		a2++
		if a2 == 3 {
			break
		}
		print(a2, " : ")
	}
	println("break") //1 : 2 : break

	/* continue */
	var a3 int
	for a3 < 5 {
		a3++
		if a3 == 3 {
			continue
		}
		print(a3, " : ")
	}
	println("continue") //1 : 2 : 4 : 5 : continue

	/* increment in same line */
	var a4 int
	for ; a4 < 5; a4++ {
		print(a4, " : ")
	}
	println("increment in same line") //0 : 1 : 2 : 3 : 4 : increment in same line

	/* increment in same line */
	for a5 := 0; a5 < 5; a5++ {
		print(a5, " : ")
	}
	println("implicit initialization") //0 : 1 : 2 : 3 : 4 : implicit initialization

}
