package main

import "fmt"

func main() {
	buah := []string{"nangka", "pisang", "durian", "apel", "pepaya"}
	x := "nangka"
	for i := 0; i < len(buah); i++ {
		if buah[i] == x {
			buah = append(buah[:i], buah[i+1:]...)
		}
	}
	fmt.Println(buah)
}
