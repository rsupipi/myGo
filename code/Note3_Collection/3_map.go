package Note3_Collection

import "fmt"

func main() {
	m := map[string]int{"foo": 22}
	fmt.Println(m)        //map[foo:22]
	fmt.Println(m["foo"]) //22

	m["foo"] = 27
	fmt.Println(m) //map[foo:27]

	delete(m, "foo")
	fmt.Println(m) //map[]

}
