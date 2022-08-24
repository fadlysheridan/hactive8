package main

import "fmt"

func main() {
	namaBanyak()
}

func namaBanyak() {
	fadli := "fadli"
	david := "david"
	list := []*string{&fadli, &david}

	function := func(l []*string) {
		for _, v := range l {
			fmt.Println(*v)
		}
	}
	function(list)

}
