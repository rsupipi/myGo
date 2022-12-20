package main

func main() {

	/* infinite loop*/
	var a1 int
	for {
		a1++
		print(a1, " : ")
		if a1 == 5 {
			break
		}
	}
	println("infinite loop with ;") //1 : 2 : 3 : 4 : 5 : infinite loop with ;

	/* infinite loop*/
	var a2 int
	for {
		a2++
		print(a2, " : ")
		if a2 == 5 {
			break
		}
	}
	println("infinite loop without ;") //1 : 2 : 3 : 4 : 5 : infinite loop without ;

}
