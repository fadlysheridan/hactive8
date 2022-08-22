package main

import "fmt"

func main() {
	names := []string{"fadli", "fisa", "rafi", "kaadek", "kevin", "iqbal", "joko", "budi", "ismed", "david"}

	for i, t := range names {
		fmt.Println(i+1, "Name", t)
	}
}
