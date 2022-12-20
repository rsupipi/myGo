package main

func main() {

	/* collection loop - slice*/
	a1 := []int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		print(a1[i], " : ")
	}
	println("collection looping") //1 : 2 : 3 : 4 : 5 : infinite loop with ;

	/* range loop - slice*/
	a2 := []int{1, 2, 3}
	for i, v := range a2 {
		println(i, v)
	}
	println("range - slice")

	/* looping map */
	a3 := map[string]int{"maths": 25, "science": 40}
	for i, v := range a3 {
		println(i, v)
	}
	println("range - map")

	/* print only values*/
	a4 := map[string]int{"maths": 25, "science": 40}
	for v := range a4 {
		println(v)
	}
	println("range only values - map")

	/* write only variable*/
	a5 := map[string]int{"maths": 25, "science": 40}
	for _, v := range a5 {
		println(v)
	}
	println("range with write only variables - map")

}
