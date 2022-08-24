package main

import "fmt"

func main() {
	appenTest()
}

func appenTest() {
	a := []string{"Hanif", "Bunda Manda", "Ayah Fadli"}
	var b []*string
	for i := 0; i < len(a); i++ {
		b = append(b, &a[i])
	}
	fmt.Println(b)

	function := func(params []*string) {
		for _, v := range params {
			fmt.Println(*v)
		}
	}

	function(b)

}
