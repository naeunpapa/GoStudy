package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		mod3, mod5 := i % 3, i % 5		
		if mod3 == 0 || mod5 == 0 {			
			sum += i			
		}
	}	
	fmt.Println("sum : ", sum)
}