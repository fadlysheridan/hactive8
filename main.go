package main

import "fmt"

type data struct {
	name string
	umur string
}

func main() {
	structTest()
}

func structTest() {
	a := []string{"Hanif", "Manda", "Fadli"}
	var b []*data

	for i := 0; i < len(a); i++ {
		dataOrang := data{
			name: a[i],
		}

		b = append(b, &dataOrang)
	}

	function := func(params []*data) {
		for _, v := range params {
			fmt.Println(v.name)
		}
	}

	function(b)

}
